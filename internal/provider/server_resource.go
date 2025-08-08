package provider

import (
	"context"
	"fmt"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &serverResource{}
	_ resource.ResourceWithConfigure   = &serverResource{}
	_ resource.ResourceWithImportState = &serverResource{}
)

func NewServerResource() resource.Resource {
	return &serverResource{}
}

type serverResource struct {
	client *pbsclient.PbsClient
}

func (r *serverResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *serverResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{ // TODO - How to avoid duplication of this schema with data source?
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this server. This is the same as the name.",
			},
			"acl_host_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Specifies whether the server obeys the host access control list in the acl_hosts server attribute.",
			},
			"acl_host_moms_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether all MoMs are automatically allowed to contact the server with the same privilege as hosts listed in the acl_hosts server attribute.",
			},
			"acl_hosts": schema.StringAttribute{
				Optional:    true,
				Description: " List of hosts from which services can be requested of this server. Requests from the server host are always honored whether or not that host is in the list.  This list contains the fully qualified domain names of the hosts. List is evaluated left-to-right; first match in list is used.",
			},
			"acl_resv_group_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Specifies whether the server obeys the group reservation access control list in the acl_resv_groups server attribute.",
			},
			"acl_resv_groups": schema.StringAttribute{
				Optional:    true,
				Description: " List of groups allowed or denied permission to create reservations in this PBS complex.  The groups in the list are groups on the server host, not submission hosts.    List is evaluated left-to-right; first match in list is used.",
			},
			"acl_resv_host_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether the server obeys the host reservation access control list in the acl_resv_hosts server attribute.",
			},
			"acl_resv_hosts": schema.StringAttribute{
				Optional:    true,
				Description: " List of hosts from which reservations can be created in this PBS complex. This list is made up  of the fully-qualified domain names of the hosts.  List is evaluated left-to-right; first match in list is used.",
			},
			"acl_resv_user_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Specifies whether the server limits which  users are allowed to create reservations, according to the access control list in the acl_resv_users server attribute.",
			},
			"acl_resv_users": schema.StringAttribute{
				Optional:    true,
				Description: " List of users allowed or denied permission to create reservations  in this PBS complex.   List is evaluated left-to-right; first match in list is used.",
			},
			"acl_roots": schema.StringAttribute{
				Optional:    true,
				Description: " List of users with root privilege who can submit and run jobs in this PBS complex.  For any job whose owner is root or Administrator, the job owner must be listed in this access control list, or the job is rejected.  List is evaluated left-to-right; first match in list is used. ",
			},
			"acl_user_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether the server limits which users are allowed to run commands at the server, according to the control list in the acl_users server attribute.",
			},
			"acl_users": schema.StringAttribute{
				Optional:    true,
				Description: " List of users allowed or denied permission to run commands at this server.   List is evaluated left-to-right; first match in list is used.",
			},
			"backfill_depth": schema.Int32Attribute{
				Optional:    true,
				Description: "Specifies backfilling behavior.  Sets the number of jobs that are to be backfilled around.  Overridden by backfill_depth queue attribute. Recommendation: set this to less than 100.",
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Description: "Informational text.  Can be set by a scheduler or other privileged client.",
			},
			"default_chunk": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "The list of resources which will be inserted into each chunk of a job's select specification if the corresponding resource is not specified by the user. This provides a means for a site to be sure a given resource is properly accounted for even if not specified by the user. ",
			},
			"default_qdel_arguments": schema.StringAttribute{
				Optional:    true,
				Description: " Argument to qdel command.  Automatically added to all qdel commands.  See qdel(1B).  Overrides standard defaults. Overridden by arguments given on the command line.",
			},
			"default_qsub_arguments": schema.StringAttribute{
				Optional:    true,
				Description: " Arguments that are automatically added to the qsub command.  Any valid arguments to qsub command, such as job attributes. Setting a job attribute via default_qsub_arguments sets that attribute for each job which does not explicitly override it. See qsub(1B). Settable by the administrator via the qmgr command. Overrides standard defaults. Overridden by arguments given on the command line and in script directives. ",
			},
			"default_queue": schema.StringAttribute{
				Optional:    true,
				Description: "The name of the default target queue.  Used for requests that do not specify a queue name.  Must be set to an existing queue.",
			},
			"eligible_time_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Enables accruing job wait time in the job's eligible_time attribute.",
			},
			"elim_on_subjobs": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether the server max_queued limit attribute counts each array job as a single job, or counts each subjob as a single job.",
			},
			"flatuid": schema.BoolAttribute{
				Optional:    true,
				Description: " Used for authorization allowing users to submit and alter jobs.   Specifies whether user names are treated as being the same across the PBS server and all submission hosts in the PBS complex.   Can be used to allow users without accounts at the server host to submit jobs. If UserA has an account at the server host, PBS requires that UserA@<server host> is the same as UserA@<execution host>.",
			},
			"job_history_duration": schema.StringAttribute{
				Optional:    true,
				Description: "The length of time PBS will keep each job's history.",
			},
			"job_history_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Enables job history management. Setting  this attribute to True enables job history management.",
			},
			"job_requeue_timeout": schema.StringAttribute{
				Optional:    true,
				Description: "The amount of time that can be taken while requeueing a job.  Minimum allowed value: 1 second.  Maximum allowed value: 3 hours.",
			},
			"job_sort_formula": schema.StringAttribute{
				Optional:    true,
				Description: " Formula for computing job priorities. Described in the PBS Professional Administrator's Guide. If the attribute job_sort_formula is set, all schedulers use the formula in it to compute job priorities.  When this scheduler sorts jobs according to the formula, it computes a priority for each job, where that priority is the value produced by the formula. Jobs with a higher value get higher priority.  ",
			},
			"jobscript_max_size": schema.StringAttribute{
				Optional:    true,
				Description: "Limit on the size of any job script.	",
			},
			"log_events": schema.Int32Attribute{
				Optional:    true,
				Description: "The types of events the server logs as an integer representation of the bits",
			},
			"mailer": schema.StringAttribute{
				Optional:    true,
				Description: " Path to mailer to be used by PBS.  This mailer should function similarly to sendmail.",
			},
			"mail_from": schema.StringAttribute{
				Optional:    true,
				Description: "The username from which server-generated mail is sent to users.  Mail is sent to this address upon failover.  ",
			},
			"managers": schema.StringAttribute{
				Optional:    true,
				Description: "List of PBS Managers.  ",
			},
			"max_array_size": schema.Int32Attribute{
				Optional:    true,
				Description: "The maximum number of subjobs allowed in any array job. ",
			},
			"max_concurrent_provision": schema.Int32Attribute{
				Optional:    true,
				Description: "The max_concurrent_provision attribute is the number of vnodes allowed to be in the process of being provisioned.  Cannot be set to zero.  When unset, default value is used.",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"max_group_res": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum amount of the specified resource that any single group may consume in this PBS complex.",
			},
			"max_group_res_soft": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The soft limit for the specified resource that any single group may consume in this complex.  If a group is consuming more than this amount of the specified resource, their jobs are eligible to be preempted by jobs from groups who are not over their soft limit.",
			},
			"max_group_run": schema.Int32Attribute{
				Optional:    true,
				Description: " Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs owned by the users in one group allowed to be running within this complex at one time. ",
			},
			"max_group_run_soft": schema.Int32Attribute{
				Optional:    true,
				Description: " Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs owned by the users in one group allowed to be running in this complex at one time. If a group has more than this number of jobs running, their jobs are eligible to be preempted by jobs from groups who are not over their soft limit. ",
			},
			"max_job_sequence_id": schema.Int64Attribute{
				Optional:    true,
				Description: " Maximum value of sequence number in a job ID, job array ID, or reservation ID. Minimum allowed is 9999999.  Maximum allowed is 999999999999.  After specified maximum for sequence number has been reached, job IDs start again at 0.",
				Validators: []validator.Int64{
					int64validator.Between(9999999, 999999999999),
				},
			},
			"max_queued": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum number of jobs allowed to be queued or running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes. The effect of this limit depends on how the elim_on_subjobs attribute is set; when elim_on_subjobs is True (the default), max_queued counts each subjob as a job; when elim_on_subjobs is False, max_queued counts each array job as a single job.",
			},
			"max_queued_res": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum amount of the specified resource allowed to be allocated to jobs queued or running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"max_run": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum number of jobs allowed to be running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"max_run_res": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum amount of the specified resource allowed to be allocated to jobs running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"max_run_res_soft": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  Soft limit on the amount of the specified resource allowed to be allocated to jobs running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"max_run_soft": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  Soft limit on the number of jobs allowed to be running in the complex.  Can be specified for projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"max_running": schema.Int32Attribute{
				Optional:    true,
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs in this complex allowed to be running at any given time.",
			},
			"max_user_res": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The maximum amount of the specified resource that any single user may consume within this complex.",
			},
			"max_user_res_soft": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Limit attribute.  The soft limit on the amount of the specified resource that any single user may consume within this complex.  If a user is consuming more than this amount of the specified resource, their jobs are eligible to be preempted by jobs from users who are not over their soft limit.",
			},
			"max_user_run": schema.Int32Attribute{
				Optional:    true,
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs owned by a single user allowed to be running within this complex at one time.",
			},
			"max_user_run_soft": schema.Int32Attribute{
				Optional:    true,
				Description: " Old limit attribute.  Incompatible with new limit  attributes.  The soft limit on the number of jobs owned by a single user that are allowed to be running within this complex at one time. If a user has more than this number of jobs running, their jobs are eligible to be preempted by jobs from users who are not over their soft limit. ",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The unique name for the PBS server",
			},
			"node_fail_requeue": schema.Int32Attribute{
				Optional:    true,
				Description: "Controls whether running jobs are automatically requeued or are deleted when the primary execution host fails.   Number of seconds to wait after losing contact with Mother Superior before requeueing or deleting jobs. Reverts to default value when server is restarted.",
			},
			"node_group_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether placement sets (which includes node grouping) are enabled.  See node_group_key server attribute.",
			},
			"node_group_key": schema.StringAttribute{
				Optional:    true,
				Description: "Specifies the resources to use for placement sets (node grouping).  Overridden by queue's node_group_key attribute.  See node_group_enable server attribute.",
			},
			"operators": schema.StringAttribute{
				Optional:    true,
				Description: " List of PBS Operators. ",
			},
			"pbs_license_info": schema.StringAttribute{
				Optional:    true,
				Description: "Location of license server(s).",
			},
			"pbs_license_linger_time": schema.Int32Attribute{
				Optional:    true,
				Description: "The number of seconds to keep an unused license, when the number of licenses is above the value given by pbs_license_min.",
			},
			"pbs_license_max": schema.Int32Attribute{
				Optional:    true,
				Description: "Maximum number of licenses to be checked out at any time, i.e maximum number of licenses to keep in the PBS local license pool. Sets a cap on the number of nodes or sockets that can be licensed at one time. ",
			},
			"pbs_license_min": schema.Int32Attribute{
				Optional:    true,
				Description: "Minimum number of nodes or sockets to permanently keep licensed, i.e. the minimum number of licenses to keep in the PBS local license pool. This is the minimum number of licenses to keep checked out. If unset, PBS automatically sets the value to 0. ",
			},
			"power_provisioning": schema.BoolAttribute{
				Optional:    true,
				Description: "Reflects use of power profiles via PBS.  Set by PBS to True when PBS_power hook is enabled.",
			},
			"python_gc_min_interval": schema.Int32Attribute{
				Optional:    true,
				Description: "Specifies interval for Python garbage collection.  For no garbage collection, set this to zero.",
			},
			"python_restart_max_pbs_servers": schema.Int32Attribute{
				Optional:    true,
				Description: " The maximum number of hooks to be serviced before the Python interpreter is restarted.  If this number is exceeded, and the time limit set in python_restart_min_interval has elapsed, the Python interpreter is restarted.",
			},
			"python_restart_max_objects": schema.Int32Attribute{
				Optional:    true,
				Description: " The maximum number of objects to be created before the Python interpreter is restarted.  If this number is exceeded, and the time limit set in python_restart_min_interval has elapsed, the Python interpreter is restarted.",
			},
			"python_restart_min_interval": schema.StringAttribute{
				Optional:    true,
				Description: " The minimum time interval before the Python interpreter is restarted.  If this interval has elapsed, and either the maximum number of hooks to be serviced (set in python_restart_max_hooks) has been exceeded or the maximum number of objects to be created (set in python_restart_max_objects) has been exceeded, the Python interpreter is restarted.",
			},
			"query_other_jobs": schema.BoolAttribute{
				Optional:    true,
				Description: "Controls whether unprivileged users are allowed to select or query the status of jobs owned by other users.  ",
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Optional:    true,
				Description: "Limit attribute.  The maximum number of jobs allowed to be queued in the complex.  Can be specified for  projects, users, groups, or all.  Cannot be used  with old limit attributes.",
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Optional:    true,
				Description: "Limit attribute.  The maximum amount of the specified resource allowed to be allocated to jobs queued in the complex.  Can be specified for  projects, users, groups, or all.  Cannot be used with old limit attributes.",
			},
			"reserve_retry_init": schema.Int32Attribute{
				Optional:           true,
				Description:        "Deprecated. The amount of time after a reservation becomes degraded that PBS waits before attempting to reconfirm the reservation.  When this value is changed, only reservations that become degraded after the change use the new value.  Must be greater than zero.",
				DeprecationMessage: "Deprecated",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"reserve_retry_time": schema.Int32Attribute{
				Optional:    true,
				Description: "The amount of time after a reservation becomes degraded that PBS waits before attempting to reconfirm the reservation, as well as amount of time between attempts to reconfirm degraded reservations.  When this value is changed, PBS uses the new value for any subsequent attempts.  Must be greater than zero.",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"resources_available": schema.MapAttribute{
				Optional:    true,
				Description: "The list of available resources and their values defined on the server.",
				ElementType: types.StringType,
			},
			"resources_default": schema.MapAttribute{
				Optional:    true,
				Description: " The list of default job-wide resource values that are set as limits for jobs in this complex when a) the job does not specify a limit, and b) there is no queue default. The value for a string array, e.g. resources_default.<string array resource>, can contain only one string. For host-level resources, see the default_chunk.<resource name> server attribute.",
				ElementType: types.StringType,
			},
			"resources_max": schema.MapAttribute{
				Optional:    true,
				Description: " The maximum amount of each resource that can be requested by any single job in this complex, if there is not a resources_max value defined for the queue at which the job is targeted.  This attribute functions as a gating value for jobs entering the PBS complex.  ",
				ElementType: types.StringType,
			},
			"restrict_res_to_release_on_suspend": schema.StringAttribute{
				Optional:    true,
				Description: " Comma-separated list of consumable resources to be released when jobs are suspended.  If unset, all consumable resources are released on suspension.",
			},
			"resv_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Specifies whether or not advance and standing reservations can be created in this complex. ",
			},
			"resv_post_processing_time": schema.StringAttribute{
				Optional:    true,
				Description: " The amount of time allowed for reservations to clean up after running jobs. Reservation duration and end time are extended by this amount of time.  Jobs are not allowed to run during the cleanup period.",
			},
			"rpp_highwater": schema.Int32Attribute{
				Optional:    true,
				Description: " The maximum number of messages.",
			},
			"rpp_max_pkt_check": schema.Int32Attribute{
				Optional:    true,
				Description: " Maximum number of TPP messages processed by the main server thread per iteration.",
			},
			"rpp_retry": schema.Int32Attribute{
				Optional:    true,
				Description: " In a fault-tolerant setup (multiple pbs_comms), when the first pbs_comm fails partway through a message, this is number of times TPP tries to use the first pbs_comm.",
			},
			"scheduler_iteration": schema.Int32Attribute{
				Optional:    true,
				Description: " The time between scheduling iterations.",
			},
			"webapi_auth_issuers": schema.StringAttribute{
				Optional:    true,
				Description: "Comma-separated list of accepted JWT token issuers.  Used only when using JWT tokens generated via hpcgentoken.",
			},
			"webapi_enable": schema.BoolAttribute{
				Optional:    true,
				Description: " Enables or disables web API support in PBS",
			},
			"webapi_oidc_clientid": schema.StringAttribute{
				Optional:    true,
				Description: " Used with external OIDC service.  The client identifier generated when registering the application with the OIDC provider.  For validation of OIDC ID tokens passed in http(s) requests.",
			},
			"webapi_oidc_provider_url": schema.StringAttribute{
				Optional:    true,
				Description: "Used with external OIDC service.  URL of the OIDC provider, for example https://accounts.google.com   For validation of OIDC ID tokens passed in http(s) requests.",
			},
		},
	}
}

