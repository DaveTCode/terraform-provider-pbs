# Test Fixtures for PBS Terraform Provider

This directory contains test configurations and utilities for testing the PBS Terraform provider.

## Structure

- `basic/` - Simple test configurations for each resource type
- `complex/` - Multi-resource test scenarios  
- `edge_cases/` - Edge case and error condition tests
- `scripts/` - Helper scripts for test setup and teardown

## Usage

These fixtures are used by the acceptance tests in `internal/provider/*_test.go`.

Test configurations should be:
- Self-contained
- Use unique naming to avoid conflicts
- Clean up resources after testing
- Test both positive and negative scenarios

## Environment Variables

Set these environment variables for testing:

- `PBS_TEST_SERVER` - PBS server hostname (default: localhost)
- `PBS_TEST_PORT` - SSH port for PBS server (default: 2222)  
- `PBS_TEST_USERNAME` - SSH username (default: root)
- `PBS_TEST_PASSWORD` - SSH password (default: pbs)
- `TF_ACC` - Set to 1 to run acceptance tests
