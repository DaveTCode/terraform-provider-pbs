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
					resource.TestCheckResourceAttr("pbs_server.pbs", "managers", "abcdefgh@*,bcdefgha@*,cdefghab@*,defghabc@*,efghabcd@*,fghabcde@*,ghabcdef@*,habcdefg@*"),
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

// TestAccServerResource_partialImportCleanPlan verifies that importing the server
// with a minimal configuration and then planning produces a clean, non-destructive
// plan. Before the fix for issue #101, attributes omitted from configuration were
// planned as null and turned into destructive qmgr `unset` commands.
func TestAccServerResource_partialImportCleanPlan(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server using a minimal (name-only) configuration.
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// A minimal configuration that omits every optional attribute must
			// produce an empty plan: omitted attributes retain their imported
			// values instead of being implicitly unset. PlanOnly fails the test
			// if the plan is non-empty.
			{
				Config:   testAccServerResourceConfigMinimal(),
				PlanOnly: true,
			},
		},
	})
}

// TestAccServerResource_partialConfigPreservesAttributes verifies that applying a
// partial configuration (as shown in the documented example) does not unset
// attributes that are present on the server but omitted from configuration.
func TestAccServerResource_partialConfigPreservesAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server.
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Establish known values for several attributes we will later omit.
			{
				Config: testAccServerResourceConfigPreserveSetup(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "managers", "operator@*,root@*"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "444"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "node_fail_requeue", "333"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_array_size", "7000"),
				),
			},
			// Apply the documented partial configuration (name + acl_users only).
			// The attributes set above are omitted here and must be preserved,
			// not unset. The framework also runs an idempotency plan after this
			// step, which must be empty.
			{
				Config: testAccServerResourceConfigDocumentedPartial(),
				Check: resource.ComposeTestCheckFunc(
					// The explicitly configured attribute is applied.
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users", "admin,staff"),
					// Omitted attributes retain their previously-set values.
					resource.TestCheckResourceAttr("pbs_server.pbs", "managers", "operator@*,root@*"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "444"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "node_fail_requeue", "333"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_array_size", "7000"),
				),
			},
		},
	})
}

// Helper functions.
func testAccServerResourceConfigMinimal() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name = "pbs"
}
`
}

func testAccServerResourceConfigPreserveSetup() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                = "pbs"
  managers            = "operator@*,root@*"
  scheduler_iteration = 444
  node_fail_requeue   = 333
  max_array_size      = 7000
}
`
}

// testAccServerResourceConfigDocumentedPartial mirrors the documented partial
// configuration example (examples/resources/pbs_server/basic.tf): only name and
// acl_users are configured, every other server attribute is omitted.
func testAccServerResourceConfigDocumentedPartial() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name      = "pbs"
  acl_users = "admin,staff"
}
`
}

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
	managers                 = "abcdefgh@*,bcdefgha@*,cdefghab@*,defghabc@*,efghabcd@*,fghabcde@*,ghabcdef@*,habcdefg@*"
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

// TestAccServerResource_partialConfigPreservesMapAttributes verifies that a map
// limit attribute present on the server (max_run_res) is populated into state on
// read and preserved across an unrelated update, instead of being destructively
// unset. This covers the createServerModel gap where max_run_res was never written
// to state.
func TestAccServerResource_partialConfigPreservesMapAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server.
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Establish a max_run_res entry on the server.
			{
				Config: testAccServerResourceConfigMaxRunRes(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_run_res.ncpus", "[u:PBS_GENERIC=10]"),
				),
			},
			// An unrelated update that omits max_run_res must preserve the entry,
			// not unset it. The framework idempotency plan after this step must
			// also be empty.
			{
				Config: testAccServerResourceConfigMaxRunResUnrelated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "comment", "max_run_res unrelated change"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "max_run_res.ncpus", "[u:PBS_GENERIC=10]"),
				),
			},
			// Explicitly clear the map attribute via unset_attributes.
			{
				Config: testAccServerResourceConfigMaxRunResUnset(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("pbs_server.pbs", "max_run_res.ncpus"),
				),
			},
		},
	})
}

func testAccServerResourceConfigMaxRunRes() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name = "pbs"
  max_run_res = {
    ncpus = "[u:PBS_GENERIC=10]"
  }
}
`
}

func testAccServerResourceConfigMaxRunResUnrelated() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name    = "pbs"
  comment = "max_run_res unrelated change"
}
`
}

func testAccServerResourceConfigMaxRunResUnset() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = ["max_run_res"]
}
`
}

