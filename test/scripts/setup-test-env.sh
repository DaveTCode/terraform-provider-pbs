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
        docker compose build
    fi
    
    # Start the container
    echo "Starting PBS container..."
    docker compose up -d
    
    # Wait for PBS to be ready
    echo "Waiting for PBS to be ready..."
    sleep 30
    
    # Check if PBS is accessible
    if ! docker compose exec -T pbs /opt/pbs/bin/qstat -s >/dev/null 2>&1; then
        echo "Warning: PBS may not be fully ready yet, waiting longer..."
        sleep 30
    fi
}

# Function to create test data in PBS
create_test_data() {
    echo "Creating test data in PBS..."
    cd "${COMPOSE_DIR}"
    
    # Create test queue
    echo "Creating test queue..."
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create queue test queue_type=execution" 2>/dev/null || echo "Queue 'test' may already exist"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test started=true"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test enabled=true"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test resources_default.nodes=1"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set queue test resources_default.walltime=3600"
    
    # Set up workq as default queue
    echo "Setting default queue..."
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set server default_queue=workq"
    
    # Create test node
    echo "Creating test node..."
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create node pbs" 2>/dev/null || echo "Node 'pbs' may already exist"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs comment='Pre-existing node for import testing'"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs resources_available.ncpus=8"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set node pbs resources_available.mem=16gb"
    
    # Create test hook
    echo "Creating test hook..."
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create hook test" 2>/dev/null || echo "Hook 'test' may already exist"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test enabled=true"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test event=execjob_begin"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test order=1"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test type=site"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test user=pbsadmin"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set hook test fail_action=none"
    
    # Create test resource
    echo "Creating test resource..."
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "create resource test type=size" 2>/dev/null || echo "Resource 'test' may already exist"
    docker compose exec -T pbs /opt/pbs/bin/qmgr -c "set resource test flag=h"
    
    echo "Test data creation complete!"
}

# Function to verify PBS is working
verify_pbs() {
    echo "Verifying PBS installation..."
    
    # Test SSH connection with password
    if ! docker compose exec -T pbs ssh -o StrictHostKeyChecking=no root@localhost -p 22 "echo 'SSH password connection successful'" 2>/dev/null; then
        echo "Error: SSH password connection to PBS container failed"
        return 1
    fi
    
    # Test PBS commands
    if ! docker compose exec -T pbs /opt/pbs/bin/qstat -s >/dev/null 2>&1; then
        echo "Error: PBS qstat command failed"
        return 1
    fi
    
    # Test qmgr
    if ! docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list server" >/dev/null 2>&1; then
        echo "Error: PBS qmgr command failed"
        return 1
    fi
    
    # Verify test resources exist
    echo "Verifying test resources..."
    
    # Check if test queue exists
    if ! docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list queue test" >/dev/null 2>&1; then
        echo "Error: Test queue 'test' not found"
        return 1
    fi
    
    # Check if test node exists
    if ! docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list node pbs" >/dev/null 2>&1; then
        echo "Error: Test node 'pbs' not found"
        return 1
    fi
    
    # Check if test hook exists
    if ! docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list hook test" >/dev/null 2>&1; then
        echo "Error: Test hook 'test' not found"
        return 1
    fi
    
    # Check if test resource exists
    if ! docker compose exec -T pbs /opt/pbs/bin/qmgr -c "list resource test" >/dev/null 2>&1; then
        echo "Error: Test resource 'test' not found"
        return 1
    fi
    
    echo "PBS verification successful!"
    return 0
}

# Function to generate SSH keys for testing
generate_ssh_keys() {
    echo "Generating SSH keys for testing..."
    
    SSH_KEY_DIR="${PROJECT_ROOT}/test/ssh-keys"
    mkdir -p "${SSH_KEY_DIR}"
    
    # Generate SSH key pair if it doesn't exist
    if [ ! -f "${SSH_KEY_DIR}/test_rsa" ]; then
        ssh-keygen -t rsa -b 4096 -f "${SSH_KEY_DIR}/test_rsa" -N "" -C "pbs-provider-test"
        echo "Generated SSH key pair in ${SSH_KEY_DIR}"
    else
        echo "SSH key pair already exists in ${SSH_KEY_DIR}"
    fi

    chmod 600 "${SSH_KEY_DIR}/test_rsa"
    chmod 644 "${SSH_KEY_DIR}/test_rsa.pub"
    
    # Copy public key to authorized_keys in container
    echo "Setting up SSH key authentication in PBS container..."
    cd "${COMPOSE_DIR}"
    
    # Create .ssh directory and authorized_keys file in container
    docker compose exec -T pbs mkdir -p /root/.ssh
    docker compose exec -T pbs chmod 700 /root/.ssh
    
    # Copy public key to container
    docker compose cp "${SSH_KEY_DIR}/test_rsa.pub" pbs:/root/.ssh/authorized_keys
    docker compose exec -T pbs chmod 600 /root/.ssh/authorized_keys
    docker compose exec -T pbs chown root:root /root/.ssh/authorized_keys
    
    # Configure SSH daemon to allow key authentication
    docker compose exec -T pbs sh -c "grep -q '^PubkeyAuthentication yes' /etc/ssh/sshd_config || echo 'PubkeyAuthentication yes' >> /etc/ssh/sshd_config"
    docker compose exec -T pbs sh -c "grep -q '^AuthorizedKeysFile' /etc/ssh/sshd_config || echo 'AuthorizedKeysFile .ssh/authorized_keys' >> /etc/ssh/sshd_config"
    docker compose exec -T pbs sh -c "grep -q '^PermitRootLogin yes' /etc/ssh/sshd_config || echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config"
    
    # Restart SSH daemon
    docker compose exec -T pbs service ssh restart || docker compose exec -T pbs /usr/sbin/sshd -D &
    
    echo "SSH key authentication configured successfully"
}

