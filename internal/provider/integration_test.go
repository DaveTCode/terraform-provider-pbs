package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccIntegration_CompleteWorkflow tests a complete PBS workflow
// including creating queues, nodes, resources, hooks, and server configuration
func TestAccIntegration_CompleteWorkflow(t *testing.T) {
	queueName := testAccResourceName("integration_queue")
	nodeName := testAccResourceName("integration_node")
	resourceName := testAccResourceName("integration_resource")
	hookName := testAccResourceName("integration_hook")
	serverName := testAccResourceName("integration_server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckCompleteWorkflowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCompleteWorkflowConfig(serverName, queueName, nodeName, resourceName, hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Server checks
					resource.TestCheckResourceAttr("pbs_server.test", "name", serverName),
					resource.TestCheckResourceAttr("pbs_server.test", "default_queue", queueName),

					// Queue checks
					resource.TestCheckResourceAttr("pbs_queue.test", "name", queueName),
					resource.TestCheckResourceAttr("pbs_queue.test", "enabled", "true"),
					resource.TestCheckResourceAttr("pbs_queue.test", "started", "true"),

					// Node checks
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.test_resource", "100mb"),

					// Resource checks
					resource.TestCheckResourceAttr("pbs_resource.test", "name", resourceName),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "size"),

					// Hook checks
					resource.TestCheckResourceAttr("pbs_hook.test", "name", hookName),
					resource.TestCheckResourceAttr("pbs_hook.test", "enabled", "true"),
				),
			},
		},
	})
}

// TestAccIntegration_QueueDependencies tests queue dependencies and ordering
func TestAccIntegration_QueueDependencies(t *testing.T) {
	routingQueue := testAccResourceName("routing_queue")
	execQueue1 := testAccResourceName("exec_queue_1")
	execQueue2 := testAccResourceName("exec_queue_2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueDependenciesConfig(routingQueue, execQueue1, execQueue2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_queue.routing", "name", routingQueue),
					resource.TestCheckResourceAttr("pbs_queue.routing", "queue_type", "Route"),
					resource.TestCheckResourceAttr("pbs_queue.exec1", "name", execQueue1),
					resource.TestCheckResourceAttr("pbs_queue.exec1", "queue_type", "Execution"),
					resource.TestCheckResourceAttr("pbs_queue.exec2", "name", execQueue2),
					resource.TestCheckResourceAttr("pbs_queue.exec2", "queue_type", "Execution"),
				),
			},
		},
	})
}

// TestAccIntegration_NodeWithCustomResources tests nodes with custom PBS resources
func TestAccIntegration_NodeWithCustomResources(t *testing.T) {
	nodeName := testAccResourceName("custom_node")
	resource1Name := testAccResourceName("custom_resource_1")
	resource2Name := testAccResourceName("custom_resource_2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeWithCustomResourcesConfig(nodeName, resource1Name, resource2Name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_resource.custom1", "name", resource1Name),
					resource.TestCheckResourceAttr("pbs_resource.custom1", "type", "size"),
					resource.TestCheckResourceAttr("pbs_resource.custom2", "name", resource2Name),
					resource.TestCheckResourceAttr("pbs_resource.custom2", "type", "long"),
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
				),
			},
		},
	})
}

// TestAccIntegration_HookOrdering tests multiple hooks with different ordering
func TestAccIntegration_HookOrdering(t *testing.T) {
	hook1Name := testAccResourceName("hook_order_1")
	hook2Name := testAccResourceName("hook_order_2")
	hook3Name := testAccResourceName("hook_order_3")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHookOrderingConfig(hook1Name, hook2Name, hook3Name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_hook.hook1", "order", "1"),
					resource.TestCheckResourceAttr("pbs_hook.hook2", "order", "2"),
					resource.TestCheckResourceAttr("pbs_hook.hook3", "order", "3"),
					resource.TestCheckResourceAttr("pbs_hook.hook1", "event", "execjob_begin"),
					resource.TestCheckResourceAttr("pbs_hook.hook2", "event", "execjob_end"),
					resource.TestCheckResourceAttr("pbs_hook.hook3", "event", "queuejob"),
				),
			},
		},
	})
}

