default: fmt lint install generate

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

generate:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

# Test targets with Docker environment
test-setup:
	@echo "Setting up test environment..."
	@chmod +x test/scripts/setup-test-env.sh
	@test/scripts/setup-test-env.sh

test-cleanup:
	@echo "Cleaning up test environment..."  
	@test/scripts/setup-test-env.sh cleanup

test-status:
	@echo "Checking test environment status..."
	@test/scripts/setup-test-env.sh status

# Run acceptance tests with Docker PBS
testacc-docker: test-setup
	@echo "Running acceptance tests against Docker PBS..."
	@export PBS_TEST_SERVER=localhost && \
	 export PBS_TEST_PORT=2222 && \
	 export PBS_TEST_USERNAME=root && \
	 export PBS_TEST_PASSWORD=pbs && \
	 export TF_ACC=1 && \
	 go test -v -cover -timeout 120m ./internal/provider/

# Run specific test suites
test-queue:
	TF_ACC=1 go test -v -run "TestAccQueue" -timeout 30m ./internal/provider/

test-node:
	TF_ACC=1 go test -v -run "TestAccNode" -timeout 30m ./internal/provider/

test-hook:
	TF_ACC=1 go test -v -run "TestAccHook" -timeout 30m ./internal/provider/

test-server:
	TF_ACC=1 go test -v -run "TestAccServer" -timeout 30m ./internal/provider/

test-resource:
	TF_ACC=1 go test -v -run "TestAccPbsResource" -timeout 30m ./internal/provider/

test-datasources:
	TF_ACC=1 go test -v -run "TestAcc.*DataSource" -timeout 30m ./internal/provider/

test-integration:
	TF_ACC=1 go test -v -run "TestAccIntegration" -timeout 60m ./internal/provider/

.PHONY: fmt lint test testacc build install generate test-setup test-cleanup test-status testacc-docker test-queue test-node test-hook test-server test-resource test-datasources test-integration
