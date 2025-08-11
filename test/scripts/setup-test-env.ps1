# PBS Test Environment Setup Script (PowerShell)
# Windows equivalent of setup-test-env.sh

param(
    [Parameter(Position=0)]
    [string]$Action = "setup"
)

$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent (Split-Path -Parent $ScriptDir)
$ComposeDir = Join-Path $ProjectRoot "/test/docker"

Write-Host "=== PBS Terraform Provider Test Setup ===" -ForegroundColor Green

function Test-DockerRunning {
    try {
        docker info | Out-Null
        return $true
    }
    catch {
        Write-Host "Error: Docker is not running or not accessible" -ForegroundColor Red
        exit 1
    }
}

function Start-PbsContainer {
    Write-Host "Building and starting PBS container..." -ForegroundColor Yellow
    Push-Location $ComposeDir
    
    try {
        # Check if image exists
        $imageExists = $false
        try {
            docker image inspect pbs | Out-Null
            $imageExists = $true
        }
        catch {
            # Image doesn't exist
        }
        
        if (-not $imageExists) {
            Write-Host "Building PBS Docker image..." -ForegroundColor Yellow
            docker compose build
        }
        
        # Start the container
        Write-Host "Starting PBS container..." -ForegroundColor Yellow
        docker compose up -d
        
        # Wait for PBS to be ready
        Write-Host "Waiting for PBS to be ready..." -ForegroundColor Yellow
        Start-Sleep 30
        
        # Check if PBS is accessible
        try {
            docker compose exec -T pbs /opt/pbs/bin/qstat -s | Out-Null
        }
        catch {
            Write-Host "Warning: PBS may not be fully ready yet, waiting longer..." -ForegroundColor Yellow
            Start-Sleep 30
        }
    }
    finally {
        Pop-Location
    }
}

function New-TestData {
    Write-Host "Creating test data in PBS..." -ForegroundColor Yellow
    Push-Location $ComposeDir
    
    try {
        # Create test queue
        Write-Host "Creating test queue..." -ForegroundColor Yellow
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create queue test queue_type=execution" 2>$null
        }
        catch {
            Write-Host "Queue 'test' may already exist" -ForegroundColor Gray
        }
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test started=true"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test enabled=true"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test resources_default.nodes=1"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test resources_default.walltime=3600"
        
        # Set up workq as default queue
        Write-Host "Setting default queue..." -ForegroundColor Yellow
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set server default_queue=workq"
        
        # Create test node
        Write-Host "Creating test node..." -ForegroundColor Yellow
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create node pbs" 2>$null
        }
        catch {
            Write-Host "Node 'pbs' may already exist" -ForegroundColor Gray
        }
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs comment='Pre-existing node for import testing'"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs resources_available.ncpus=8"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs resources_available.mem=16gb"
        
        # Create test hook
        Write-Host "Creating test hook..." -ForegroundColor Yellow
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create hook test" 2>$null
        }
        catch {
            Write-Host "Hook 'test' may already exist" -ForegroundColor Gray
        }
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test enabled=true"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test event=execjob_begin"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test order=1"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test type=site"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test user=pbsadmin"
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test fail_action=none"
        
        # Create test resource
        Write-Host "Creating test resource..." -ForegroundColor Yellow
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create resource test type=size" 2>$null
        }
        catch {
            Write-Host "Resource 'test' may already exist" -ForegroundColor Gray
        }
        docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set resource test flag=h"
        
        Write-Host "Test data creation complete!" -ForegroundColor Green
    }
    finally {
        Pop-Location
    }
}

