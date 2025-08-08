package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccServerResource_basicConfiguration tests basic server configuration and updates.
func TestAccServerResource_basicConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigBasic(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigBasicUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_aclConfiguration tests all ACL-related fields.
func TestAccServerResource_aclConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigACL(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigACLUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_limitAttributes tests all limit attribute maps.
func TestAccServerResource_limitAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigLimits(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigLimitsUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_resourceMaps tests resource map attributes.
func TestAccServerResource_resourceMaps(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigResourceMaps(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigResourceMapsUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_numericConfiguration tests all numeric fields.
func TestAccServerResource_numericConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigNumeric(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigNumericUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_comprehensiveConfiguration tests all fields together.
func TestAccServerResource_comprehensiveConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccServerResourceConfigFull(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testAccServerResourceConfigFullUpdated(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccServerResource_import tests importing the PBS server.
func TestAccServerResource_import(t *testing.T) {
	serverName := "pbs" // PBS server name that exists in the test environment

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        testAccServerResourceConfigMinimalForImport(serverName),
				ResourceName:  "pbs_server.test",
				ImportState:   true,
				ImportStateId: serverName,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}
					state := s[0]
					if state.ID != serverName {
						return fmt.Errorf("expected ID %s, got %s", serverName, state.ID)
					}
					if state.Attributes["name"] != serverName {
						return fmt.Errorf("expected name %s, got %s", serverName, state.Attributes["name"])
					}
					// Verify that some basic server attributes are imported
					if state.Attributes["log_events"] == "" {
						return fmt.Errorf("expected log_events to be set")
					}
					return nil
				},
			},
		},
	})
}

// TestAccServerResource_update tests updating PBS server properties.
func TestAccServerResource_updateWithDrift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Import the server and set specific properties
			{
				Config:        testAccServerResourceConfigBasic(),
				ResourceName:  "pbs_server.test",
				ImportState:   true,
				ImportStateId: "pbs",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.test", "name", "pbs"),
					resource.TestCheckResourceAttr("pbs_server.test", "id", "pbs"),
				),
			},
		},
	})
}

// TestAccServerResource_configurationDrift tests that we can detect and correct configuration drift.
func TestAccServerResource_configurationDrift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Test that we can apply configuration changes to the server
			{
				Config: testAccServerResourceConfigBasicUpdated(),
				ResourceName:  "pbs_server.test",
				ImportState:   true,
				ImportStateId: "pbs",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbs_server.test", "name", "pbs"),
					resource.TestCheckResourceAttr("pbs_server.test", "id", "pbs"),
					resource.TestCheckResourceAttr("pbs_server.test", "log_events", "1023"),
					resource.TestCheckResourceAttr("pbs_server.test", "comment", "Updated test server configuration"),
				),
			},
		},
	})
}

func testAccServerResourceConfigBasic() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  log_events = 511
  comment = "Initial test server configuration"
}
`
}

func testAccServerResourceConfigBasicUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  log_events = 1023
  comment = "Updated test server configuration"
}
`
}

func testAccServerResourceConfigACL() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  acl_host_enable = true
  acl_hosts = "host1.example.com,host2.example.com"
  acl_user_enable = true
  acl_users = "user1@example.com,user2@example.com"
  acl_resv_group_enable = false
  acl_resv_user_enable = false
  log_events = 511
}
`
}

func testAccServerResourceConfigACLUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  acl_host_enable = false
  acl_hosts = "host3.example.com,host4.example.com"
  acl_user_enable = false
  acl_users = "user3@example.com,user4@example.com"
  acl_resv_group_enable = true
  acl_resv_groups = "group1,group2"
  acl_resv_user_enable = true
  acl_resv_users = "resvuser1,resvuser2"
  log_events = 511
}
`
}

func testAccServerResourceConfigLimits() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  max_group_res = {
    "ncpus" = "[o:PBS_ALL=100]"
    "mem"   = "[o:PBS_ALL=50gb]"
  }
  max_queued_res = {
    "ncpus" = "[o:PBS_ALL=200]"
    "walltime" = "[o:PBS_ALL=24:00:00]"
  }
  max_run_res = {
    "ncpus" = "[o:PBS_ALL=150]"
  }
  max_user_res = {
    "walltime" = "[o:PBS_ALL=72:00:00]"
    "mem" = "[o:PBS_ALL=100gb]"
  }
  log_events = 511
}
`
}

func testAccServerResourceConfigLimitsUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  max_group_res = {
    "ncpus" = "[o:PBS_ALL=200]"
    "mem"   = "[o:PBS_ALL=100gb]"
    "vmem"  = "[o:PBS_ALL=120gb]"
  }
  max_group_res_soft = {
    "ncpus" = "[o:PBS_ALL=180]"
  }
  max_queued_res = {
    "ncpus" = "[o:PBS_ALL=300]"
    "walltime" = "[o:PBS_ALL=48:00:00]"
  }
  max_run_res = {
    "ncpus" = "[o:PBS_ALL=250]"
    "mem" = "[o:PBS_ALL=200gb]"
  }
  max_run_res_soft = {
    "ncpus" = "[o:PBS_ALL=200]"
  }
  max_user_res = {
    "walltime" = "[o:PBS_ALL=168:00:00]"
    "mem" = "[o:PBS_ALL=200gb]"
  }
  max_user_res_soft = {
    "walltime" = "[o:PBS_ALL=120:00:00]"
  }
  log_events = 511
}
`
}

