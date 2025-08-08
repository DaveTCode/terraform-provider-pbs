# PBS Terraform Provider Testing Guide

This guide covers the comprehensive testing strategy for the PBS Terraform Provider, including unit tests, acceptance tests, and integration tests against Docker-based PBS instances.

## Test Structure

### 1. Unit Tests (`internal/pbsclient/*_test.go`)
- Tests for PBS client functionality
- Parser validation for PBS command outputs
- No external dependencies

### 2. Acceptance Tests (`internal/provider/*_test.go`)
- Full Terraform provider acceptance tests
- Tests against real PBS instances
- Resource lifecycle testing (CRUD operations)
- Data source validation

### 3. Integration Tests (`internal/provider/integration_test.go`)
- Complex multi-resource scenarios
- Workflow testing
- Resource dependency validation
- End-to-end PBS configuration testing

## Running Tests

### Prerequisites

1. **Docker**: For running PBS in containers
2. **Go 1.20+**: For running the tests
3. **Make**: For using the provided Makefile targets

### Quick Start

1. **Setup test environment:**
   ```bash
   # On Linux/Mac
   ./test/scripts/setup-test-env.sh
   
   # On Windows PowerShell
   .\test\scripts\setup-test-env.ps1
   ```

2. **Run all acceptance tests:**
   ```bash
   make testacc-docker
   ```

3. **Run specific test suites:**
   ```bash
   make test-queue      # Queue resource tests
   make test-node       # Node resource tests  
   make test-hook       # Hook resource tests
   make test-server     # Server resource tests
   make test-resource   # PBS resource tests
   make test-datasources # Data source tests
   make test-integration # Integration tests
   ```

### Manual Test Setup

If you prefer manual setup:

1. **Start PBS container:**
   ```bash
   cd docker_compose
   docker-compose up -d
   
   # Wait for PBS to be ready
   sleep 60
   docker-compose exec pbs /opt/pbs/bin/qstat -s
   ```

2. **Set environment variables:**
   ```bash
   export TF_ACC=1
   export PBS_TEST_SERVER=localhost
   export PBS_TEST_PORT=2222
   export PBS_TEST_USERNAME=root
   export PBS_TEST_PASSWORD=pbs
   ```

3. **Run tests:**
   ```bash
   go test -v -timeout 120m ./internal/provider/
   ```

4. **Cleanup:**
   ```bash
   cd docker_compose
   docker-compose down
   ```

## Test Categories

### Resource Tests

Each PBS resource type has comprehensive tests covering:

#### Queue Resources (`queue_resource_test.go`)
- **Basic CRUD**: Create, read, update, delete operations
- **Resource defaults**: Testing walltime, ncpus, etc.
- **State management**: Enabled/disabled, started/stopped
- **Priority and limits**: max_running, max_queued
- **Import functionality**: Terraform import validation

#### Node Resources (`node_resource_test.go`)
- **Basic node configuration**: Name, port, state
- **Resource allocation**: Memory, CPUs, custom resources
- **Reservation settings**: resv_enable functionality
- **Custom PBS resources**: Integration with pbs_resource

#### Hook Resources (`hook_resource_test.go`)
- **Event handling**: Different hook events
- **Hook ordering**: Multiple hooks with ordering
- **Debug and alarm settings**: Advanced hook configuration
- **State transitions**: Enabled/disabled hooks
- **Multiple events**: Hooks with multiple event types

#### Server Resources (`server_resource_test.go`)
- **Basic configuration**: Server settings and parameters
- **Default chunks**: Resource chunk configuration
- **Resource defaults**: Server-wide resource defaults
- **Advanced settings**: Licensing, scheduling, provisioning

#### PBS Resources (`pbs_resource_resource_test.go`)
- **Resource types**: size, string, float, long, boolean, string_array
- **Resource flags**: Various flag combinations
- **Custom resources**: User-defined PBS resources

### Data Source Tests (`data_source_test.go`)

- **Queue data sources**: Reading existing queue configurations
- **Node data sources**: Querying node information
- **Server data sources**: Server configuration retrieval
- **Resource data sources**: PBS resource definitions
- **Hook data sources**: Hook configuration queries