function Test-PbsInstallation {
    Write-Host "Verifying PBS installation..." -ForegroundColor Yellow
    Push-Location $ComposeDir
    
    try {
        # Test SSH connection
        try {
            docker compose exec -T pbs ssh -o StrictHostKeyChecking=no root@localhost -p 22 "echo 'SSH connection successful'" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: SSH connection to PBS container failed" -ForegroundColor Red
            return $false
        }
        
        # Test PBS commands
        try {
            docker compose exec -T pbs /opt/pbs/bin/qstat -s 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: PBS qstat command failed" -ForegroundColor Red
            return $false
        }
        
        # Test qmgr
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list server" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: PBS qmgr command failed" -ForegroundColor Red
            return $false
        }
        
        # Verify test resources exist
        Write-Host "Verifying test resources..." -ForegroundColor Yellow
        
        # Check if test queue exists
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list queue test" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: Test queue 'test' not found" -ForegroundColor Red
            return $false
        }
        
        # Check if test node exists
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list node pbs" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: Test node 'pbs' not found" -ForegroundColor Red
            return $false
        }
        
        # Check if test hook exists
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list hook test" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: Test hook 'test' not found" -ForegroundColor Red
            return $false
        }
        
        # Check if test resource exists
        try {
            docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list resource test" 2>$null | Out-Null
        }
        catch {
            Write-Host "Error: Test resource 'test' not found" -ForegroundColor Red
            return $false
        }
        
        Write-Host "PBS verification successful!" -ForegroundColor Green
        return $true
    }
    finally {
        Pop-Location
    }
}

function Set-TestEnvironment {
    Write-Host "Setting up test environment variables..." -ForegroundColor Yellow
    
    $envContent = @"
# PBS Test Environment Variables
`$env:PBS_TEST_SERVER = "localhost"
`$env:PBS_TEST_PORT = "2222"  
`$env:PBS_TEST_USERNAME = "root"
`$env:PBS_TEST_PASSWORD = "pbs"
`$env:TF_ACC = "1"
"@
    
    $envFile = Join-Path $ProjectRoot ".env.test.ps1"
    $envContent | Out-File -FilePath $envFile -Encoding UTF8
    
    Write-Host "Test environment file created: $envFile" -ForegroundColor Green
    Write-Host "To use these variables, run: . .\.env.test.ps1" -ForegroundColor Green
}

function Show-ContainerStatus {
    Write-Host "=== Container Status ===" -ForegroundColor Cyan
    Push-Location $ComposeDir
    docker compose ps
    
    Write-Host ""
    Write-Host "=== PBS Service Status ===" -ForegroundColor Cyan
    try {
        docker compose exec -T pbs /opt/pbs/bin/qstat -s 2>$null
        Write-Host "PBS services are running" -ForegroundColor Green
    }
    catch {
        Write-Host "PBS services may not be ready yet" -ForegroundColor Yellow
    }
    
    Pop-Location
}

function Start-TestSetup {
    Write-Host "Starting PBS test environment setup..." -ForegroundColor Green
    
    Test-DockerRunning
    Start-PbsContainer
    
    # Create test data first
    New-TestData
    
    # Try verification a few times
    for ($i = 1; $i -le 3; $i++) {
        if (Test-PbsInstallation) {
            break
        }
        else {
            if ($i -lt 3) {
                Write-Host "Verification failed, waiting and retrying... (attempt $i/3)" -ForegroundColor Yellow
                Start-Sleep 30
            }
            else {
                Write-Host "Warning: PBS verification failed after 3 attempts" -ForegroundColor Yellow
                Write-Host "Container may still be starting up. Check manually with:" -ForegroundColor Yellow
                Write-Host "  cd $ComposeDir && docker compose logs pbs" -ForegroundColor Yellow
            }
        }
    }
    
    Set-TestEnvironment
    Show-ContainerStatus
    
    Write-Host ""
    Write-Host "=== Setup Complete ===" -ForegroundColor Green
    Write-Host "PBS container is running and accessible on localhost:2222" -ForegroundColor Green
    Write-Host "Run tests with: make testacc" -ForegroundColor Green
    Write-Host "Or: . .\.env.test.ps1; go test -v -timeout 120m ./internal/provider/" -ForegroundColor Green
}

function Stop-TestEnvironment {
    Write-Host "=== Cleaning up ===" -ForegroundColor Red
    Push-Location $ComposeDir
    docker compose down
    Write-Host "PBS container stopped" -ForegroundColor Green
    Pop-Location
}

# Main execution based on action
switch ($Action.ToLower()) {
    "cleanup" {
        Stop-TestEnvironment
    }
    "status" {
        Show-ContainerStatus
    }
    default {
        Start-TestSetup
    }
}