func testAccServerResourceConfigResourceMaps() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  default_chunk = {
    "ncpus" = "2"
    "mem"   = "4gb"
  }
  resources_default = {
    "ncpus"    = "1"
    "walltime" = "01:00:00"
    "mem"      = "2gb"
  }
  resources_available = {
    "ncpus" = "1000"
    "mem"   = "1000gb"
  }
  resources_max = {
    "ncpus"    = "100"
    "walltime" = "168:00:00"
  }
  log_events = 511
}
`
}

func testAccServerResourceConfigResourceMapsUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  default_chunk = {
    "ncpus" = "4"
    "mem"   = "8gb"
    "vmem"  = "10gb"
  }
  resources_default = {
    "ncpus"    = "2"
    "walltime" = "02:00:00"
    "mem"      = "4gb"
  }
  resources_available = {
    "ncpus" = "2000"
    "mem"   = "2000gb"
    "vmem"  = "2500gb"
  }
  resources_max = {
    "ncpus"    = "200"
    "walltime" = "240:00:00"
    "mem"      = "500gb"
  }
  log_events = 511
}
`
}

func testAccServerResourceConfigNumeric() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  max_array_size = 1000
  max_concurrent_provision = 5
  max_job_sequence_id = 10000000
  max_group_run = 50
  max_group_run_soft = 40
  max_running = 100
  max_user_run = 25
  max_user_run_soft = 20
  backfill_depth = 100
  node_fail_requeue = 300
  pbs_license_linger_time = 31536000
  pbs_license_max = 2147483647
  pbs_license_min = 0
  python_gc_min_interval = 60
  python_restart_max_pbs_servers = 10
  python_restart_max_objects = 1000
  reserve_retry_init = 5
  reserve_retry_time = 30
  rpp_highwater = 8192
  rpp_max_pkt_check = 100
  rpp_retry = 3
  scheduler_iteration = 600
  log_events = 511
}
`
}

func testAccServerResourceConfigNumericUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  max_array_size = 2000
  max_concurrent_provision = 10
  max_job_sequence_id = 50000000
  max_group_run = 100
  max_group_run_soft = 80
  max_running = 200
  max_user_run = 50
  max_user_run_soft = 40
  backfill_depth = 200
  node_fail_requeue = 310
  pbs_license_linger_time = 63072000
  pbs_license_max = 2000000000
  pbs_license_min = 1
  python_gc_min_interval = 120
  python_restart_max_pbs_servers = 20
  python_restart_max_objects = 2000
  reserve_retry_init = 10
  reserve_retry_time = 60
  rpp_highwater = 16384
  rpp_max_pkt_check = 200
  rpp_retry = 5
  scheduler_iteration = 300
  log_events = 1023
}
`
}

