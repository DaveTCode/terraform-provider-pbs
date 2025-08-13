package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccQueueResource_basic(t *testing.T) {
	queueName := testAccResourceName("tq_basic")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccQueueResourceConfig(queueName, "Execution", true, true, 100),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "queue_type", "Execution"),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "100"),
				),
			},
			// Update and Read testing
			{
				Config: testAccQueueResourceConfig(queueName, "Execution", false, false, 200),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccQueueResource_import tests importing an existing queue.
func TestAccQueueResource_import(t *testing.T) {
	queueName := "test" // Use pre-created queue from setup script

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import testing - test importing the pre-existing test queue
			{
				Config:        testAccQueueResourceConfigMinimalForImport(queueName),
				ResourceName:  "pbs_queue.test",
				ImportState:   true,
				ImportStateId: queueName,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != queueName {
						return fmt.Errorf("expected ID %s, got %s", queueName, state.ID)
					}

					if state.Attributes["name"] != queueName {
						return fmt.Errorf("expected name %s, got %s", queueName, state.Attributes["name"])
					}

					// Verify that queue_type is set (since it's required)
					if queueType := state.Attributes["queue_type"]; queueType == "" {
						return fmt.Errorf("expected queue_type to be set")
					}

					// Verify enabled and started are set (since they're required)
					if enabled := state.Attributes["enabled"]; enabled == "" {
						return fmt.Errorf("expected enabled to be set")
					}

					if started := state.Attributes["started"]; started == "" {
						return fmt.Errorf("expected started to be set")
					}

					return nil
				},
			},
		},
	})
}

func TestAccQueueResource_withResources(t *testing.T) {
	queueName := testAccResourceName("tq_res")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigWithResources(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.mem", "4gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.ncpus", "1"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.nodect", "1"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.nodes", "1"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.walltime", "01:30:00"),
					resource.TestCheckResourceAttr("pbs_queue.test", "default_chunk.mem", "300gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "default_chunk.ncpus", "256"),
				),
			},
		},
	})
}

func TestAccQueueResource_maxValues(t *testing.T) {
	queueName := testAccResourceName("tq_max")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigWithMaxValues(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_queued_res.ncpus", "[o:PBS_ALL=100]"),
				),
			},
		},
	})
}

func testAccCheckQueueExists(resourceName string) resource.TestCheckFunc { //nolint:unparam
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Queue ID is set")
		}

		// TODO: Add actual PBS connection check here
		// You would connect to PBS and verify the queue exists

		return nil
	}
}

func testAccCheckQueueDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pbs_queue" {
			continue
		}

		// TODO: Add actual PBS connection check here
		// You would connect to PBS and verify the queue is destroyed
	}

	return nil
}

func testAccQueueResourceConfig(name, queueType string, enabled, started bool, priority int) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name        = %[1]q
  queue_type  = %[2]q
  enabled     = %[3]t
  started     = %[4]t
  priority    = %[5]d
}
`, name, queueType, enabled, started, priority)
}

func testAccQueueResourceConfigWithResources(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name       = %[1]q
  queue_type = "Execution"
  enabled    = true
  started    = true
  resources_default = {
    mem      = "4gb"
    ncpus    = "1"
    nodect   = "1"
    nodes    = "1"
    walltime = "01:30:00"
  }
  default_chunk = {
    mem      = "300gb"
    ncpus    = 256
  }
}
`, name)
}

func testAccQueueResourceConfigWithMaxValues(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name        = %[1]q
  queue_type  = "Execution"
  enabled     = true
  started     = true
  max_queued_res  = {
    ncpus = "[o:PBS_ALL=100]"
  }
}
`, name)
}

func testAccQueueResourceConfigMinimalForImport(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name       = %[1]q
  queue_type = "Execution"
  enabled    = true
  started    = true
}
`, name)
}

// Integration tests for comprehensive queue attributes

