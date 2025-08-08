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
	AclHostEnable          types.Bool              `tfsdk:"acl_host_enable"`
	AclHosts               types.String            `tfsdk:"acl_hosts"`
	AclUserEnable          types.Bool              `tfsdk:"acl_user_enable"`
	AclUsers               types.String            `tfsdk:"acl_users"`
	AltRouter              types.String            `tfsdk:"alt_router"`
	BackfillDepth          types.Int32             `tfsdk:"backfill_depth"`
	CheckpointMin          types.Int32             `tfsdk:"checkpoint_min"`
	DefaultChunk           types.String            `tfsdk:"default_chunk"`
	Enabled                types.Bool              `tfsdk:"enabled"`
	FromRouteOnly          types.Bool              `tfsdk:"from_route_only"`
	KillDelay              types.Int32             `tfsdk:"kill_delay"`
	MaxArraySize           types.Int32             `tfsdk:"max_array_size"`
	MaxGroupRes            map[string]types.String `tfsdk:"max_group_res"`
	MaxGroupResSoft        map[string]types.String `tfsdk:"max_group_res_soft"`
	MaxGroupRun            types.Int32             `tfsdk:"max_group_run"`
	MaxGroupRunSoft        types.Int32             `tfsdk:"max_group_run_soft"`
	MaxQueuable            types.Int32             `tfsdk:"max_queuable"`
	MaxQueued              map[string]types.String `tfsdk:"max_queued"`
	MaxQueuedRes           map[string]types.String `tfsdk:"max_queued_res"`
	MaxRun                 map[string]types.String `tfsdk:"max_run"`
	MaxRunRes              map[string]types.String `tfsdk:"max_run_res"`
	MaxRunResSoft          map[string]types.String `tfsdk:"max_run_res_soft"`
	MaxRunSoft             map[string]types.String `tfsdk:"max_run_soft"`
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
	ResourcesAssigned      map[string]types.String `tfsdk:"resources_assigned"`
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
	if !m.AclGroupEnable.IsNull() {
		queue.AclGroupEnable = m.AclGroupEnable.ValueBoolPointer()
	}
	if !m.AclGroups.IsNull() {
		queue.AclGroups = m.AclGroups.ValueStringPointer()
	}
	if !m.AclHostEnable.IsNull() {
		queue.AclHostEnable = m.AclHostEnable.ValueBoolPointer()
	}
	if !m.AclHosts.IsNull() {
		queue.AclHosts = m.AclHosts.ValueStringPointer()
	}
	if !m.AclUserEnable.IsNull() {
		queue.AclUserEnable = m.AclUserEnable.ValueBoolPointer()
	}
	if !m.AclUsers.IsNull() {
		queue.AclUsers = m.AclUsers.ValueStringPointer()
	}
	if !m.AltRouter.IsNull() {
		queue.AltRouter = m.AltRouter.ValueStringPointer()
	}
	if !m.BackfillDepth.IsNull() {
		queue.BackfillDepth = m.BackfillDepth.ValueInt32Pointer()
	}
	if !m.CheckpointMin.IsNull() {
		queue.CheckpointMin = m.CheckpointMin.ValueInt32Pointer()
	}
	if !m.DefaultChunk.IsNull() {
		queue.DefaultChunk = m.DefaultChunk.ValueStringPointer()
	}
	if !m.FromRouteOnly.IsNull() {
		queue.FromRouteOnly = m.FromRouteOnly.ValueBoolPointer()
	}
	if !m.KillDelay.IsNull() {
		queue.KillDelay = m.KillDelay.ValueInt32Pointer()
	}
	if !m.MaxArraySize.IsNull() {
		queue.MaxArraySize = m.MaxArraySize.ValueInt32Pointer()
	}
	if !m.MaxGroupRun.IsNull() {
		queue.MaxGroupRun = m.MaxGroupRun.ValueInt32Pointer()
	}
	if !m.MaxGroupRunSoft.IsNull() {
		queue.MaxGroupRunSoft = m.MaxGroupRunSoft.ValueInt32Pointer()
	}
	if !m.MaxQueuable.IsNull() {
		queue.MaxQueuable = m.MaxQueuable.ValueInt32Pointer()
	}
	if !m.MaxRunning.IsNull() {
		queue.MaxRunning = m.MaxRunning.ValueInt32Pointer()
	}
	if !m.MaxUserRun.IsNull() {
		queue.MaxUserRun = m.MaxUserRun.ValueInt32Pointer()
	}
	if !m.MaxUserRunSoft.IsNull() {
		queue.MaxUserRunSoft = m.MaxUserRunSoft.ValueInt32Pointer()
	}
	if !m.NodeGroupKey.IsNull() {
		queue.NodeGroupKey = m.NodeGroupKey.ValueStringPointer()
	}
	if !m.Partition.IsNull() {
		queue.Partition = m.Partition.ValueStringPointer()
	}
	if !m.Priority.IsNull() {
		queue.Priority = m.Priority.ValueInt32Pointer()
	}
	if !m.QueuedJobsThreshold.IsNull() {
		queue.QueuedJobsThreshold = m.QueuedJobsThreshold.ValueStringPointer()
	}
	if !m.QueuedJobsThresholdRes.IsNull() {
		queue.QueuedJobsThresholdRes = m.QueuedJobsThresholdRes.ValueStringPointer()
	}
	if !m.RouteDestinations.IsNull() {
		queue.RouteDestinations = m.RouteDestinations.ValueStringPointer()
	}
	if !m.RouteHeldJobs.IsNull() {
		queue.RouteHeldJobs = m.RouteHeldJobs.ValueBoolPointer()
	}
	if !m.RouteLifetime.IsNull() {
		queue.RouteLifetime = m.RouteLifetime.ValueInt32Pointer()
	}
	if !m.RouteRetryTime.IsNull() {
		queue.RouteRetryTime = m.RouteRetryTime.ValueInt32Pointer()
	}
	if !m.RouteWaitingJobs.IsNull() {
		queue.RouteWaitingJobs = m.RouteWaitingJobs.ValueBoolPointer()
	}

	var diags diag.Diagnostics

	// Set non-pointer computed fields
	queue.MaxGroupRes = make(map[string]string)
	for k, v := range m.MaxGroupRes {
		queue.MaxGroupRes[k] = v.ValueString()
	}
	queue.MaxGroupResSoft = make(map[string]string)
	for k, v := range m.MaxGroupResSoft {
		queue.MaxGroupResSoft[k] = v.ValueString()
	}
	queue.MaxQueued = make(map[string]string)
	for k, v := range m.MaxQueued {
		queue.MaxQueued[k] = v.ValueString()
	}
	queue.MaxQueuedRes = make(map[string]string)
	for k, v := range m.MaxQueuedRes {
		queue.MaxQueuedRes[k] = v.ValueString()
	}
	queue.MaxRun = make(map[string]string)
	for k, v := range m.MaxRun {
		queue.MaxRun[k] = v.ValueString()
	}
	queue.MaxRunRes = make(map[string]string)
	for k, v := range m.MaxRunRes {
		queue.MaxRunRes[k] = v.ValueString()
	}
	queue.MaxRunResSoft = make(map[string]string)
	for k, v := range m.MaxRunResSoft {
		queue.MaxRunResSoft[k] = v.ValueString()
	}
	queue.MaxRunSoft = make(map[string]string)
	for k, v := range m.MaxRunSoft {
		queue.MaxRunSoft[k] = v.ValueString()
	}
	queue.MaxUserRes = make(map[string]string)
	for k, v := range m.MaxUserRes {
		queue.MaxUserRes[k] = v.ValueString()
	}
	queue.MaxUserResSoft = make(map[string]string)
	for k, v := range m.MaxUserResSoft {
		queue.MaxUserResSoft[k] = v.ValueString()
	}
	queue.ResourcesAssigned = make(map[string]string)
	for k, v := range m.ResourcesAssigned {
		queue.ResourcesAssigned[k] = v.ValueString()
	}
	queue.ResourcesAvailable = make(map[string]string)
	for k, v := range m.ResourcesAvailable {
		queue.ResourcesAvailable[k] = v.ValueString()
	}
	queue.ResourcesDefault = make(map[string]string)
	for k, v := range m.ResourcesDefault {
		queue.ResourcesDefault[k] = v.ValueString()
	}
	queue.ResourcesMax = make(map[string]string)
	for k, v := range m.ResourcesMax {
		queue.ResourcesMax[k] = v.ValueString()
	}
	queue.ResourcesMin = make(map[string]string)
	for k, v := range m.ResourcesMin {
		queue.ResourcesMin[k] = v.ValueString()
	}

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
	model.AclGroups = types.StringPointerValue(queue.AclGroups)
	model.AclHostEnable = types.BoolPointerValue(queue.AclHostEnable)
	model.AclHosts = types.StringPointerValue(queue.AclHosts)
	model.AclUserEnable = types.BoolPointerValue(queue.AclUserEnable)
	model.AclUsers = types.StringPointerValue(queue.AclUsers)
	model.AltRouter = types.StringPointerValue(queue.AltRouter)
	model.BackfillDepth = types.Int32PointerValue(queue.BackfillDepth)
	model.CheckpointMin = types.Int32PointerValue(queue.CheckpointMin)
	model.DefaultChunk = types.StringPointerValue(queue.DefaultChunk)
	model.FromRouteOnly = types.BoolPointerValue(queue.FromRouteOnly)
	model.KillDelay = types.Int32PointerValue(queue.KillDelay)
	model.MaxArraySize = types.Int32PointerValue(queue.MaxArraySize)
	model.MaxGroupRun = types.Int32PointerValue(queue.MaxGroupRun)
	model.MaxGroupRunSoft = types.Int32PointerValue(queue.MaxGroupRunSoft)
	model.MaxQueuable = types.Int32PointerValue(queue.MaxQueuable)
	model.MaxRunning = types.Int32PointerValue(queue.MaxRunning)
	model.MaxUserRun = types.Int32PointerValue(queue.MaxUserRun)
	model.MaxUserRunSoft = types.Int32PointerValue(queue.MaxUserRunSoft)
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

	if queue.MaxGroupRes != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxGroupRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxGroupRes = elements
	}
	if queue.MaxGroupResSoft != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxGroupResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxGroupResSoft = elements
	}
	if queue.MaxQueued != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxQueued {
			elements[k] = types.StringValue(v)
		}
		model.MaxQueued = elements
	}
	if queue.MaxQueuedRes != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxQueuedRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxQueuedRes = elements
	}
	if queue.MaxRun != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxRun {
			elements[k] = types.StringValue(v)
		}
		model.MaxRun = elements
	}
	if queue.MaxRunRes != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxRunRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunRes = elements
	}
	if queue.MaxRunResSoft != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxRunResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunResSoft = elements
	}
	if queue.MaxRunSoft != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxRunSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunSoft = elements
	}
	if queue.MaxUserRes != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxUserRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxUserRes = elements
	}
	if queue.MaxUserResSoft != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.MaxUserResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxUserResSoft = elements
	}

	if queue.ResourcesAssigned != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.ResourcesAssigned {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesAssigned = elements
	}
	if queue.ResourcesAvailable != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.ResourcesAvailable {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesAvailable = elements
	}
	if queue.ResourcesDefault != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.ResourcesDefault {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesDefault = elements
	}
	if queue.ResourcesMax != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.ResourcesMax {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesMax = elements
	}
	if queue.ResourcesMin != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range queue.ResourcesMin {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesMin = elements
	}

	return model
}
