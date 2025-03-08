package provider

import (
	"context"
	"fmt"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ resource.Resource                = &queueResource{}
	_ resource.ResourceWithConfigure   = &queueResource{}
	_ resource.ResourceWithImportState = &queueResource{}
)

func NewQueueResource() resource.Resource {
	return &queueResource{}
}

type queueResource struct {
	client *pbsclient.PbsClient
}

func (r *queueResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue"
}

func (r *queueResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{ // TODO - How to avoid duplication of this schema with data source?
		Attributes: map[string]schema.Attribute{
			"acl_group_enable": schema.BoolAttribute{
				Description: "Controls whether group access to the queue obeys the access control list defined in the  acl_groups queue attribute.",
				Optional:    true,
			},
			"acl_groups": schema.StringAttribute{
				Description: "List of groups which are allowed or denied access to this queue. The groups in the list are groups on the server host, not submitting hosts.  List is evaluated left-to-right; first match in list is used.",
				Optional:    true,
			},
			"acl_host_enable": schema.BoolAttribute{
				Description: "Controls whether host access to the queue obeys the access control list defined in the acl_hosts queue attribute.",
				Optional:    true,
			},
			"acl_hosts": schema.StringAttribute{
				Description: "List of hosts from which jobs may be submitted to this queue.  List is evaluated left-to-right; first match in list is used.",
				Optional:    true,
			},
			"acl_user_enable": schema.BoolAttribute{
				Description: "Controls whether user access to the queue obeys the access control list defined in the acl_users queue attribute",
				Optional:    true,
			},
			"acl_users": schema.StringAttribute{
				Description: "List of users allowed or denied access to this queue. List is evaluated left-to-right; first match in list is used.",
				Optional:    true,
			},
			"alt_router": schema.StringAttribute{
				DeprecationMessage: "No longer used.",
				Optional:           true,
			},
			"backfill_depth": schema.Int32Attribute{
				Description: "Specifies backfilling behavior for this queue.  Sets the number of jobs that are to be backfilled around in this queue. Overrides backfill_depth server attribute. Recommendation: set this to less than 100.",
				Optional:    true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"checkpoint_min": schema.Int32Attribute{
				Description: "Minimum number of minutes of CPU time or walltime allowed between checkpoints of a job.  If a user specifies a time less than this value, this value is used instead.  The value given in checkpoint_min is used for both CPU minutes and walltime minutes.",
				Optional:    true,
			},
			"default_chunk": schema.StringAttribute{
				Description: "The list of resources which will be inserted into each chunk of a job's select specification if the corresponding resource is not specified by the user. This provides a means for a site to be sure a given resource is properly accounted for even if not specified by the user. ",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Specifies whether this queue accepts new jobs.",
				Optional:    true,
			},
			"from_route_only": schema.BoolAttribute{
				Description: " Specifies whether this queue accepts jobs only from routing queues, or from both execution and routing queues.",
				Optional:    true,
			},
			"kill_delay": schema.Int32Attribute{
				MarkdownDescription: "The time delay (seconds) between sending SIGTERM and SIGKILL when a `qdel` command is issued against a running job. Default value is 10 seconds.",
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"max_array_size": schema.Int32Attribute{
				Description: "The maximum number of subjobs that are allowed in an array job.",
				Optional:    true,
			},
			"max_group_res": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit attributes.  The maximum amount of the specified resource that any single group may consume in a complex.",
				Optional:    true,
			},
			"max_group_res_soft": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The soft limit on the amount of the specified resource that any single group may consume in a complex. If a group is consuming more than this amount of the specified resource, their jobs are eligible to be preempted by jobs from groups who are not over their soft limit.",
				Optional:    true,
			},
			"max_group_run": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs owned by users in a single group that are allowed to be running from this queue at one time. ",
				Optional:    true,
			},
			"max_group_run_soft": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The maximum number of jobs owned by users in a single group that are allowed to be running from this queue at one time. If a group has more than this number of jobs running, their jobs are eligible to be preempted by jobs from groups who are not over their soft limit. ",
				Optional:    true,
			},
			"max_queuable": schema.Int32Attribute{
				Description: " Old limit attribute.  Incompatible with new limit attributes.  The maximum number of jobs allowed to reside in this queue at any given time.",
				Optional:    true,
			},
			"max_queued": schema.StringAttribute{
				Description: "Limit attribute.  The maximum number of jobs allowed to be queued  in  or running from this queue.  Can be specified for  projects, users, groups, or all.  Cannot  be used  with old limit attributes.",
				Optional:    true,
			},
			"max_queued_res": schema.StringAttribute{
				Description: "Limit attribute.  The maximum amount of the specified resource allowed to be allocated to jobs queued in or running from this queue.  Can be specified for  projects, users, groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"max_run": schema.StringAttribute{
				Description: "Limit attribute. The maximum number of jobs allowed to be running from this queue.  Can be specified for projects, users,  groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"max_run_res": schema.StringAttribute{
				Description: "Limit attribute. The maximum amount of the specified resource allowed to be allocated to jobs running from this queue.  Can be specified for  projects, users, groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"max_run_res_soft": schema.StringAttribute{
				Description: "Limit attribute.  Soft limit on the amount of the specified resource allowed to be allocated to jobs running from this queue.  Can be specified for  projects, users, groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"max_run_soft": schema.StringAttribute{
				Description: "Limit attribute.  Soft limit on the number of jobs allowed to be running from this  queue.   Can be specified  for   projects, users, groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"max_running": schema.Int32Attribute{
				Description: "Old limit attribute. Incompatible with new limit  attributes.For an execution queue, this is the largest number of jobs allowed to be running at any given time. For a routing queue, this is the largest number of jobs allowed to be transiting from this queue at any given time.",
				Optional:    true,
			},
			"max_user_res": schema.StringAttribute{
				Description: "Old limit attribute.  Incompatible with new limit attributes.  The maximum amount of the specified resource that any single user may consume. ",
				Optional:    true,
			},
			"max_user_res_soft": schema.StringAttribute{
				Description: "Old limit attribute.  Incompatible with new limit  attributes.  The soft limit on the amount of the specified resource that any single user may consume.  If a user is consuming more than this amount of the specified resource, their jobs are eligible to be preempted by jobs from users who are not over their soft limit. ",
				Optional:    true,
			},
			"max_user_run": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit attributes.  The maximum number of jobs owned by a single user that are allowed to be running from this queue at one time. ",
				Optional:    true,
			},
			"max_user_run_soft": schema.Int32Attribute{
				Description: "Old limit attribute.  Incompatible with new limit attributes.  The soft limit on the number of jobs owned by any single user that are allowed to be running from this queue at one time. If a user has more than this number of jobs running, their jobs are eligible to be preempted by jobs from users who are not over their soft limit. ",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The unique name of the queue on the server",
				Required:    true,
			},
			"node_group_key": schema.StringAttribute{
				Description: "Specifies the resources to use for placement sets (node grouping). Overrides server's node_group_key attribute.  Specified resources must be of type string_array.",
				Optional:    true,
			},
			"partition": schema.StringAttribute{
				Description: "Name of partition to which this queue is assigned.  Cannot be set for routing queue.  An execution queue cannot be changed to a routing queue while this attribute is set.",
				Optional:    true,
			},
			"priority": schema.Int32Attribute{
				Description: "The priority of this queue compared to other queues of the same type in this PBS complex.   Priority can define a queue as an express queue.  See preempt_queue_prio in Chapter 4, \"Scheduler Parameters\", on page 251. Used for execution queues only; the value of Priority has no meaning for routing queues.",
				Optional:    true,
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Description: "Limit attribute.  The maximum number of jobs allowed to be queued in this queue.  Can be specified for  projects, users, groups, or all.  Cannot be used  with old limit attributes.",
				Optional:    true,
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Description: " Limit attribute.  The maximum amount of the specified resource allowed to be allocated to jobs queued in this queue.  Can be specified for  projects, users, groups, or all.  Cannot be used with old limit attributes.",
				Optional:    true,
			},
			"queue_type": schema.StringAttribute{
				Description: "The type of this queue. This attribute must be explicitly set at queue creation to one of Execution/Route",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("Execution", "Route"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resources_assigned": schema.StringAttribute{
				Description: "The total for each kind of resource allocated to running and exiting jobs in this queue.",
				Optional:    true,
			},
			"resources_available": schema.StringAttribute{
				Description: "The list of resources and amounts available to jobs running in this queue. The sum of the resource of each type used by all jobs running from this queue cannot exceed the total amount listed here. ",
				Optional:    true,
			},
			"resources_default": schema.StringAttribute{
				MarkdownDescription: " The list of default resource values which are set as limits for a job residing in this queue and for which the job did not specify a limit.  If not set, the default limit for a job is determined by the first of the following attributes which is set: server's `resources_default`, queue's `resources_max`, server's `resources_max`. If none of these is set, the job gets unlimited resource usage.",
				Optional:            true,
			},
			"resources_max": schema.StringAttribute{
				Description: "The maximum amount of each resource that can be requested by a single job in this queue. This queue value supersedes any server wide maximum limit. ",
				Optional:    true,
			},
			"resources_min": schema.StringAttribute{
				Description: "The minimum amount of each resource that can be requested by a single job in this queue. ",
				Optional:    true,
			},
			"route_destinations": schema.StringAttribute{
				Description: "The list of destinations to which jobs may be routed. Must be set to at least one valid destination.",
				Optional:    true,
			},
			"route_held_jobs": schema.BoolAttribute{
				Description: "Specifies whether jobs in the held state can be routed from this queue.",
				Optional:    true,
			},
			"route_lifetime": schema.Int32Attribute{
				Description: "The maximum time a job is allowed to reside in this routing queue. If a job cannot be routed in this amount of time, the job is aborted. ",
				Optional:    true,
			},
			"route_retry_time": schema.Int32Attribute{
				Description: "Time delay between routing retries. Typically used when the network between servers is down. ",
				Optional:    true,
			},
			"route_waiting_jobs": schema.BoolAttribute{
				MarkdownDescription: "Specifies whether jobs whose `Execution_Time` attribute value is in the future can be routed from this queue.",
				Optional:            true,
			},
			"started": schema.BoolAttribute{
				Description: " If this is an execution queue, specifies whether jobs in this queue can be scheduled for execution, or if this is a routing queue, whether jobs can be routed.",
				Optional:    true,
			},
		},
	}
}

func (r *queueResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *queueResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var queueModel queueModel
	var queue pbsclient.PbsQueue
	diags := req.Plan.Get(ctx, &queueModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	queue, err := r.client.CreateQueue(queueModel.ToPbsQueue())
	if err != nil {
		resp.Diagnostics.AddError("Error creating queue", "Could not create queue, unexpected error: "+err.Error())
		return
	}

	queueModel = createQueueModel(queue)

	diags = resp.State.Set(ctx, queueModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *queueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var queue queueModel

	resp.Diagnostics.Append(req.State.Get(ctx, &queue)...)

	if resp.Diagnostics.HasError() {
		return
	}

	q, err := r.client.GetQueue(queue.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read queues, got error: %s", err))
		return
	}

	// If the queue doesn't exist then remove from state
	if q.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	queue = createQueueModel(q)

	resp.Diagnostics.Append(resp.State.Set(ctx, &queue)...)
}

func (r *queueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data queueModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	_, err := r.client.UpdateQueue(data.ToPbsQueue())
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

func (r *queueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var queue queueModel

	resp.Diagnostics.Append(req.State.Get(ctx, &queue)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteQueue(queue.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete queue, got error: %s", err))
		return
	}
}

func (r *queueResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