// TestAccQueueResource_comprehensive_ACL tests ACL configurations.
func TestAccQueueResource_comprehensive_ACL(t *testing.T) {
	queueName := testAccResourceName("tq_acl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			// Create with ACL settings
			{
				Config: testAccQueueResourceConfigACL(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_user_enable", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_users", "testuser,pbsuser"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_users_normalized", "pbsuser,testuser"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_host_enable", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_hosts", "server2,localhost,server1"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_hosts_normalized", "localhost,server1,server2"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_group_enable", "false"),
				),
			},
			// Update ACL settings
			{
				Config: testAccQueueResourceConfigACLUpdated(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_user_enable", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_host_enable", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_group_enable", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_groups", "staff,admin"),
					resource.TestCheckResourceAttr("pbs_queue.test", "acl_groups_normalized", "admin,staff"),
				),
			},
		},
	})
}

// TestAccQueueResource_comprehensive_ResourceLimits tests resource limit configurations.
func TestAccQueueResource_comprehensive_ResourceLimits(t *testing.T) {
	queueName := testAccResourceName("tq_limits")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			// Create with resource limits
			{
				Config: testAccQueueResourceConfigResourceLimits(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_running", "50"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_run", "10"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_run", "20"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.ncpus", "1"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.mem", "1gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.ncpus", "8"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.mem", "16gb"),
				),
			},
			// Update resource limits
			{
				Config: testAccQueueResourceConfigResourceLimitsUpdated(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_running", "100"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_run", "20"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_run", "40"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.ncpus", "2"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.mem", "2gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.ncpus", "16"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.mem", "32gb"),
				),
			},
		},
	})
}

// TestAccQueueResource_comprehensive_RoutingConfig tests routing configurations.
func TestAccQueueResource_comprehensive_RoutingConfig(t *testing.T) {
	queueName := testAccResourceName("tq_routing")
	destQueueName := testAccResourceName("tq_dest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			// Create destination queue first
			{
				Config: testAccQueueResourceConfigRouting(queueName, destQueueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					testAccCheckQueueExists("pbs_queue.dest"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "queue_type", "Route"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_destinations", destQueueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_held_jobs", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_waiting_jobs", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_lifetime", "300"),
				),
			},
			// Update routing config
			{
				Config: testAccQueueResourceConfigRoutingUpdated(queueName, destQueueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_held_jobs", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_waiting_jobs", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_lifetime", "600"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_retry_time", "120"),
				),
			},
		},
	})
}

// TestAccQueueResource_comprehensive_AdvancedLimits tests advanced limit configurations.
func TestAccQueueResource_comprehensive_AdvancedLimits(t *testing.T) {
	queueName := testAccResourceName("tq_advanced")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			// Create with advanced limits
			{
				Config: testAccQueueResourceConfigAdvancedLimits(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_array_size", "1000"),
					resource.TestCheckResourceAttr("pbs_queue.test", "backfill_depth", "10"),
					resource.TestCheckResourceAttr("pbs_queue.test", "kill_delay", "30"),
					resource.TestCheckResourceAttr("pbs_queue.test", "checkpoint_min", "600"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "150"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_run", "[o:PBS_ALL=2]"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_run_soft", "[o:PBS_ALL=1]"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_queued", "[u:PBS_GENERIC=200]"),
				),
			},
			// Update advanced limits
			{
				Config: testAccQueueResourceConfigAdvancedLimitsUpdated(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_array_size", "2000"),
					resource.TestCheckResourceAttr("pbs_queue.test", "backfill_depth", "20"),
					resource.TestCheckResourceAttr("pbs_queue.test", "kill_delay", "60"),
					resource.TestCheckResourceAttr("pbs_queue.test", "checkpoint_min", "1200"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "200"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_run", "[o:PBS_ALL=2]"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_run_soft", "[o:PBS_ALL=1]"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_queued", "[u:PBS_GENERIC=200]"),
				),
			},
		},
	})
}

// Configuration functions for comprehensive queue tests.

