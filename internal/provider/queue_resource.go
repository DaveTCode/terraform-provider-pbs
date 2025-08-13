package provider

import (
	"context"
	"fmt"
	"strings"
	"terraform-provider-pbs/internal/pbsclient"
	validators "terraform-provider-pbs/internal/provider/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueID,
			},
			"acl_group_enable": schema.BoolAttribute{
				MarkdownDescription: DescQueueAclGroupEnable,
				Optional:            true,
			},
			"acl_groups": schema.StringAttribute{
				MarkdownDescription: DescQueueAclGroups,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_groups_normalized": schema.StringAttribute{
				MarkdownDescription: DescQueueAclGroupsNormalized + " This field is computed and reflects the actual value used by PBS.",
				Computed:            true,
			},
			"acl_host_enable": schema.BoolAttribute{
				MarkdownDescription: DescQueueAclHostEnable,
				Optional:            true,
			},
			"acl_hosts": schema.StringAttribute{
				MarkdownDescription: DescQueueAclHosts,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_hosts_normalized": schema.StringAttribute{
				MarkdownDescription: DescQueueAclHostsNormalized + " This field is computed and reflects the actual value used by PBS.",
				Computed:            true,
			},
			"acl_user_enable": schema.BoolAttribute{
				MarkdownDescription: DescQueueAclUserEnable,
				Optional:            true,
			},
			"acl_users": schema.StringAttribute{
				MarkdownDescription: DescQueueAclUsers,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_users_normalized": schema.StringAttribute{
				MarkdownDescription: DescQueueAclUsersNormalized + " This field is computed and reflects the actual value used by PBS.",
				Computed:            true,
			},
			"alt_router": schema.StringAttribute{
				DeprecationMessage:  DescQueueAltRouter,
				MarkdownDescription: DescQueueAltRouter,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"backfill_depth": schema.Int32Attribute{
				MarkdownDescription: DescQueueBackfillDepth,
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"checkpoint_min": schema.Int32Attribute{
				MarkdownDescription: DescQueueCheckpointMin,
				Optional:            true,
			},
			"default_chunk": schema.MapAttribute{
				MarkdownDescription: DescQueueDefaultChunk,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: DescQueueEnabled,
				Optional:            true,
			},
			"from_route_only": schema.BoolAttribute{
				MarkdownDescription: DescQueueFromRouteOnly,
				Optional:            true,
			},
			"kill_delay": schema.Int32Attribute{
				MarkdownDescription: DescQueueKillDelay,
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"max_array_size": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxArraySize,
				Optional:            true,
			},
			"max_group_res": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxGroupRes,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_group_res_soft": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxGroupResSoft,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_group_run": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxGroupRun,
				Optional:            true,
			},
			"max_group_run_soft": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxGroupRunSoft,
				Optional:            true,
			},
			"max_queuable": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxQueuable,
				Optional:            true,
			},
			"max_queued": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxQueued,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_queued_res": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxQueuedRes,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_run": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxRun,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_run_res": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxRunRes,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_run_res_soft": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxRunResSoft,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_run_soft": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxRunSoft,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_running": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxRunning,
				Optional:            true,
			},
			"max_user_res": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxUserRes,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_user_res_soft": schema.MapAttribute{
				MarkdownDescription: DescQueueMaxUserResSoft,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_user_run": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxUserRun,
				Optional:            true,
			},
			"max_user_run_soft": schema.Int32Attribute{
				MarkdownDescription: DescQueueMaxUserRunSoft,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: DescQueueName,
				Required:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"node_group_key": schema.StringAttribute{
				MarkdownDescription: DescQueueNodeGroupKey,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"partition": schema.StringAttribute{
				MarkdownDescription: DescQueuePartition,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"priority": schema.Int32Attribute{
				MarkdownDescription: DescQueuePriority,
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"queued_jobs_threshold": schema.StringAttribute{
				MarkdownDescription: DescQueueQueuedJobsThreshold,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				MarkdownDescription: DescQueueQueuedJobsThresholdRes,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"queue_type": schema.StringAttribute{
				MarkdownDescription: DescQueueQtype,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Execution", "Route"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resources_assigned": schema.MapAttribute{
				MarkdownDescription: DescQueueResourcesAssigned,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"resources_available": schema.MapAttribute{
				MarkdownDescription: DescQueueResourcesAvailable,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"resources_default": schema.MapAttribute{
				MarkdownDescription: DescQueueResourcesDefault,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"resources_max": schema.MapAttribute{
				MarkdownDescription: DescQueueResourcesMax,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"resources_min": schema.MapAttribute{
				MarkdownDescription: DescQueueResourcesMin,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"route_destinations": schema.StringAttribute{
				MarkdownDescription: DescQueueRouteDestinations,
				Optional:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"route_held_jobs": schema.BoolAttribute{
				MarkdownDescription: DescQueueRouteHeldJobs,
				Optional:            true,
			},
			"route_lifetime": schema.Int32Attribute{
				MarkdownDescription: DescQueueRouteLifetime,
				Optional:            true,
			},
			"route_retry_time": schema.Int32Attribute{
				MarkdownDescription: DescQueueRouteRetryTime,
				Optional:            true,
			},
			"route_waiting_jobs": schema.BoolAttribute{
				MarkdownDescription: DescQueueRouteWaitingJobs,
				Optional:            true,
			},
			"started": schema.BoolAttribute{
				MarkdownDescription: DescQueueStarted,
				Optional:            true,
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
	var planModel queueModel
	diags := req.Plan.Get(ctx, &planModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pbsQueueObj, diags := planModel.ToPbsQueue(ctx)
	resp.Diagnostics.Append(diags...)
	queue, err := r.client.CreateQueue(pbsQueueObj)
	if err != nil {
		resp.Diagnostics.AddError("Error creating queue", "Could not create queue, unexpected error: "+err.Error())
		return
	}

	// Create the model from the queue returned by PBS.
	resultModel := createQueueModel(queue)

	// Preserve the user's original format for ACL fields from the plan.
	preserveUserAclFormat(&planModel, &resultModel)

	diags = resp.State.Set(ctx, resultModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *queueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state queueModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	queueName := state.Name.ValueString()
	if queueName == "" && !state.ID.IsNull() {
		queueName = state.ID.ValueString()
	}

	q, err := r.client.GetQueue(queueName)
	if err != nil {
		// Check if queue doesn't exist and remove from state
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "Unknown queue") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read queue, got error: %s", err))
		return
	}

	// If the queue doesn't exist then remove from state
	if q.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	// Update state with current values, preserving plan-only values
	updatedState := createQueueModel(q)
	// Preserve the name from the original state to avoid unnecessary changes,
	// but only if it's not empty (during import, state.Name will be empty)
	if !state.Name.IsNull() && state.Name.ValueString() != "" {
		updatedState.Name = state.Name
	}

	// Preserve the user's ACL format if it's semantically equivalent to what PBS returned
	preserveUserAclFormatFromState(&state, &updatedState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedState)...)
}

func (r *queueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planModel queueModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planModel)...)

	queue, diags := planModel.ToPbsQueue(ctx)
	resp.Diagnostics.Append(diags...)
	updatedQueue, err := r.client.UpdateQueue(queue)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Create the model from the updated queue to ensure all fields including ID are properly set
	updatedModel := createQueueModel(updatedQueue)

	// Preserve the user's original format for ACL fields from the plan
	preserveUserAclFormat(&planModel, &updatedModel)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedModel)...)
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
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
