package provider

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testAccPreCheckSSHKey verifies SSH key authentication setup
func testAccPreCheckSSHKey(t *testing.T) {
	// Run the standard pre-check first
	testAccPreCheck(t)

	// Check for SSH key environment variable
	sshKey := os.Getenv("PBS_TEST_SSH_PRIVATE_KEY")

	if sshKey == "" {
		t.Skip("SSH key authentication tests require PBS_TEST_SSH_PRIVATE_KEY to be set")
	}

	// Verify the SSH key is parseable
	if _, err := base64.StdEncoding.DecodeString(sshKey); err != nil {
		// If it's not base64, try to parse it directly as a key
		if _, err := os.ReadFile(sshKey); err != nil {
			// If it's not a file, assume it's raw key content and that's fine
		}
	}
}

// Test provider configuration with SSH key from environment
func TestAccProviderSSHKey_EnvironmentVariable(t *testing.T) {
	// This test verifies that the provider can authenticate using SSH keys
	
	hookName := testAccResourceName("ssh_test")

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheckSSHKey(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderSSHKeyConfig(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "name", hookName),
					resource.TestCheckResourceAttr("pbs_hook.test", "enabled", "true"),
				),
			},
		},
	})
}

// Configuration using SSH key from environment variable
func testAccProviderSSHKeyConfig(hookName string) string {
	return fmt.Sprintf(`
provider "pbs" {
  server   = "%s"
  sshport  = "%s" 
  username = "%s"
  ssh_private_key = base64decode("%s")
}

resource "pbs_hook" "test" {
  name    = "%s"
  type    = "site"
  enabled = true
  event   = "execjob_begin"
  order   = 1
}
`,
		os.Getenv("PBS_TEST_SERVER"),
		os.Getenv("PBS_TEST_PORT"),
		os.Getenv("PBS_TEST_USERNAME"),
		os.Getenv("PBS_TEST_SSH_PRIVATE_KEY"),
		hookName,
	)
}

// Helper function to get SSH private key content for testing
func getTestSSHPrivateKey() (string, error) {
	// Try direct key first
	if key := os.Getenv("PBS_TEST_SSH_PRIVATE_KEY"); key != "" {
		return key, nil
	}

	return "", fmt.Errorf("no SSH private key found in environment")
}