func (r *serverResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pbsclient.PbsClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *pbsclient.PbsClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *serverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model serverModel
	var server pbsclient.PbsServer
	diags := req.Plan.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pbsServerObj := model.ToPbsServer(ctx)
	resp.Diagnostics.Append(diags...)
	server, err := r.client.CreatePbsServer(pbsServerObj)
	if err != nil {
		resp.Diagnostics.AddError("Error creating server", "Could not create server, unexpected error: "+err.Error())
		return
	}

	_ = createServerModel(server)

	diags = resp.State.Set(ctx, model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var server serverModel

	resp.Diagnostics.Append(req.State.Get(ctx, &server)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	serverName := server.Name.ValueString()
	if serverName == "" && !server.ID.IsNull() {
		serverName = server.ID.ValueString()
	}

	q, err := r.client.GetPbsServer(serverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read servers, got error: %s", err))
		return
	}

	// If the server doesn't exist then remove from state
	if q.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	server = createServerModel(q)

	resp.Diagnostics.Append(resp.State.Set(ctx, &server)...)
}

func (r *serverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data serverModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	server := data.ToPbsServer(ctx)
	_, err := r.client.UpdatePbsServer(server)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var server serverModel

	resp.Diagnostics.Append(req.State.Get(ctx, &server)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeletePbsServer(server.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete server, got error: %s", err))
		return
	}
}

func (r *serverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