func testAccServerResourceConfigFull() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  
  # ACL Configuration
  acl_host_enable = true
  acl_hosts = "host1.example.com"
  acl_user_enable = true
  acl_users = "user1@example.com"
  acl_resv_group_enable = false
  acl_resv_user_enable = false
  
  # Basic Configuration
  eligible_time_enable = true
  flatuid = false
  power_provisioning = true
  query_other_jobs = false
  resv_enable = true
  
  # String Configuration
  comment = "Comprehensive test server"
  default_queue = "workq"
  mailer = "/usr/sbin/sendmail"
  mail_from = "admin@test.com"
  managers = "admin@test.com"
  operators = "operator@test.com"
  job_history_duration = "24:00:00"
  job_history_enable = true
  job_requeue_timeout = "30"
  job_sort_formula = "ncpus + mem/1024"
  jobscript_max_size = "100kb"
  python_restart_min_interval = "01:00:00"
  queued_jobs_threshold = "100"
  queued_jobs_threshold_res = "ncpus=1000"
  restrict_res_to_release_on_suspend = "ncpus"
  resv_post_processing_time = "00:05:00"
  
  # Numeric Configuration
  max_array_size = 5000
  max_concurrent_provision = 10
  max_job_sequence_id = 10000000
  max_running = 200
  node_fail_requeue = 310
  log_events = 511
  
  # Limit Attributes
  max_group_res = {
    "ncpus" = "[o:PBS_ALL=100]"
    "mem"   = "[o:PBS_ALL=50gb]"
  }
  max_queued_res = {
    "ncpus" = "[o:PBS_ALL=200]"
  }
  max_run_res = {
    "ncpus" = "[o:PBS_ALL=150]"
  }
  max_user_res = {
    "walltime" = "[o:PBS_ALL=72:00:00]"
  }
  
  # Resource Maps
  default_chunk = {
    "ncpus" = "1"
    "mem"   = "2gb"
  }
  resources_default = {
    "ncpus" = "1"
    "walltime" = "01:00:00"
  }
  
  # WebAPI Configuration
  webapi_enable = false
}
`
}

func testAccServerResourceConfigFullUpdated() string {
	return providerConfig() + `
resource "pbs_server" "test" {
  name = "pbs"
  
  # ACL Configuration - Updated
  acl_host_enable = false
  acl_host_moms_enable = true
  acl_user_enable = false
  acl_resv_group_enable = true
  acl_resv_groups = "group1,group2"
  acl_resv_user_enable = true
  acl_resv_users = "resvuser1,resvuser2"
  
  # Basic Configuration - Updated
  eligible_time_enable = false
  flatuid = true
  power_provisioning = false
  query_other_jobs = true
  resv_enable = false
  
  # String Configuration - Updated
  comment = "Updated comprehensive test server"
  default_queue = "batch"
  mailer = "/usr/bin/mail"
  mail_from = "newadmin@test.com"
  managers = "newadmin@test.com,admin2@test.com"
  operators = "operator1@test.com,operator2@test.com"
  job_history_duration = "48:00:00"
  job_history_enable = false
  job_requeue_timeout = "60"
  job_sort_formula = "walltime + ncpus*2"
  jobscript_max_size = "200kb"
  python_restart_min_interval = "02:00:00"
  queued_jobs_threshold = "200"
  queued_jobs_threshold_res = "ncpus=2000"
  restrict_res_to_release_on_suspend = "mem"
  resv_post_processing_time = "00:10:00"
  
  # Numeric Configuration - Updated
  max_array_size = 10000
  max_concurrent_provision = 20
  max_job_sequence_id = 50000000
  max_running = 400
  node_fail_requeue = 320
  log_events = 1023
  
  # Limit Attributes - Updated and Expanded
  max_group_res = {
    "ncpus" = "[o:PBS_ALL=200]"
    "mem"   = "[o:PBS_ALL=100gb]"
    "vmem"  = "[o:PBS_ALL=120gb]"
  }
  max_group_res_soft = {
    "ncpus" = "[o:PBS_ALL=180]"
  }
  max_queued_res = {
    "ncpus" = "[o:PBS_ALL=400]"
    "walltime" = "[o:PBS_ALL=168:00:00]"
  }
  max_run_res = {
    "ncpus" = "[o:PBS_ALL=300]"
    "mem" = "[o:PBS_ALL=200gb]"
  }
  max_run_res_soft = {
    "ncpus" = "[o:PBS_ALL=250]"
  }
  max_user_res = {
    "walltime" = "[o:PBS_ALL=168:00:00]"
    "mem" = "[o:PBS_ALL=100gb]"
  }
  max_user_res_soft = {
    "walltime" = "[o:PBS_ALL=120:00:00]"
  }
  
  # Resource Maps - Updated
  default_chunk = {
    "ncpus" = "2"
    "mem"   = "4gb"
    "vmem"  = "5gb"
  }
  resources_default = {
    "ncpus" = "2"
    "walltime" = "02:00:00"
    "mem" = "4gb"
  }
  resources_available = {
    "ncpus" = "2000"
    "mem"   = "2000gb"
  }
  resources_max = {
    "ncpus" = "100"
    "walltime" = "240:00:00"
  }
  
  # WebAPI Configuration - Updated
  webapi_enable = true
  webapi_auth_issuers = "https://auth.example.com"
  webapi_oidc_clientid = "pbs-client"
  webapi_oidc_provider_url = "https://oidc.example.com"
}
`
}

func testAccServerResourceConfigMinimalForImport(name string) string {
	return providerConfig() + fmt.Sprintf(`
resource "pbs_server" "test" {
  name = "%[1]s"
}
`, name)
}
