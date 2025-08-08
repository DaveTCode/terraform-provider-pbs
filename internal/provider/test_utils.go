package provider

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"pbs": providerserver.NewProtocol6WithError(New("test")()),
}

// testAccPreCheck verifies and sets up any required providers external
// configuration. This should be run during every acceptance test.
func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	// Check if we're running in Docker environment or if PBS is accessible
	if os.Getenv("PBS_TEST_SERVER") == "" {
		t.Setenv("PBS_TEST_SERVER", "localhost")
	}
	if os.Getenv("PBS_TEST_PORT") == "" {
		t.Setenv("PBS_TEST_PORT", "2222")
	}
	if os.Getenv("PBS_TEST_USERNAME") == "" {
		t.Setenv("PBS_TEST_USERNAME", "root")
	}
	if os.Getenv("PBS_TEST_PASSWORD") == "" {
		t.Setenv("PBS_TEST_PASSWORD", "pbs")
	}
}

// Helper function to generate unique names for test resources.
// PBS has strict naming limits (typically 15 chars max), so we keep names short.
func testAccResourceName(prefix string) string {
	// Use only the last 4 digits of timestamp to keep names short
	timestamp := time.Now().Unix() % 10000
	// Truncate prefix if too long and add short timestamp
	maxPrefixLen := 10 // Leave room for underscore and 4-digit timestamp
	if len(prefix) > maxPrefixLen {
		prefix = prefix[:maxPrefixLen]
	}
	return fmt.Sprintf("%s_%d", prefix, timestamp)
}

// providerConfig returns a basic provider configuration for testing.
func providerConfig() string {
	return fmt.Sprintf(`
provider "pbs" {
  server   = "%s"
  sshport  = "%s"
  username = "%s"
  password = "%s"
}
`,
		getEnvWithDefault("PBS_TEST_SERVER", "localhost"),
		getEnvWithDefault("PBS_TEST_PORT", "2222"),
		getEnvWithDefault("PBS_TEST_USERNAME", "root"),
		getEnvWithDefault("PBS_TEST_PASSWORD", "pbs"),
	)
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