### Integration Tests (`integration_test.go`)

#### Complete Workflow Tests
- **Full PBS setup**: Server, queues, nodes, resources, hooks
- **Resource dependencies**: Testing resource relationships
- **Configuration consistency**: End-to-end configuration validation

#### Dependency Tests
- **Queue hierarchies**: Routing and execution queues
- **Node resources**: Nodes using custom PBS resources
- **Hook ordering**: Multiple hooks with proper ordering

#### Complex Scenarios
- **Multi-resource workflows**: Real-world PBS configurations
- **Resource conflicts**: Testing error conditions
- **State management**: Complex state transitions

## Test Data and Fixtures

### Test Utilities (`test_utils.go`)
- **Provider configuration**: Reusable provider setup
- **Resource naming**: Unique test resource names
- **Environment setup**: Test environment validation
- **Cleanup functions**: Resource destruction verification

### Test Fixtures (`test/fixtures/`)
- **Basic configurations**: Simple resource examples
- **Complex scenarios**: Multi-resource setups
- **Edge cases**: Error condition testing
- **Template configurations**: Reusable test templates

## Continuous Integration

### GitHub Actions (`pbs-tests.yml`)
- **Automated testing**: Tests run on every PR and push
- **Matrix testing**: Multiple scenarios and configurations
- **PBS container management**: Automated Docker setup
- **Failure diagnostics**: Log collection and analysis

### Test Scenarios
1. **Basic queue management**
2. **Node resource allocation**
3. **Hook lifecycle management**
4. **Complex workflows**

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `TF_ACC` | - | Set to `1` to enable acceptance tests |
| `PBS_TEST_SERVER` | `localhost` | PBS server hostname |
| `PBS_TEST_PORT` | `2222` | SSH port for PBS server |
| `PBS_TEST_USERNAME` | `root` | SSH username |
| `PBS_TEST_PASSWORD` | `pbs` | SSH password |

## Debugging Tests

### Verbose Output
```bash
go test -v -run TestAccQueue -timeout 30m ./internal/provider/
```

### Single Test
```bash
go test -v -run TestAccQueueResource_basic -timeout 10m ./internal/provider/
```

### With Logs
```bash
TF_LOG=DEBUG go test -v -run TestAccQueue -timeout 30m ./internal/provider/
```

### PBS Diagnostics
```bash
# Check PBS status in container
cd docker_compose
docker-compose exec pbs /opt/pbs/bin/qstat -f
docker-compose exec pbs /opt/pbs/bin/pbsnodes -a
docker-compose exec pbs /opt/pbs/bin/qmgr -c "list server"
```

## Test Best Practices

1. **Unique naming**: Use `testAccResourceName()` for unique resource names
2. **Cleanup**: Always implement proper destroy checks
3. **Dependencies**: Test resource dependencies explicitly
4. **Error cases**: Include negative test scenarios
5. **Import testing**: Validate Terraform import functionality
6. **Timeouts**: Use appropriate timeouts for PBS operations
7. **Parallel execution**: Design tests to run in parallel when possible

## Troubleshooting

### Common Issues

1. **PBS not ready**: Increase wait time in container startup
2. **SSH connection failures**: Check container networking
3. **Resource conflicts**: Ensure unique resource naming
4. **Timeout errors**: Increase test timeouts for slow operations
5. **Permission issues**: Verify PBS permissions and authentication

### Debug Commands
```bash
# Container logs
docker-compose logs pbs

# PBS service status
docker-compose exec pbs systemctl status pbs

# Network connectivity
docker-compose exec pbs netstat -tlnp

# PBS configuration
docker-compose exec pbs cat /etc/pbs.conf
```

## Contributing

When adding new tests:

1. Follow the existing test patterns
2. Add both positive and negative test cases
3. Include proper cleanup functions
4. Update this documentation
5. Ensure tests pass in CI/CD pipeline
6. Add integration tests for complex scenarios

## Performance Considerations

- Tests run in parallel where possible
- Resource cleanup is thorough but efficient
- Container startup is optimized
- Test data is minimal but comprehensive
- Timeouts are realistic for PBS operations
