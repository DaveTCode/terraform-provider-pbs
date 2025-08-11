package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testAccPreCheckSSHKey verifies SSH key authentication setup.
func testAccPreCheckSSHKey(t *testing.T) {
	// Run the standard pre-check first
	testAccPreCheck(t)

	// Check for SSH key content environment variable
	if sshKeyContent := os.Getenv("PBS_TEST_SSH_PRIVATE_KEY"); sshKeyContent == "" {
		t.Skip("SSH key authentication tests require PBS_TEST_SSH_PRIVATE_KEY to be set with SSH private key content")
	}
}

// Test provider configuration with SSH key from environment.
func TestAccProviderSSHKey_EnvironmentVariable(t *testing.T) {
	// This test verifies that the provider can authenticate using SSH keys

	hookName := testAccResourceName("ssh_test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckSSHKey(t) },
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

// Configuration using SSH key from environment variable.
func testAccProviderSSHKeyConfig(hookName string) string {
	return fmt.Sprintf(`
provider "pbs" {
  server   = "%s"
  sshport  = "%s" 
  username = "%s"
  ssh_private_key = <<-EOT
%s
EOT
}

resource "pbs_hook" "test" {
  name        = "%s"
  enabled     = true
  debug       = true
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  alarm       = 60
  fail_action = "none"
}
`,
		os.Getenv("PBS_TEST_SERVER"),
		os.Getenv("PBS_TEST_PORT"),
		os.Getenv("PBS_TEST_USERNAME"),
		os.Getenv("PBS_TEST_SSH_PRIVATE_KEY"),
		hookName,
	)
}
