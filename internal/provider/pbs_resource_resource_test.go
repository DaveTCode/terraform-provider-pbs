package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccPbsResourceResource_basic(t *testing.T) {
	resourceName := testAccResourceName("test_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "size", "h"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "name", resourceName),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "size"),
					resource.TestCheckResourceAttr("pbs_resource.test", "flag", "h"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "pbs_resource.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "size", "hf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "flag", "hf"),
				),
			},
		},
	})
}

func TestAccPbsResourceResource_stringType(t *testing.T) {
	resourceName := testAccResourceName("test_string_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "string", "h"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "string"),
				),
			},
		},
	})
}

func TestAccPbsResourceResource_floatType(t *testing.T) {
	resourceName := testAccResourceName("test_float_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "float", "hn"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "float"),
					resource.TestCheckResourceAttr("pbs_resource.test", "flag", "hn"),
				),
			},
		},
	})
}

func TestAccPbsResourceResource_longType(t *testing.T) {
	resourceName := testAccResourceName("test_long_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "long", "h"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "long"),
				),
			},
		},
	})
}

func TestAccPbsResourceResource_booleanType(t *testing.T) {
	resourceName := testAccResourceName("test_boolean_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "boolean", "h"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "boolean"),
				),
			},
		},
	})
}

func TestAccPbsResourceResource_stringArrayType(t *testing.T) {
	resourceName := testAccResourceName("test_string_array_resource")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckPbsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPbsResourceResourceConfig(resourceName, "string_array", "h"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPbsResourceExists("pbs_resource.test"),
					resource.TestCheckResourceAttr("pbs_resource.test", "type", "string_array"),
				),
			},
		},
	})
}

// TestAccPbsResourceResource_import tests importing an existing PBS resource
func TestAccPbsResourceResource_import(t *testing.T) {
	resourceName := "test" // Use pre-created resource from setup script

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import testing - test importing the pre-existing test resource
			{
				Config:        testAccPbsResourceResourceConfigMinimalForImport(resourceName),
				ResourceName:  "pbs_resource.test",
				ImportState:   true,
				ImportStateId: resourceName,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != resourceName {
						return fmt.Errorf("expected ID %s, got %s", resourceName, state.ID)
					}

					if state.Attributes["name"] != resourceName {
						return fmt.Errorf("expected name %s, got %s", resourceName, state.Attributes["name"])
					}

					// Verify that type is set (since it's required)
					if resourceType := state.Attributes["type"]; resourceType == "" {
						return fmt.Errorf("expected type to be set")
					}

					// Verify that flag is set (since it's required)
					if flag := state.Attributes["flag"]; flag == "" {
						return fmt.Errorf("expected flag to be set")
					}

					return nil
				},
			},
		},
	})
}

func testAccCheckPbsResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No PBS Resource ID is set")
		}

		// TODO: Add actual PBS connection check here

		return nil
	}
}

func testAccCheckPbsResourceDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pbs_resource" {
			continue
		}

		// TODO: Add actual PBS connection check here
	}

	return nil
}

func testAccPbsResourceResourceConfig(name, resourceType, flag string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_resource" "test" {
  name = %[1]q
  type = %[2]q
  flag = %[3]q
}
`, name, resourceType, flag)
}

func testAccPbsResourceResourceConfigMinimalForImport(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_resource" "test" {
  name = %[1]q
  type = "size"
  flag = "h"
}
`, name)
}
