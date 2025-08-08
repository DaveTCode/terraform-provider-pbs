package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Available node names that correspond to Docker containers
var availableTestNodes = []string{
	"compute1", "compute2", "compute3", "node1", "node2",
}

// getTestNodeName returns a specific node name for each test to avoid conflicts
func getTestNodeName(testName string) string {
	testNodeMap := map[string]string{
		"basic":                "compute1",
		"withResources":        "compute2",
		"powerAndProvisioning": "compute3",
		"comprehensive":        "node1",
		"minimal":              "node2",
	}

	if nodeName, exists := testNodeMap[testName]; exists {
		return nodeName
	}
	// Fallback to first available node
	return availableTestNodes[0]
}

// TestAccNodeResource_basic tests basic node creation and updates
func TestAccNodeResource_basic(t *testing.T) {
	nodeName := getTestNodeName("basic")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckNodeDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccNodeResourceConfigBasic(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
					resource.TestCheckResourceAttr("pbs_node.test", "port", "15002"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "comment", "Basic test node"),
				),
			},
			// Update and Read testing
			{
				Config: testAccNodeResourceConfigBasicUpdated(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "port", "15003"),
					resource.TestCheckResourceAttr("pbs_node.test", "comment", "Updated basic test node"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "false"),
				),
			},
		},
	})
}

// TestAccNodeResource_import tests importing an existing node
func TestAccNodeResource_import(t *testing.T) {
	nodeName := "pbs" // Use pre-created node from setup script (server's own hostname)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import testing - test importing the pre-existing PBS server node
			{
				Config:        testAccNodeResourceConfigMinimalForImport(nodeName),
				ResourceName:  "pbs_node.test",
				ImportState:   true,
				ImportStateId: nodeName,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != nodeName {
						return fmt.Errorf("expected ID %s, got %s", nodeName, state.ID)
					}

					if state.Attributes["name"] != nodeName {
						return fmt.Errorf("expected name %s, got %s", nodeName, state.Attributes["name"])
					}

					// Verify that resv_enable is set (since it's required)
					if resvEnable := state.Attributes["resv_enable"]; resvEnable == "" {
						return fmt.Errorf("expected resv_enable to be set")
					}

					return nil
				},
			},
		},
	})
}

// TestAccNodeResource_withResources tests node with custom resource specifications
func TestAccNodeResource_withResources(t *testing.T) {
	nodeName := getTestNodeName("withResources")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeResourceConfigWithResources(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.mem", "8gb"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.ncpus", "4"),
				),
			},
			{
				Config: testAccNodeResourceConfigWithResourcesUpdated(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.mem", "16gb"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.ncpus", "8"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.vmem", "18gb"),
				),
			},
		},
	})
}

// TestAccNodeResource_powerAndProvisioning tests power and provisioning fields
func TestAccNodeResource_powerAndProvisioning(t *testing.T) {
	nodeName := getTestNodeName("powerAndProvisioning")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeResourceConfigPowerProvisioning(nodeName, true, true, true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "poweroff_eligible", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "power_provisioning", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "provision_enable", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "true"),
				),
			},
			{
				Config: testAccNodeResourceConfigPowerProvisioning(nodeName, false, false, false, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "poweroff_eligible", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "power_provisioning", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "provision_enable", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "false"),
				),
			},
		},
	})
}

// TestAccNodeResource_comprehensive tests all supported configurable fields
func TestAccNodeResource_comprehensive(t *testing.T) {
	nodeName := getTestNodeName("comprehensive")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeResourceConfigComprehensive(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
					resource.TestCheckResourceAttr("pbs_node.test", "comment", "Comprehensive test node"),
					resource.TestCheckResourceAttr("pbs_node.test", "port", "15004"),
					resource.TestCheckResourceAttr("pbs_node.test", "priority", "100"),
					resource.TestCheckResourceAttr("pbs_node.test", "poweroff_eligible", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "power_provisioning", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "provision_enable", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.ncpus", "12"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.mem", "32gb"),
				),
			},
			{
				Config: testAccNodeResourceConfigComprehensiveUpdated(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "comment", "Updated comprehensive test node"),
					resource.TestCheckResourceAttr("pbs_node.test", "port", "15005"),
					resource.TestCheckResourceAttr("pbs_node.test", "priority", "200"),
					resource.TestCheckResourceAttr("pbs_node.test", "poweroff_eligible", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "power_provisioning", "true"),
					resource.TestCheckResourceAttr("pbs_node.test", "provision_enable", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "false"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.ncpus", "24"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.mem", "64gb"),
					resource.TestCheckResourceAttr("pbs_node.test", "resources_available.vmem", "72gb"),
				),
			},
		},
	})
}

// TestAccNodeResource_minimalSupported tests only the most basic supported attributes
func TestAccNodeResource_minimalSupported(t *testing.T) {
	nodeName := getTestNodeName("minimal")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeResourceConfigMinimal(nodeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNodeExists("pbs_node.test"),
					resource.TestCheckResourceAttr("pbs_node.test", "name", nodeName),
					resource.TestCheckResourceAttr("pbs_node.test", "port", "15002"),
					resource.TestCheckResourceAttr("pbs_node.test", "resv_enable", "true"),
				),
			},
		},
	})
}

func testAccCheckNodeExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Node ID is set")
		}

		// TODO: Add actual PBS connection check here

		return nil
	}
}

func testAccCheckNodeDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pbs_node" {
			continue
		}

		// TODO: Add actual PBS connection check here
	}

	return nil
}

func testAccNodeResourceConfigBasic(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resv_enable = true
  comment = "Basic test node"
}
`, name)
}

func testAccNodeResourceConfigBasicUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15003
  resv_enable = false
  comment = "Updated basic test node"
}
`, name)
}

func testAccNodeResourceConfigMinimalForImport(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  resv_enable = true
}
`, name)
}

func testAccNodeResourceConfigWithResources(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resv_enable = true
  resources_available = {
    mem   = "8gb"
    ncpus = "4"
  }
}
`, name)
}

func testAccNodeResourceConfigWithResourcesUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resv_enable = true
  resources_available = {
    mem   = "16gb"
    ncpus = "8"
    vmem  = "18gb"
  }
}
`, name)
}

func testAccNodeResourceConfigPowerProvisioning(name string, powerOff, powerProv, provisionEnable, resvEnable bool) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resv_enable = %[5]t
  poweroff_eligible = %[2]t
  power_provisioning = %[3]t
  provision_enable = %[4]t
}
`, name, powerOff, powerProv, provisionEnable, resvEnable)
}

func testAccNodeResourceConfigComprehensive(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  comment = "Comprehensive test node"
  port = 15004
  priority = 100
  poweroff_eligible = true
  power_provisioning = false
  provision_enable = true
  resv_enable = true
  resources_available = {
    ncpus = "12"
    mem = "32gb"
    vmem = "36gb"
  }
}
`, name)
}

func testAccNodeResourceConfigComprehensiveUpdated(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  comment = "Updated comprehensive test node"
  port = 15005
  priority = 200
  poweroff_eligible = false
  power_provisioning = true
  provision_enable = false
  resv_enable = false
  resources_available = {
    ncpus = "24"
    mem = "64gb"
    vmem = "72gb"
  }
}
`, name)
}

func testAccNodeResourceConfigMinimal(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_node" "test" {
  name = %[1]q
  port = 15002
  resv_enable = true
}
`, name)
}
