package provider

import (
	"context"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type queueModel struct {
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
	MaxGroupRes            types.Int32             `tfsdk:"max_group_res"`
	MaxGroupResSoft        types.Int32             `tfsdk:"max_group_res_soft"`
	MaxGroupRun            types.Int32             `tfsdk:"max_group_run"`
	MaxGroupRunSoft        types.Int32             `tfsdk:"max_group_run_soft"`
	MaxQueuable            types.Int32             `tfsdk:"max_queuable"`
	MaxQueued              types.String            `tfsdk:"max_queued"`
	MaxQueuedRes           types.String            `tfsdk:"max_queued_res"`
	MaxRun                 types.String            `tfsdk:"max_run"`
	MaxRunRes              types.String            `tfsdk:"max_run_res"`
	MaxRunResSoft          types.String            `tfsdk:"max_run_res_soft"`
	MaxRunSoft             types.String            `tfsdk:"max_run_soft"`
	MaxRunning             types.Int32             `tfsdk:"max_running"`
	MaxUserRes             types.String            `tfsdk:"max_user_res"`
	MaxUserResSoft         types.String            `tfsdk:"max_user_res_soft"`
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
		Name:                   m.Name.ValueString(),
		Enabled:                m.Enabled.ValueBool(),
		Started:                m.Started.ValueBool(),
		QueueType:              m.QueueType.ValueString(),
		AclGroupEnable:         m.AclGroupEnable.ValueBoolPointer(),
		AclGroups:              m.AclGroups.ValueStringPointer(),
		AclHostEnable:          m.AclHostEnable.ValueBoolPointer(),
		AclHosts:               m.AclHosts.ValueStringPointer(),
		AclUserEnable:          m.AclUserEnable.ValueBoolPointer(),
		AclUsers:               m.AclUsers.ValueStringPointer(),
		AltRouter:              m.AltRouter.ValueStringPointer(),
		BackfillDepth:          m.BackfillDepth.ValueInt32Pointer(),
		CheckpointMin:          m.CheckpointMin.ValueInt32Pointer(),
		DefaultChunk:           m.DefaultChunk.ValueStringPointer(),
		FromRouteOnly:          m.FromRouteOnly.ValueBoolPointer(),
		KillDelay:              m.KillDelay.ValueInt32Pointer(),
		MaxArraySize:           m.MaxArraySize.ValueInt32Pointer(),
		MaxGroupRes:            m.MaxGroupRes.ValueInt32Pointer(),
		MaxGroupResSoft:        m.MaxGroupResSoft.ValueInt32Pointer(),
		MaxGroupRun:            m.MaxGroupRun.ValueInt32Pointer(),
		MaxGroupRunSoft:        m.MaxGroupRunSoft.ValueInt32Pointer(),
		MaxQueuable:            m.MaxQueuable.ValueInt32Pointer(),
		MaxQueued:              m.MaxQueued.ValueStringPointer(),
		MaxQueuedRes:           m.MaxQueuedRes.ValueStringPointer(),
		MaxRun:                 m.MaxRun.ValueStringPointer(),
		MaxRunRes:              m.MaxRunRes.ValueStringPointer(),
		MaxRunResSoft:          m.MaxRunResSoft.ValueStringPointer(),
		MaxRunSoft:             m.MaxRunSoft.ValueStringPointer(),
		MaxRunning:             m.MaxRunning.ValueInt32Pointer(),
		MaxUserRes:             m.MaxUserRes.ValueStringPointer(),
		MaxUserResSoft:         m.MaxUserResSoft.ValueStringPointer(),
		MaxUserRun:             m.MaxUserRun.ValueInt32Pointer(),
		MaxUserRunSoft:         m.MaxUserRunSoft.ValueInt32Pointer(),
		NodeGroupKey:           m.NodeGroupKey.ValueStringPointer(),
		Partition:              m.Partition.ValueStringPointer(),
		Priority:               m.Priority.ValueInt32Pointer(),
		QueuedJobsThreshold:    m.QueuedJobsThreshold.ValueStringPointer(),
		QueuedJobsThresholdRes: m.QueuedJobsThresholdRes.ValueStringPointer(),
		RouteDestinations:      m.RouteDestinations.ValueStringPointer(),
		RouteHeldJobs:          m.RouteHeldJobs.ValueBoolPointer(),
		RouteLifetime:          m.RouteLifetime.ValueInt32Pointer(),
		RouteRetryTime:         m.RouteRetryTime.ValueInt32Pointer(),
		RouteWaitingJobs:       m.RouteWaitingJobs.ValueBoolPointer(),
	}

	var diags diag.Diagnostics
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
	model.MaxGroupRes = types.Int32PointerValue(queue.MaxGroupRes)
	model.MaxGroupResSoft = types.Int32PointerValue(queue.MaxGroupResSoft)
	model.MaxGroupRun = types.Int32PointerValue(queue.MaxGroupRun)
	model.MaxGroupRunSoft = types.Int32PointerValue(queue.MaxGroupRunSoft)
	model.MaxQueuable = types.Int32PointerValue(queue.MaxQueuable)
	model.MaxQueued = types.StringPointerValue(queue.MaxQueued)
	model.MaxQueuedRes = types.StringPointerValue(queue.MaxQueuedRes)
	model.MaxRun = types.StringPointerValue(queue.MaxRun)
	model.MaxRunRes = types.StringPointerValue(queue.MaxRunRes)
	model.MaxRunResSoft = types.StringPointerValue(queue.MaxRunResSoft)
	model.MaxRunSoft = types.StringPointerValue(queue.MaxRunSoft)
	model.MaxRunning = types.Int32PointerValue(queue.MaxRunning)
	model.MaxUserRes = types.StringPointerValue(queue.MaxUserRes)
	model.MaxUserResSoft = types.StringPointerValue(queue.MaxUserResSoft)
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