# Function to verify SSH key authentication
verify_ssh_key_auth() {
    echo "Verifying SSH key authentication..."
    
    SSH_KEY_DIR="${PROJECT_ROOT}/test/ssh-keys"
    
    # Debug: Check if key files exist
    if [ ! -f "${SSH_KEY_DIR}/test_rsa" ]; then
        echo "Error: Private key not found at ${SSH_KEY_DIR}/test_rsa"
        return 1
    fi
    
    if [ ! -f "${SSH_KEY_DIR}/test_rsa.pub" ]; then
        echo "Error: Public key not found at ${SSH_KEY_DIR}/test_rsa.pub"
        return 1
    fi
    
    # Debug: Check authorized_keys in container
    echo "Checking authorized_keys in container..."
    docker compose exec -T pbs cat /root/.ssh/authorized_keys 2>/dev/null || echo "No authorized_keys file found"
    
    # Test SSH connection with key
    echo "Testing SSH connection with key..."
    if ssh -i "${SSH_KEY_DIR}/test_rsa" -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o PasswordAuthentication=no -p 2222 root@localhost "echo 'SSH key authentication successful'" 2>/dev/null; then
        echo "SSH key authentication verified successfully"
        return 0
    else
        echo "Warning: SSH key authentication verification failed"
        echo "Debugging SSH connection..."
        ssh -i "${SSH_KEY_DIR}/test_rsa" -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o PasswordAuthentication=no -p 2222 -v root@localhost "echo 'SSH key authentication test'" 2>&1 | head -20
        echo "Password authentication will still work"
        return 1
    fi
}

# Function to setup test environment variables
setup_test_env() {
    echo "Setting up test environment variables..."
    
    SSH_KEY_DIR="${PROJECT_ROOT}/test/ssh-keys"
    SSH_PRIVATE_KEY=""
    
    # Read private key and encode it for environment variable
    if [ -f "${SSH_KEY_DIR}/test_rsa" ]; then
        # Use -w 0 for Linux base64, -i for macOS
        if base64 --help 2>&1 | grep -q "wrap"; then
            SSH_PRIVATE_KEY=$(base64 -w 0 "${SSH_KEY_DIR}/test_rsa")
        else
            SSH_PRIVATE_KEY=$(base64 -i "${SSH_KEY_DIR}/test_rsa" | tr -d '\n')
        fi
    fi
    
    cat > "${PROJECT_ROOT}/.env.test" << EOF
# PBS Test Environment Variables - Password Authentication
PBS_TEST_SERVER=localhost
PBS_TEST_PORT=2222
PBS_TEST_USERNAME=root
PBS_TEST_PASSWORD=pbs

# PBS Test Environment Variables - SSH Key Authentication
PBS_TEST_SSH_PRIVATE_KEY=${SSH_PRIVATE_KEY}

# Common Test Variables
TF_ACC=1
EOF
    
    echo "Test environment file created: ${PROJECT_ROOT}/.env.test"
    echo "To use these variables, run: source .env.test"
}

# Function to display container status
show_container_status() {
    echo "=== Container Status ==="
    cd "${COMPOSE_DIR}"
    docker compose ps
    
    echo ""
    echo "=== PBS Service Status ==="
    if docker compose exec -T pbs /opt/pbs/bin/qstat -s 2>/dev/null; then
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
    
    # Generate SSH keys for testing
    generate_ssh_keys
    
    # Create test data first
    create_test_data
    
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
                echo "  cd ${COMPOSE_DIR} && docker compose logs pbs"
            fi
        fi
    done
    
    # Verify SSH key authentication
    verify_ssh_key_auth
    
    setup_test_env
    show_container_status
    
    echo ""
    echo "=== Setup Complete ==="
    echo "PBS container is running and accessible on localhost:2222"
    echo "Authentication methods available:"
    echo "  - Password: root/pbs"
    echo "  - SSH Key: test/ssh-keys/test_rsa"
    echo "Run tests with: make testacc"
    echo "Or: source .env.test && go test -v -timeout 120m ./internal/provider/"
}

# Cleanup function
cleanup() {
    echo "=== Cleaning up ==="
    cd "${COMPOSE_DIR}"
    docker compose down
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
