package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccHookResource_basic(t *testing.T) {
	hookName := testAccResourceName("test_hook")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccHookResourceConfig(hookName, "execjob_begin", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "name", hookName),
					resource.TestCheckResourceAttr("pbs_hook.test", "event", "execjob_begin"),
					resource.TestCheckResourceAttr("pbs_hook.test", "order", "1"),
					resource.TestCheckResourceAttr("pbs_hook.test", "enabled", "true"),
					resource.TestCheckResourceAttr("pbs_hook.test", "alarm", "30"),
					resource.TestCheckResourceAttr("pbs_hook.test", "debug", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "pbs_hook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccHookResourceConfig(hookName, "execjob_end", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "event", "execjob_end"),
					resource.TestCheckResourceAttr("pbs_hook.test", "order", "2"),
				),
			},
		},
	})
}

func TestAccHookResource_withDebugAndAlarm(t *testing.T) {
	hookName := testAccResourceName("test_hook_debug")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHookResourceConfigWithDebugAndAlarm(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "debug", "true"),
					resource.TestCheckResourceAttr("pbs_hook.test", "alarm", "60"),
					resource.TestCheckResourceAttr("pbs_hook.test", "fail_action", "offline_vnodes"),
				),
			},
		},
	})
}

func TestAccHookResource_multipleEvents(t *testing.T) {
	hookName := testAccResourceName("test_hook_multi")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHookResourceConfigMultipleEvents(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "event", "execjob_begin,execjob_end"),
				),
			},
		},
	})
}

func TestAccHookResource_disabledState(t *testing.T) {
	hookName := testAccResourceName("test_hook_disabled")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHookResourceConfigDisabled(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "enabled", "false"),
				),
			},
			{
				Config: testAccHookResourceConfigEnabled(hookName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists("pbs_hook.test"),
					resource.TestCheckResourceAttr("pbs_hook.test", "enabled", "true"),
				),
			},
		},
	})
}

// TestAccHookResource_import tests importing an existing hook
func TestAccHookResource_import(t *testing.T) {
	hookName := "test" // Use pre-created hook from setup script

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import testing - test importing the pre-existing test hook
			{
				Config:        testAccHookResourceConfigMinimalForImport(hookName),
				ResourceName:  "pbs_hook.test",
				ImportState:   true,
				ImportStateId: hookName,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					state := s[0]
					if state.ID != hookName {
						return fmt.Errorf("expected ID %s, got %s", hookName, state.ID)
					}

					if state.Attributes["name"] != hookName {
						return fmt.Errorf("expected name %s, got %s", hookName, state.Attributes["name"])
					}

					// Verify that event is set (since it's required)
					if event := state.Attributes["event"]; event == "" {
						return fmt.Errorf("expected event to be set")
					}

					// Verify enabled and type are set (since they're required)
					if enabled := state.Attributes["enabled"]; enabled == "" {
						return fmt.Errorf("expected enabled to be set")
					}

					if hookType := state.Attributes["type"]; hookType == "" {
						return fmt.Errorf("expected type to be set")
					}

					// Verify user is set (since it's required)
					if user := state.Attributes["user"]; user == "" {
						return fmt.Errorf("expected user to be set")
					}

					return nil
				},
			},
		},
	})
}

func testAccCheckHookExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Hook ID is set")
		}

		// TODO: Add actual PBS connection check here

		return nil
	}
}

func testAccCheckHookDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pbs_hook" {
			continue
		}

		// TODO: Add actual PBS connection check here
	}

	return nil
}

func testAccHookResourceConfig(name, event string, order int) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "test" {
  name        = %[1]q
  enabled     = true
  event       = %[2]q
  order       = %[3]d
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
  alarm       = 30
  debug       = false
}
`, name, event, order)
}

func testAccHookResourceConfigWithDebugAndAlarm(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "test" {
  name        = %[1]q
  enabled     = true
  debug       = true
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  alarm       = 60
  fail_action = "offline_vnodes"
}
`, name)
}

func testAccHookResourceConfigMultipleEvents(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "test" {
  name        = %[1]q
  enabled     = true
  event       = "execjob_begin,execjob_end"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
  alarm       = 30
  debug       = false
}
`, name)
}

func testAccHookResourceConfigDisabled(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_hook" "test" {
  name        = %[1]q
  enabled     = false
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  user        = "pbsadmin"
  fail_action = "none"
  alarm       = 30
  debug       = false
}
`, name)
}

func testAccHookResourceConfigEnabled(name string) string {
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
`, name)
}

func testAccHookResourceConfigMinimalForImport(name string) string {
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
`, name)
}