func testAccQueueResourceConfigACL(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name             = "%[1]s"
  queue_type       = "Execution"
  enabled          = true
  started          = true
  acl_user_enable  = true
  acl_users        = "testuser,pbsuser"
  acl_host_enable  = true
  acl_hosts        = "server2,localhost,server1"
  acl_group_enable = false
  priority         = 100
}
`, name)
}

func testAccQueueResourceConfigACLUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name             = "%[1]s"
  queue_type       = "Execution"
  enabled          = true
  started          = true
  acl_user_enable  = false
  acl_host_enable  = false
  acl_group_enable = true
  acl_groups       = "staff,admin"
  priority         = 100
}
`, name)
}

func testAccQueueResourceConfigResourceLimits(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name           = "%[1]s"
  queue_type     = "Execution"
  enabled        = true
  started        = true
  max_running    = 50
  max_user_run   = 10
  max_group_run  = 20
  priority       = 100
  
  resources_default = {
    ncpus = "1"
    mem   = "1gb"
  }
  
  resources_max = {
    ncpus = "8"
    mem   = "16gb"
  }
}
`, name)
}

func testAccQueueResourceConfigResourceLimitsUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name           = "%[1]s"
  queue_type     = "Execution"
  enabled        = true
  started        = true
  max_running    = 100
  max_user_run   = 20
  max_group_run  = 40
  priority       = 150
  
  resources_default = {
    ncpus = "2"
    mem   = "2gb"
  }
  
  resources_max = {
    ncpus = "16"
    mem   = "32gb"
  }
}
`, name)
}

func testAccQueueResourceConfigRouting(name, destName string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "dest" {
  name       = "%[2]s"
  queue_type = "Execution"
  enabled    = true
  started    = true
}

resource "pbs_queue" "test" {
  name                = "%[1]s"
  queue_type          = "Route"
  enabled             = true
  started             = true
  route_destinations  = "%[2]s"
  route_held_jobs     = true
  route_waiting_jobs  = true
  route_lifetime      = 300
  priority            = 100
  
  depends_on = [pbs_queue.dest]
}
`, name, destName)
}

func testAccQueueResourceConfigRoutingUpdated(name, destName string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "dest" {
  name       = "%[2]s"
  queue_type = "Execution"
  enabled    = true
  started    = true
}

resource "pbs_queue" "test" {
  name                = "%[1]s"
  queue_type          = "Route"
  enabled             = true
  started             = true
  route_destinations  = "%[2]s"
  route_held_jobs     = false
  route_waiting_jobs  = false
  route_lifetime      = 600
  route_retry_time    = 120
  priority            = 100
  
  depends_on = [pbs_queue.dest]
}
`, name, destName)
}

func testAccQueueResourceConfigAdvancedLimits(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name             = "%[1]s"
  queue_type       = "Execution"
  enabled          = true
  started          = true
  max_array_size   = 1000
  backfill_depth   = 10
  kill_delay       = 30
  checkpoint_min   = 600
  priority         = 150
	max_run          = "[o:PBS_ALL=2]"
	max_run_soft     = "[o:PBS_ALL=1]"
	max_queued       = "[u:PBS_GENERIC=200]"
}
`, name)
}

func testAccQueueResourceConfigAdvancedLimitsUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name             = "%[1]s"
  queue_type       = "Execution"
  enabled          = true
  started          = true
  max_array_size   = 2000
  backfill_depth   = 20
  kill_delay       = 60
  checkpoint_min   = 1200
  priority         = 200
	max_run          = "[o:PBS_ALL=2]"
	max_run_soft     = "[o:PBS_ALL=1]"
	max_queued       = "[u:PBS_GENERIC=200]"
}
`, name)
}