// TestAccServerResource_unsetAttributes verifies that a scalar attribute can be
// explicitly reset via unset_attributes, while an attribute that is merely omitted
// is still preserved.
func TestAccServerResource_unsetAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server.
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Establish known values for default_queue and scheduler_iteration.
			{
				Config: testAccServerResourceConfigUnsetSetup(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "default_queue", "workq"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "450"),
				),
			},
			// Explicitly unset default_queue. scheduler_iteration is omitted and
			// must be preserved (not unset).
			{
				Config: testAccServerResourceConfigUnset(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("pbs_server.pbs", "default_queue"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "450"),
					resource.TestCheckResourceAttr("pbs_server.pbs", "unset_attributes.#", "1"),
				),
			},
		},
	})
}

func testAccServerResourceConfigUnsetSetup() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                = "pbs"
  default_queue       = "workq"
  scheduler_iteration = 450
}
`
}

func testAccServerResourceConfigUnset() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = ["default_queue"]
}
`
}

// TestAccServerResource_unsetAttributesRejectsUnknown verifies that an unknown
// attribute name in unset_attributes is rejected with a helpful error.
func TestAccServerResource_unsetAttributesRejectsUnknown(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the existing server.
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			{
				Config:      testAccServerResourceConfigUnsetUnknown(),
				ExpectError: regexp.MustCompile("not a pbs_server attribute"),
			},
		},
	})
}

func testAccServerResourceConfigUnsetUnknown() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = ["does_not_exist"]
}
`
}

// TestAccServerResource_unsetDefaultReturningScalar verifies that unsetting a
// scalar which PBS reports at a default value afterwards (scheduler_iteration
// reverts to 600) keeps Terraform state equal to that real PBS value, does not
// produce a perpetual diff (the one-shot unset is tracked in private state), and
// stays consistent when the unset marker is later removed.
func TestAccServerResource_unsetDefaultReturningScalar(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Set a custom scheduler_iteration.
			{
				Config: testAccServerResourceConfigSchedulerIteration(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "450"),
				),
			},
			// Unset it: PBS reverts it to its default (600), and state records that
			// real value. The post-apply idempotency plan must be empty.
			{
				Config: testAccServerResourceConfigUnsetSchedulerIteration(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "600"),
				),
			},
			// A second plan with the same configuration must remain clean (the
			// one-shot unset is not re-issued).
			{
				Config:   testAccServerResourceConfigUnsetSchedulerIteration(),
				PlanOnly: true,
			},
			// Removing the unset marker keeps the real value and stays consistent.
			{
				Config: testAccServerResourceConfigMinimal(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "scheduler_iteration", "600"),
				),
			},
		},
	})
}

func testAccServerResourceConfigSchedulerIteration() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name                = "pbs"
  scheduler_iteration = 450
}
`
}

func testAccServerResourceConfigUnsetSchedulerIteration() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = ["scheduler_iteration"]
}
`
}

// TestAccServerResource_unsetAclAttribute verifies that an ACL string attribute
// can be unset without leaving an invalid unknown value in state (the ACL
// format-preservation path must skip unknown planned values).
func TestAccServerResource_unsetAclAttribute(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			// Set an ACL user list.
			{
				Config: testAccServerResourceConfigAclUsersSet(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.pbs", "acl_users", "admin,staff"),
				),
			},
			// Unset acl_users; it must be cleared without an unknown-state error.
			{
				Config: testAccServerResourceConfigUnsetAclUsers(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("pbs_server.pbs", "acl_users"),
				),
			},
		},
	})
}

func testAccServerResourceConfigAclUsersSet() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name            = "pbs"
  acl_user_enable = true
  acl_users       = "admin,staff"
}
`
}

func testAccServerResourceConfigUnsetAclUsers() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  acl_user_enable  = true
  unset_attributes = ["acl_users"]
}
`
}

// TestAccServerResource_unsetAttributesRejectsComputed verifies that an
// unset_attributes value that is not known at plan time (an element computed from
// another resource) is rejected with a clear error rather than guessed at.
func TestAccServerResource_unsetAttributesRejectsComputed(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			{
				Config:      testAccServerResourceConfigUnsetComputed(),
				ExpectError: regexp.MustCompile("must be known at plan time"),
			},
		},
	})
}

func testAccServerResourceConfigUnsetComputed() string {
	return providerConfig() + `
resource "terraform_data" "unset" {
  input = "default_queue"
}

resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = [terraform_data.unset.output]
}
`
}

// TestAccServerResource_unsetAttributesRejectsNull verifies that a null element in
// unset_attributes is rejected instead of being treated as a wildcard.
func TestAccServerResource_unsetAttributesRejectsNull(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigMinimal(),
				ResourceName:       "pbs_server.pbs",
				ImportState:        true,
				ImportStateId:      "pbs",
				ImportStatePersist: true,
			},
			{
				Config:      testAccServerResourceConfigUnsetNull(),
				ExpectError: regexp.MustCompile("null"),
			},
		},
	})
}

func testAccServerResourceConfigUnsetNull() string {
	return providerConfig() + `
resource "pbs_server" "pbs" {
  name             = "pbs"
  unset_attributes = [null]
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
