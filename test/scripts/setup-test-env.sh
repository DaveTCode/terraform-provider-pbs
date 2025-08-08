#!/bin/bash

# Test setup script for PBS Terraform Provider
# This script sets up the Docker environment for testing

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
COMPOSE_DIR="${PROJECT_ROOT}/test/docker"

echo "=== PBS Terraform Provider Test Setup ==="

# Function to check if Docker is running
check_docker() {
    if ! docker info >/dev/null 2>&1; then
        echo "Error: Docker is not running or not accessible"
        exit 1
    fi
}

# Function to build and start PBS container
start_pbs_container() {
    echo "Building and starting PBS container..."
    cd "${COMPOSE_DIR}"
    
    # Build the container if it doesn't exist
    if ! docker image inspect pbs >/dev/null 2>&1; then
        echo "Building PBS Docker image..."
        docker-compose build
    fi
    
    # Start the container
    echo "Starting PBS container..."
    docker-compose up -d
    
    # Wait for PBS to be ready
    echo "Waiting for PBS to be ready..."
    sleep 30
    
    # Check if PBS is accessible
    if ! docker-compose exec -T pbs /opt/pbs/bin/qstat -s >/dev/null 2>&1; then
        echo "Warning: PBS may not be fully ready yet, waiting longer..."
        sleep 30
    fi
}

# Function to verify PBS is working
verify_pbs() {
    echo "Verifying PBS installation..."
    
    # Test SSH connection
    if ! docker-compose exec -T pbs ssh -o StrictHostKeyChecking=no root@localhost -p 22 "echo 'SSH connection successful'" 2>/dev/null; then
        echo "Error: SSH connection to PBS container failed"
        return 1
    fi
    
    # Test PBS commands
    if ! docker-compose exec -T pbs /opt/pbs/bin/qstat -s >/dev/null 2>&1; then
        echo "Error: PBS qstat command failed"
        return 1
    fi
    
    # Test qmgr
    if ! docker-compose exec -T pbs /opt/pbs/bin/qmgr -c "list server" >/dev/null 2>&1; then
        echo "Error: PBS qmgr command failed"
        return 1
    fi
    
    # Verify test resources exist
    echo "Verifying test resources..."
    
    # Check if test queue exists
    if ! docker-compose exec -T pbs /opt/pbs/bin/qmgr -c "list queue test" >/dev/null 2>&1; then
        echo "Error: Test queue 'test' not found"
        return 1
    fi
    
    # Check if test node exists
    if ! docker-compose exec -T pbs /opt/pbs/bin/qmgr -c "list node pbs" >/dev/null 2>&1; then
        echo "Error: Test node 'pbs' not found"
        return 1
    fi
    
    # Check if test hook exists
    if ! docker-compose exec -T pbs /opt/pbs/bin/qmgr -c "list hook test" >/dev/null 2>&1; then
        echo "Error: Test hook 'test' not found"
        return 1
    fi
    
    # Check if test resource exists
    if ! docker-compose exec -T pbs /opt/pbs/bin/qmgr -c "list resource test" >/dev/null 2>&1; then
        echo "Error: Test resource 'test' not found"
        return 1
    fi
    
    echo "PBS verification successful!"
    return 0
}

# Function to setup test environment variables
setup_test_env() {
    echo "Setting up test environment variables..."
    
    cat > "${PROJECT_ROOT}/.env.test" << EOF
# PBS Test Environment Variables
PBS_TEST_SERVER=localhost
PBS_TEST_PORT=2222
PBS_TEST_USERNAME=root
PBS_TEST_PASSWORD=pbs
TF_ACC=1
EOF
    
    echo "Test environment file created: ${PROJECT_ROOT}/.env.test"
    echo "To use these variables, run: source .env.test"
}

# Function to display container status
show_container_status() {
    echo "=== Container Status ==="
    cd "${COMPOSE_DIR}"
    docker-compose ps
    
    echo ""
    echo "=== PBS Service Status ==="
    if docker-compose exec -T pbs /opt/pbs/bin/qstat -s 2>/dev/null; then
        echo "PBS services are running"
    else
        echo "PBS services may not be ready yet"
    fi
}

# Main execution
main() {
    echo "Starting PBS test environment setup..."
    
    check_docker
    start_pbs_container
    
    # Try verification a few times
    for i in {1..3}; do
        if verify_pbs; then
            break
        else
            if [ $i -lt 3 ]; then
                echo "Verification failed, waiting and retrying... (attempt $i/3)"
                sleep 30
            else
                echo "Warning: PBS verification failed after 3 attempts"
                echo "Container may still be starting up. Check manually with:"
                echo "  cd docker_compose && docker-compose logs pbs"
            fi
        fi
    done
    
    setup_test_env
    show_container_status
    
    echo ""
    echo "=== Setup Complete ==="
    echo "PBS container is running and accessible on localhost:2222"
    echo "Run tests with: make testacc"
    echo "Or: source .env.test && go test -v -timeout 120m ./internal/provider/"
}

# Cleanup function
cleanup() {
    echo "=== Cleaning up ==="
    cd "${COMPOSE_DIR}"
    docker-compose down
    echo "PBS container stopped"
}

# Handle script arguments
case "${1:-}" in
    "cleanup")
        cleanup
        ;;
    "status")
        show_container_status
        ;;
    *)
        main
        ;;
esac