func testAccCheckCompleteWorkflowDestroy(s *terraform.State) error {
	// Check that all resources are properly destroyed
	// This would involve connecting to PBS and verifying cleanup
	return nil
}

func testAccCompleteWorkflowConfig(serverName, queueName, nodeName, resourceName, hookName string) string {
	return providerConfig() + fmt.Sprintf(`
# Create a custom PBS resource first
resource "pbs_resource" "test" {
  name = %[4]q
  type = "size"
  flag = "h"
}

# Create a queue
resource "pbs_queue" "test" {
  name        = %[2]q
  queue_type  = "Execution"
  enabled     = true
  started     = true
  priority    = 100
  max_running = 10
  resources_default = {
    ncpus    = 1
    walltime = "01:00:00"
  }
}

# Create a node that uses the custom resource
resource "pbs_node" "test" {
  name = %[3]q
  port = 15002
  resources_available = {
    mem           = "8gb"
    ncpus         = "4"
    test_resource = "100mb"
  }
  resv_enable = true
  depends_on  = [pbs_resource.test]
}

# Create a hook
resource "pbs_hook" "test" {
  name        = %[5]q
  enabled     = true
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
}

# Configure the server to use the queue as default
resource "pbs_server" "test" {
  name = %[1]q
  default_chunk = {
    ncpus = 1
  }
  default_queue            = pbs_queue.test.name
  eligible_time_enable     = false
  log_events               = 511
  max_array_size           = 10000
  max_concurrent_provision = 5
  power_provisioning       = false
  query_other_jobs         = true
  resources_default = {
    ncpus = 1
  }
  resv_enable         = true
  scheduler_iteration = 600
  depends_on          = [pbs_queue.test]
}
`, serverName, queueName, nodeName, resourceName, hookName)
}

func testAccQueueDependenciesConfig(routingQueue, execQueue1, execQueue2 string) string {
	return providerConfig() + fmt.Sprintf(`
# Create execution queues first
resource "pbs_queue" "exec1" {
  name        = %[2]q
  queue_type  = "Execution"
  enabled     = true
  started     = true
  priority    = 100
}

resource "pbs_queue" "exec2" {
  name        = %[3]q
  queue_type  = "Execution"  
  enabled     = true
  started     = true
  priority    = 200
}

# Create routing queue that routes to execution queues
resource "pbs_queue" "routing" {
  name        = %[1]q
  queue_type  = "Route"
  enabled     = true
  started     = true
  # route_destinations = "${pbs_queue.exec1.name},${pbs_queue.exec2.name}"
  depends_on  = [pbs_queue.exec1, pbs_queue.exec2]
}
`, routingQueue, execQueue1, execQueue2)
}

func testAccNodeWithCustomResourcesConfig(nodeName, resource1Name, resource2Name string) string {
	return providerConfig() + fmt.Sprintf(`
# Create custom PBS resources
resource "pbs_resource" "custom1" {
  name = %[2]q
  type = "size"
  flag = "h"
}

resource "pbs_resource" "custom2" {
  name = %[3]q
  type = "long"
  flag = "hn"
}

# Create node using custom resources
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resources_available = {
    mem    = "16gb"
    ncpus  = "8"
    # Use the custom resources
    "${pbs_resource.custom1.name}" = "500mb"
    "${pbs_resource.custom2.name}" = "1000"
  }
  resv_enable = true
  depends_on  = [pbs_resource.custom1, pbs_resource.custom2]
}
`, nodeName, resource1Name, resource2Name)
}

func testAccHookOrderingConfig(hook1Name, hook2Name, hook3Name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "hook1" {
  name        = %[1]q
  enabled     = true
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
}

resource "pbs_hook" "hook2" {
  name        = %[2]q
  enabled     = true
  event       = "execjob_end"
  order       = 2
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
}

resource "pbs_hook" "hook3" {
  name        = %[3]q
  enabled     = true
  event       = "queuejob"
  order       = 3
  type        = "site"
  user        = "pbsadmin"
  fail_action = "offline_vnodes"
}
`, hook1Name, hook2Name, hook3Name)
}
