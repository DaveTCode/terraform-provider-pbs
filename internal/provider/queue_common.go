package provider

import (
	"context"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type queueModel struct {
	ID                     types.String            `tfsdk:"id"`
	AclGroupEnable         types.Bool              `tfsdk:"acl_group_enable"`
	AclGroups              types.String            `tfsdk:"acl_groups"`
	AclGroupsNormalized    types.String            `tfsdk:"acl_groups_normalized"`
	AclHostEnable          types.Bool              `tfsdk:"acl_host_enable"`
	AclHosts               types.String            `tfsdk:"acl_hosts"`
	AclHostsNormalized     types.String            `tfsdk:"acl_hosts_normalized"`
	AclUserEnable          types.Bool              `tfsdk:"acl_user_enable"`
	AclUsers               types.String            `tfsdk:"acl_users"`
	AclUsersNormalized     types.String            `tfsdk:"acl_users_normalized"`
	AltRouter              types.String            `tfsdk:"alt_router"`
	BackfillDepth          types.Int32             `tfsdk:"backfill_depth"`
	CheckpointMin          types.Int32             `tfsdk:"checkpoint_min"`
	DefaultChunk           map[string]types.String `tfsdk:"default_chunk"`
	Enabled                types.Bool              `tfsdk:"enabled"`
	FromRouteOnly          types.Bool              `tfsdk:"from_route_only"`
	KillDelay              types.Int32             `tfsdk:"kill_delay"`
	MaxArraySize           types.Int32             `tfsdk:"max_array_size"`
	MaxGroupRes            map[string]types.String `tfsdk:"max_group_res"`
	MaxGroupResSoft        map[string]types.String `tfsdk:"max_group_res_soft"`
	MaxGroupRun            types.Int32             `tfsdk:"max_group_run"`
	MaxGroupRunSoft        types.Int32             `tfsdk:"max_group_run_soft"`
	MaxQueuable            types.Int32             `tfsdk:"max_queuable"`
	MaxQueued              types.String            `tfsdk:"max_queued"`
	MaxQueuedRes           map[string]types.String `tfsdk:"max_queued_res"`
	MaxRun                 types.String            `tfsdk:"max_run"`
	MaxRunRes              map[string]types.String `tfsdk:"max_run_res"`
	MaxRunResSoft          map[string]types.String `tfsdk:"max_run_res_soft"`
	MaxRunSoft             types.String            `tfsdk:"max_run_soft"`
	MaxRunning             types.Int32             `tfsdk:"max_running"`
	MaxUserRes             map[string]types.String `tfsdk:"max_user_res"`
	MaxUserResSoft         map[string]types.String `tfsdk:"max_user_res_soft"`
	MaxUserRun             types.Int32             `tfsdk:"max_user_run"`
	MaxUserRunSoft         types.Int32             `tfsdk:"max_user_run_soft"`
	Name                   types.String            `tfsdk:"name"`
	NodeGroupKey           types.String            `tfsdk:"node_group_key"`
	Partition              types.String            `tfsdk:"partition"`
	Priority               types.Int32             `tfsdk:"priority"`
	QueuedJobsThreshold    types.String            `tfsdk:"queued_jobs_threshold"`
	QueuedJobsThresholdRes types.String            `tfsdk:"queued_jobs_threshold_res"`
	QueueType              types.String            `tfsdk:"queue_type"`
	ResourcesAvailable     map[string]types.String `tfsdk:"resources_available"`
	ResourcesDefault       map[string]types.String `tfsdk:"resources_default"`
	ResourcesMax           map[string]types.String `tfsdk:"resources_max"`
	ResourcesMin           map[string]types.String `tfsdk:"resources_min"`
	RouteDestinations      types.String            `tfsdk:"route_destinations"`
	RouteHeldJobs          types.Bool              `tfsdk:"route_held_jobs"`
	RouteLifetime          types.Int32             `tfsdk:"route_lifetime"`
	RouteRetryTime         types.Int32             `tfsdk:"route_retry_time"`
	RouteWaitingJobs       types.Bool              `tfsdk:"route_waiting_jobs"`
	Started                types.Bool              `tfsdk:"started"`
}

func (m queueModel) ToPbsQueue(ctx context.Context) (pbsclient.PbsQueue, diag.Diagnostics) {
	queue := pbsclient.PbsQueue{
		Name:      m.Name.ValueString(),
		Enabled:   m.Enabled.ValueBool(),
		Started:   m.Started.ValueBool(),
		QueueType: m.QueueType.ValueString(),
	}

	// Only set pointer fields if the value is not null
	SetBoolPointerIfNotNull(m.AclGroupEnable, &queue.AclGroupEnable)
	SetStringPointerIfNotNull(m.AclGroups, &queue.AclGroups)
	SetBoolPointerIfNotNull(m.AclHostEnable, &queue.AclHostEnable)
	SetStringPointerIfNotNull(m.AclHosts, &queue.AclHosts)
	SetBoolPointerIfNotNull(m.AclUserEnable, &queue.AclUserEnable)
	SetStringPointerIfNotNull(m.AclUsers, &queue.AclUsers)
	SetStringPointerIfNotNull(m.AltRouter, &queue.AltRouter)
	SetInt32PointerIfNotNull(m.BackfillDepth, &queue.BackfillDepth)
	SetInt32PointerIfNotNull(m.CheckpointMin, &queue.CheckpointMin)
	SetBoolPointerIfNotNull(m.FromRouteOnly, &queue.FromRouteOnly)
	SetInt32PointerIfNotNull(m.KillDelay, &queue.KillDelay)
	SetInt32PointerIfNotNull(m.MaxArraySize, &queue.MaxArraySize)
	SetInt32PointerIfNotNull(m.MaxGroupRun, &queue.MaxGroupRun)
	SetInt32PointerIfNotNull(m.MaxGroupRunSoft, &queue.MaxGroupRunSoft)
	SetStringPointerIfNotNull(m.MaxQueued, &queue.MaxQueued)
	SetInt32PointerIfNotNull(m.MaxQueuable, &queue.MaxQueuable)
	SetStringPointerIfNotNull(m.MaxRun, &queue.MaxRun)
	SetStringPointerIfNotNull(m.MaxRunSoft, &queue.MaxRunSoft)
	SetInt32PointerIfNotNull(m.MaxRunning, &queue.MaxRunning)
	SetInt32PointerIfNotNull(m.MaxUserRun, &queue.MaxUserRun)
	SetInt32PointerIfNotNull(m.MaxUserRunSoft, &queue.MaxUserRunSoft)
	SetStringPointerIfNotNull(m.NodeGroupKey, &queue.NodeGroupKey)
	SetStringPointerIfNotNull(m.Partition, &queue.Partition)
	SetInt32PointerIfNotNull(m.Priority, &queue.Priority)
	SetStringPointerIfNotNull(m.QueuedJobsThreshold, &queue.QueuedJobsThreshold)
	SetStringPointerIfNotNull(m.QueuedJobsThresholdRes, &queue.QueuedJobsThresholdRes)
	SetStringPointerIfNotNull(m.RouteDestinations, &queue.RouteDestinations)
	SetBoolPointerIfNotNull(m.RouteHeldJobs, &queue.RouteHeldJobs)
	SetInt32PointerIfNotNull(m.RouteLifetime, &queue.RouteLifetime)
	SetInt32PointerIfNotNull(m.RouteRetryTime, &queue.RouteRetryTime)
	SetBoolPointerIfNotNull(m.RouteWaitingJobs, &queue.RouteWaitingJobs)

	var diags diag.Diagnostics

	// Set non-pointer computed fields using utility functions
	queue.DefaultChunk = ConvertTypesStringMap(m.DefaultChunk)
	queue.MaxGroupRes = ConvertTypesStringMap(m.MaxGroupRes)
	queue.MaxGroupResSoft = ConvertTypesStringMap(m.MaxGroupResSoft)
	queue.MaxQueuedRes = ConvertTypesStringMap(m.MaxQueuedRes)
	queue.MaxRunRes = ConvertTypesStringMap(m.MaxRunRes)
	queue.MaxRunResSoft = ConvertTypesStringMap(m.MaxRunResSoft)
	queue.MaxUserRes = ConvertTypesStringMap(m.MaxUserRes)
	queue.MaxUserResSoft = ConvertTypesStringMap(m.MaxUserResSoft)
	queue.ResourcesAvailable = ConvertTypesStringMap(m.ResourcesAvailable)
	queue.ResourcesDefault = ConvertTypesStringMap(m.ResourcesDefault)
	queue.ResourcesMax = ConvertTypesStringMap(m.ResourcesMax)
	queue.ResourcesMin = ConvertTypesStringMap(m.ResourcesMin)

	return queue, diags
}

func createQueueModel(queue pbsclient.PbsQueue) queueModel {
	model := queueModel{
		ID:        types.StringValue(queue.Name), // Use name as ID
		Name:      types.StringValue(queue.Name),
		Enabled:   types.BoolValue(queue.Enabled),
		Started:   types.BoolValue(queue.Started),
		QueueType: types.StringValue(queue.QueueType),
	}

	model.AclGroupEnable = types.BoolPointerValue(queue.AclGroupEnable)
	// Store the normalized version as received from PBS.
	addNormalizedAclField(queue.AclGroups, &model.AclGroupsNormalized)
	// For AclGroups, we'll preserve the user's format when available (handled in Create/Update).
	// For Read operations, we'll use the PBS value.
	model.AclGroups = types.StringPointerValue(queue.AclGroups)
	model.AclHostEnable = types.BoolPointerValue(queue.AclHostEnable)
	// Store the normalized version as received from PBS.
	addNormalizedAclField(queue.AclHosts, &model.AclHostsNormalized)
	// For AclHosts, we'll preserve the user's format when available (handled in Create/Update).
	model.AclHosts = types.StringPointerValue(queue.AclHosts)
	model.AclUserEnable = types.BoolPointerValue(queue.AclUserEnable)
	// Store the normalized version as received from PBS.
	addNormalizedAclField(queue.AclUsers, &model.AclUsersNormalized)
	// For AclUsers, we'll preserve the user's format when available (handled in Create/Update).
	model.AclUsers = types.StringPointerValue(queue.AclUsers)
	model.AltRouter = types.StringPointerValue(queue.AltRouter)
	model.BackfillDepth = types.Int32PointerValue(queue.BackfillDepth)
	model.CheckpointMin = types.Int32PointerValue(queue.CheckpointMin)
	model.FromRouteOnly = types.BoolPointerValue(queue.FromRouteOnly)
	model.KillDelay = types.Int32PointerValue(queue.KillDelay)
	model.MaxArraySize = types.Int32PointerValue(queue.MaxArraySize)
	model.MaxGroupRun = types.Int32PointerValue(queue.MaxGroupRun)
	model.MaxGroupRunSoft = types.Int32PointerValue(queue.MaxGroupRunSoft)
	model.MaxQueued = types.StringPointerValue(queue.MaxQueued)
	model.MaxQueuable = types.Int32PointerValue(queue.MaxQueuable)
	model.MaxRunning = types.Int32PointerValue(queue.MaxRunning)
	model.MaxUserRun = types.Int32PointerValue(queue.MaxUserRun)
	model.MaxUserRunSoft = types.Int32PointerValue(queue.MaxUserRunSoft)
	model.MaxRun = types.StringPointerValue(queue.MaxRun)
	model.MaxRunSoft = types.StringPointerValue(queue.MaxRunSoft)
	model.NodeGroupKey = types.StringPointerValue(queue.NodeGroupKey)
	model.Partition = types.StringPointerValue(queue.Partition)
	model.Priority = types.Int32PointerValue(queue.Priority)
	model.QueuedJobsThreshold = types.StringPointerValue(queue.QueuedJobsThreshold)
	model.QueuedJobsThresholdRes = types.StringPointerValue(queue.QueuedJobsThresholdRes)
	model.RouteDestinations = types.StringPointerValue(queue.RouteDestinations)
	model.RouteHeldJobs = types.BoolPointerValue(queue.RouteHeldJobs)
	model.RouteLifetime = types.Int32PointerValue(queue.RouteLifetime)
	model.RouteRetryTime = types.Int32PointerValue(queue.RouteRetryTime)
	model.RouteWaitingJobs = types.BoolPointerValue(queue.RouteWaitingJobs)

	if queue.DefaultChunk != nil {
		model.DefaultChunk = convertStringMapToTypesStringMap(queue.DefaultChunk)
	}
	if queue.MaxGroupRes != nil {
		model.MaxGroupRes = convertStringMapToTypesStringMap(queue.MaxGroupRes)
	}
	if queue.MaxGroupResSoft != nil {
		model.MaxGroupResSoft = convertStringMapToTypesStringMap(queue.MaxGroupResSoft)
	}
	if queue.MaxQueuedRes != nil {
		model.MaxQueuedRes = convertStringMapToTypesStringMap(queue.MaxQueuedRes)
	}
	if queue.MaxRunRes != nil {
		model.MaxRunRes = convertStringMapToTypesStringMap(queue.MaxRunRes)
	}
	if queue.MaxRunResSoft != nil {
		model.MaxRunResSoft = convertStringMapToTypesStringMap(queue.MaxRunResSoft)
	}
	if queue.MaxUserRes != nil {
		model.MaxUserRes = convertStringMapToTypesStringMap(queue.MaxUserRes)
	}
	if queue.MaxUserResSoft != nil {
		model.MaxUserResSoft = convertStringMapToTypesStringMap(queue.MaxUserResSoft)
	}
	if queue.ResourcesAvailable != nil {
		model.ResourcesAvailable = convertStringMapToTypesStringMap(queue.ResourcesAvailable)
	}
	if queue.ResourcesDefault != nil {
		model.ResourcesDefault = convertStringMapToTypesStringMap(queue.ResourcesDefault)
	}
	if queue.ResourcesMax != nil {
		model.ResourcesMax = convertStringMapToTypesStringMap(queue.ResourcesMax)
	}
	if queue.ResourcesMin != nil {
		model.ResourcesMin = convertStringMapToTypesStringMap(queue.ResourcesMin)
	}

	return model
}

// preserveUserAclFormat preserves user-provided ACL field formats from plan in the result model.
func preserveUserAclFormat(planModel, resultModel *queueModel) {
	planFields := []AclFieldPair{
		{UserField: planModel.AclGroups, NormalizedField: planModel.AclGroupsNormalized},
		{UserField: planModel.AclHosts, NormalizedField: planModel.AclHostsNormalized},
		{UserField: planModel.AclUsers, NormalizedField: planModel.AclUsersNormalized},
	}

	resultFields := []AclFieldPair{
		{UserField: resultModel.AclGroups, NormalizedField: resultModel.AclGroupsNormalized},
		{UserField: resultModel.AclHosts, NormalizedField: resultModel.AclHostsNormalized},
		{UserField: resultModel.AclUsers, NormalizedField: resultModel.AclUsersNormalized},
	}

	preserveUserAclFormats(planFields, resultFields)

	// Update the result model with preserved values
	resultModel.AclGroups = resultFields[0].UserField
	resultModel.AclHosts = resultFields[1].UserField
	resultModel.AclUsers = resultFields[2].UserField
}

// preserveUserAclFormatFromState preserves user-provided ACL field formats from state when semantically equivalent.
func preserveUserAclFormatFromState(state, updatedState *queueModel) {
	stateFields := []AclFieldPair{
		{UserField: state.AclGroups, NormalizedField: state.AclGroupsNormalized},
		{UserField: state.AclHosts, NormalizedField: state.AclHostsNormalized},
		{UserField: state.AclUsers, NormalizedField: state.AclUsersNormalized},
	}

	updatedFields := []AclFieldPair{
		{UserField: updatedState.AclGroups, NormalizedField: updatedState.AclGroupsNormalized},
		{UserField: updatedState.AclHosts, NormalizedField: updatedState.AclHostsNormalized},
		{UserField: updatedState.AclUsers, NormalizedField: updatedState.AclUsersNormalized},
	}

	preserveUserAclFormatsFromState(stateFields, updatedFields)

	// Update the updated state with preserved values
	updatedState.AclGroups = updatedFields[0].UserField
	updatedState.AclHosts = updatedFields[1].UserField
	updatedState.AclUsers = updatedFields[2].UserField
}
