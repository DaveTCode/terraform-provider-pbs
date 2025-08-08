package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccQueueDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccQueueDataSourceConfig("workq"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_queue.test", "name", "workq"),
					resource.TestCheckResourceAttr("data.pbs_queue.test", "queue_type", "Execution"),
					resource.TestCheckResourceAttrSet("data.pbs_queue.test", "enabled"),
					resource.TestCheckResourceAttrSet("data.pbs_queue.test", "started"),
				),
			},
		},
	})
}

func TestAccServerDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfig("pbs"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_server.test", "name", "pbs"),
					resource.TestCheckResourceAttrSet("data.pbs_server.test", "log_events"),
				),
			},
		},
	})
}

func TestAccNodeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeDataSourceConfig("pbs"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_node.test", "name", "pbs"),
					resource.TestCheckResourceAttrSet("data.pbs_node.test", "port"),
				),
			},
		},
	})
}

func TestAccPbsResourceDataSource_basic(t *testing.T) {
	// First create a resource to query
	resourceName := testAccResourceName("test_data_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceDataSourceConfig(resourceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_resource.test", "name", resourceName),
					resource.TestCheckResourceAttr("data.pbs_resource.test", "type", "size"),
					resource.TestCheckResourceAttr("data.pbs_resource.test", "flag", "h"),
				),
			},
		},
	})
}

func TestAccHookDataSource_basic(t *testing.T) {
	// First create a hook to query
	hookName := testAccResourceName("test_data_hook")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHookDataSourceConfig(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pbs_hook.test", "name", hookName),
					resource.TestCheckResourceAttr("data.pbs_hook.test", "event", "execjob_begin"),
					resource.TestCheckResourceAttr("data.pbs_hook.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.pbs_hook.test", "alarm", "30"),
					resource.TestCheckResourceAttr("data.pbs_hook.test", "debug", "false"),
				),
			},
		},
	})
}

func testAccQueueDataSourceConfig(name string) string {
	return providerConfig() + fmt.Sprintf(`
data "pbs_queue" "test" {
  name = %[1]q
}
`, name)
}

func testAccServerDataSourceConfig(name string) string {
	return providerConfig() + fmt.Sprintf(`
data "pbs_server" "test" {
  name = %[1]q
}
`, name)
}

func testAccNodeDataSourceConfig(name string) string {
	return providerConfig() + fmt.Sprintf(`
data "pbs_node" "test" {
  name = %[1]q
}
`, name)
}

func testAccPbsResourceDataSourceConfig(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_resource" "test" {
  name = %[1]q
  type = "size"
  flag = "h"
}

data "pbs_resource" "test" {
  name = pbs_resource.test.name
  depends_on = [pbs_resource.test]
}
`, name)
}

func testAccHookDataSourceConfig(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "test" {
  name        = %[1]q
  enabled     = true
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
  alarm       = 30
  debug       = false
}

data "pbs_hook" "test" {
  name = pbs_hook.test.name
  depends_on = [pbs_hook.test]
}
`, name)
}
