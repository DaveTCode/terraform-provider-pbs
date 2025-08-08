package provider

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"

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
// Uses a cryptographically random suffix for better uniqueness than timestamps.
func testAccResourceName(prefix string) string {
	// Generate 6 random hex characters (3 bytes = 24 bits of entropy)
	// This gives us 16^6 = 16,777,216 possible combinations
	randomBytes := make([]byte, 3)
	if _, err := rand.Read(randomBytes); err != nil {
		// Fallback to a simple counter if crypto/rand fails (very unlikely)
		return fmt.Sprintf("%s_fb", prefix)
	}
	randomSuffix := fmt.Sprintf("%06x", randomBytes)[:6]

	// Truncate prefix if too long and add random suffix
	maxPrefixLen := 8 // Leave room for underscore and 6-char random suffix
	if len(prefix) > maxPrefixLen {
		prefix = prefix[:maxPrefixLen]
	}
	return fmt.Sprintf("%s_%s", prefix, randomSuffix)
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
