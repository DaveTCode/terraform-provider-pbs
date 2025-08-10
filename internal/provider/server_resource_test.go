package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccServerResource_createShouldFail tests that creating a server resource fails with a clear error.
func TestAccServerResource_createShouldFail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccServerResourceConfigBasic(),
				ExpectError: regexp.MustCompile("Server Resource Cannot Be Created"),
			},
		},
	})
}

// TestAccServerResource_dataSourceFirst tests reading server data before import.
func TestAccServerResource_dataSourceFirst(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfig("pbs"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_server.test", "name", "pbs"),
					resource.TestCheckResourceAttr("data.pbs_server.test", "id", "pbs"),
				),
			},
		},
	})
}

// TestAccServerResource_import tests importing a server resource.
func TestAccServerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        testAccServerResourceConfigBasic(),
				ResourceName:  "pbs_server.pbs",
				ImportState:   true,
				ImportStateId: "pbs",
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != "pbs" {
						return fmt.Errorf("expected ID %s, got %s", "pbs", state.ID)
					}

					if state.Attributes["name"] != "pbs" {
						return fmt.Errorf("expected name %s, got %s", "pbs", state.Attributes["name"])
					}

					return nil
				},
			},
		},
	})
}

// TestAccServerResource_importAndUpdate tests importing a server resource and then updating it.
func TestAccServerResource_importAndUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server
			{
				Config:             testAccServerResourceConfigForImport(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true, // This allows the state to persist to the next step
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != "pbs" {
						return fmt.Errorf("expected ID %s, got %s", "pbs", state.ID)
					}

					if state.Attributes["name"] != "pbs" {
						return fmt.Errorf("expected name %s, got %s", "pbs", state.Attributes["name"])
					}

					return nil
				},
			},
			// Update the server with new configuration
			{
				Config: testAccServerResourceConfigUpdated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "name", "pbs"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "comment", "Updated test server"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "300"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_array_size", "5000"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "node_fail_requeue", "600"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "eligible_time_enable", "true"),
				),
			},
			// Update again with different values to test multiple updates
			{
				Config: testAccServerResourceConfigUpdatedAgain(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "name", "pbs"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "comment", "Final test server configuration"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "900"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_array_size", "15000"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "node_fail_requeue", "120"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "eligible_time_enable", "false"),
				),
			},
		},
	})
}

// Helper functions.
func testAccServerResourceConfigBasic() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name    = "pbs"
  comment = "Basic test server"
}
`
}

func testAccServerResourceConfigForImport() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                     = "pbs"
  comment                  = "Imported test server"
  log_events               = 511
  mailer                   = "/usr/sbin/sendmail"
  mail_from                = "adm"
  query_other_jobs         = true
  resources_default = {
    ncpus = "1"
  }
  scheduler_iteration      = 600
  resv_enable              = true
  node_fail_requeue        = 310
  max_array_size           = 10000
  pbs_license_min          = 0
  pbs_license_max          = 2147483647
  pbs_license_linger_time  = 31536000
  eligible_time_enable     = false
  max_concurrent_provision = 5
  power_provisioning       = false
}
`
}

func testAccServerResourceConfigUpdated() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                     = "pbs"
  comment                  = "Updated test server"
  scheduler_iteration      = 300
  max_array_size           = 5000
  node_fail_requeue        = 600
  eligible_time_enable     = true
  log_events               = 511
  mailer                   = "/usr/sbin/sendmail"
  mail_from                = "adm"
  query_other_jobs         = true
  resources_default = {
    ncpus = "1"
  }
  resv_enable              = true
  pbs_license_min          = 0
  pbs_license_max          = 2147483647
  pbs_license_linger_time  = 31536000
  max_concurrent_provision = 5
  power_provisioning       = false
}
`
}

func testAccServerResourceConfigUpdatedAgain() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                     = "pbs"
  comment                  = "Final test server configuration"
  scheduler_iteration      = 900
  max_array_size           = 15000
  node_fail_requeue        = 120
  eligible_time_enable     = false
  log_events               = 511
  mailer                   = "/usr/sbin/sendmail"
  mail_from                = "adm"
  query_other_jobs         = true
  resources_default = {
    ncpus = "1"
  }
  resv_enable              = true
  pbs_license_min          = 0
  pbs_license_max          = 2147483647
  pbs_license_linger_time  = 31536000
  max_concurrent_provision = 5
  power_provisioning       = false
}
`
}

// TestAccServerResource_comprehensive_ACL tests comprehensive ACL handling including user format preservation.
func TestAccServerResource_comprehensive_ACL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server.
			{
				Config:             testAccServerResourceConfigACLImport(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Update with ACL settings in user-preferred order.
			{
				Config: testAccServerResourceConfigACLUpdated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "name", "pbs"),
					// Verify user format is preserved.
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users", "staff,admin"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_hosts", "host2.example.com,host1.example.com"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_resv_users", "user2,user1"),
					// Verify normalized format shows PBS ordering.
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users_normalized", "admin,staff"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_hosts_normalized", "host1.example.com,host2.example.com"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_resv_users_normalized", "user1,user2"),
				),
			},
			// Update ACL with different order to verify format preservation.
			{
				Config: testAccServerResourceConfigACLReordered(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "name", "pbs"),
					// Verify new user format is preserved.
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users", "admin,staff,manager"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_hosts", "host1.example.com,host3.example.com"),
					// Verify normalized format shows PBS ordering.
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users_normalized", "admin,manager,staff"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_hosts_normalized", "host1.example.com,host3.example.com"),
				),
			},
		},
	})
}

func testAccServerResourceConfigACLImport() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name = "pbs"
}
`
}

func testAccServerResourceConfigACLUpdated() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  comment          = "ACL test server"
  acl_user_enable  = true
  acl_users        = "staff,admin"
  acl_host_enable  = true
  acl_hosts        = "host2.example.com,host1.example.com"
  acl_resv_user_enable = true
  acl_resv_users   = "user2,user1"

  scheduler_iteration      = 900
  max_array_size           = 15000
  node_fail_requeue        = 120
  eligible_time_enable     = false
  log_events               = 511
  mailer                   = "/usr/sbin/sendmail"
  mail_from                = "adm"
  query_other_jobs         = true
  resources_default = {
    ncpus = "1"
  }
  resv_enable              = true
  pbs_license_min          = 0
  pbs_license_max          = 2147483647
  pbs_license_linger_time  = 31536000
  max_concurrent_provision = 5
  power_provisioning       = false
}
`
}

func testAccServerResourceConfigACLReordered() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  comment          = "ACL test server reordered"
  acl_user_enable  = true
  acl_users        = "admin,staff,manager"
  acl_host_enable  = true
  acl_hosts        = "host1.example.com,host3.example.com"
  acl_resv_user_enable = false
	
	scheduler_iteration      = 900
  max_array_size           = 15000
  node_fail_requeue        = 120
  eligible_time_enable     = false
  log_events               = 511
  mailer                   = "/usr/sbin/sendmail"
  mail_from                = "adm"
  query_other_jobs         = true
  resources_default = {
    ncpus = "1"
  }
  resv_enable              = true
  pbs_license_min          = 0
  pbs_license_max          = 2147483647
  pbs_license_linger_time  = 31536000
  max_concurrent_provision = 5
  power_provisioning       = false
}
`
}
