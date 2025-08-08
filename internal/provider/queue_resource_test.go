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
			// Skip import test for now - we'll test it separately
			// {
			//     ResourceName:      "pbs_queue.test",
			//     ImportState:       true,
			//     ImportStateVerify: true,
			// },
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

// TestAccQueueResource_import tests importing an existing queue
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
					resource.TestCheckResourceAttr("pbs_queue.test", "max_running", "50"),
					resource.TestCheckResourceAttr("pbs_queue.test", "max_queued_res.ncpus", "[o:PBS_ALL=100]"),
				),
			},
		},
	})
}

func testAccCheckQueueExists(resourceName string) resource.TestCheckFunc {
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
  max_running = 50
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