// TestAccQueueResource_comprehensiveAttributes tests queue with comprehensive attribute coverage.
func TestAccQueueResource_comprehensiveAttributes(t *testing.T) {
	queueName := testAccResourceName("tq_comprehensive")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigComprehensive(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "queue_type", "Execution"),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "150"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_running", "50"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_run", "10"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_run", "20"),
					resource.TestCheckResourceAttr("pbs_queue.test", "checkpoint_min", "900"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.%", "3"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.ncpus", "2"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.mem", "4gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.walltime", "02:00:00"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.%", "3"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.ncpus", "16"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.mem", "128gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.walltime", "168:00:00"),
				),
			},
			// Update test
			{
				Config: testAccQueueResourceConfigComprehensiveUpdated(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "priority", "200"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_running", "100"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_default.ncpus", "4"),
					resource.TestCheckResourceAttr("pbs_queue.test", "resources_max.ncpus", "32"),
				),
			},
		},
	})
}

// TestAccQueueResource_routingQueueSimple tests routing queue specific attributes.
func TestAccQueueResource_routingQueueSimple(t *testing.T) {
	queueName := testAccResourceName("tq_routing_simple")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigRoutingSimple(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "queue_type", "Route"),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "route_destinations", "workq,batch"),
				),
			},
		},
	})
}

// TestAccQueueResource_disabledQueue tests disabled queue state.
func TestAccQueueResource_disabledQueue(t *testing.T) {
	queueName := testAccResourceName("tq_disabled")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigDisabled(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "false"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "false"),
				),
			},
		},
	})
}

// TestAccQueueResource_limitsByUser tests user-specific limit configurations.
func TestAccQueueResource_limitsByUser(t *testing.T) {
	queueName := testAccResourceName("tq_userlimits")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueResourceConfigUserLimits(queueName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckQueueExists("pbs_queue.test"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_res.%", "3"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_res.ncpus", "8"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_res.mem", "32gb"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_user_res.walltime", "12:00:00"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_res.%", "2"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_res.ncpus", "64"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_group_res.mem", "256gb"),
				),
			},
		},
	})
}

func testAccQueueResourceConfigComprehensive(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name           = %[1]q
  queue_type     = "Execution"
  enabled        = true
  started        = true
  priority       = 150
  
  # Limits (using old-style attributes since they're Int32)
  max_running    = 50
  max_user_run   = 10
  max_group_run  = 20
  max_queuable   = 500
  
  # Time limits
  checkpoint_min = 900
  
  # Resource defaults
  resources_default = {
    ncpus    = "2"
    mem      = "4gb"
    walltime = "02:00:00"
  }
  
  # Resource maximums
  resources_max = {
    ncpus    = "16"
    mem      = "128gb"
    walltime = "168:00:00"
  }
}
`, name)
}

func testAccQueueResourceConfigComprehensiveUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name           = %[1]q
  queue_type     = "Execution"
  enabled        = true
  started        = true
  priority       = 200
  
  # Updated limits
  max_running    = 100
  max_user_run   = 20
  max_group_run  = 40
  max_queuable   = 1000
  
  # Time limits
  checkpoint_min = 1200
  
  # Updated resource defaults
  resources_default = {
    ncpus    = "4"
    mem      = "8gb"
    walltime = "04:00:00"
  }
  
  # Updated resource maximums
  resources_max = {
    ncpus    = "32"
    mem      = "256gb"
    walltime = "336:00:00"
  }
}
`, name)
}

func testAccQueueResourceConfigRoutingSimple(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name               = %[1]q
  queue_type         = "Route"
  enabled            = true
  started            = true
  route_destinations = "workq,batch"
}
`, name)
}

func testAccQueueResourceConfigDisabled(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name       = %[1]q
  queue_type = "Execution"
  enabled    = false
  started    = false
  priority   = 50
}
`, name)
}

func testAccQueueResourceConfigUserLimits(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_queue" "test" {
  name       = %[1]q
  queue_type = "Execution"
  enabled    = true
  started    = true
  priority   = 100
  
  # User-specific resource limits
  max_user_res = {
    ncpus    = "8"
    mem      = "32gb"
    walltime = "12:00:00"
  }
  
  # Group-specific resource limits
  max_group_res = {
    ncpus = "64"
    mem   = "256gb"
  }
}
`, name)
}
