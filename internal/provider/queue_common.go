package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type queueModel struct {
	AclGroupEnable         types.Bool   `tfsdk:"acl_group_enable"`
	AclGroups              types.String `tfsdk:"acl_groups"`
	AclHostEnable          types.Bool   `tfsdk:"acl_host_enable"`
	AclHosts               types.String `tfsdk:"acl_hosts"`
	AclUserEnable          types.Bool   `tfsdk:"acl_user_enable"`
	AclUsers               types.String `tfsdk:"acl_users"`
	AltRouter              types.String `tfsdk:"alt_router"`
	BackfillDepth          types.Int32  `tfsdk:"backfill_depth"`
	CheckpointMin          types.Int32  `tfsdk:"checkpoint_min"`
	DefaultChunk           types.String `tfsdk:"default_chunk"`
	Enabled                types.Bool   `tfsdk:"enabled"`
	FromRouteOnly          types.Bool   `tfsdk:"from_route_only"`
	KillDelay              types.Int32  `tfsdk:"kill_delay"`
	MaxArraySize           types.Int32  `tfsdk:"max_array_size"`
	MaxGroupRes            types.Int32  `tfsdk:"max_group_res"`
	MaxGroupResSoft        types.Int32  `tfsdk:"max_group_res_soft"`
	MaxGroupRun            types.Int32  `tfsdk:"max_group_run"`
	MaxGroupRunSoft        types.Int32  `tfsdk:"max_group_run_soft"`
	MaxQueuable            types.Int32  `tfsdk:"max_queuable"`
	MaxQueued              types.String `tfsdk:"max_queued"`
	MaxQueuedRes           types.String `tfsdk:"max_queued_res"`
	MaxRun                 types.String `tfsdk:"max_run"`
	MaxRunRes              types.String `tfsdk:"max_run_res"`
	MaxRunResSoft          types.String `tfsdk:"max_run_res_soft"`
	MaxRunSoft             types.String `tfsdk:"max_run_soft"`
	MaxRunning             types.Int32  `tfsdk:"max_running"`
	MaxUserRes             types.String `tfsdk:"max_user_res"`
	MaxUserResSoft         types.String `tfsdk:"max_user_res_soft"`
	MaxUserRun             types.Int32  `tfsdk:"max_user_run"`
	MaxUserRunSoft         types.Int32  `tfsdk:"max_user_run_soft"`
	Name                   types.String `tfsdk:"name"`
	NodeGroupKey           types.String `tfsdk:"node_group_key"`
	Partition              types.String `tfsdk:"partition"`
	Priority               types.Int32  `tfsdk:"priority"`
	QueuedJobsThreshold    types.String `tfsdk:"queued_jobs_threshold"`
	QueuedJobsThresholdRes types.String `tfsdk:"queued_jobs_threshold_res"`
	QueueType              types.String `tfsdk:"queue_type"`
	ResourcesAssigned      types.String `tfsdk:"resources_assigned"`
	ResourcesAvailable     types.String `tfsdk:"resources_available"`
	ResourcesDefault       types.String `tfsdk:"resources_default"`
	ResourcesMax           types.String `tfsdk:"resources_max"`
	ResourcesMin           types.String `tfsdk:"resources_min"`
	RouteDestinations      types.String `tfsdk:"route_destinations"`
	RouteHeldJobs          types.Bool   `tfsdk:"route_held_jobs"`
	RouteLifetime          types.Int32  `tfsdk:"route_lifetime"`
	RouteRetryTime         types.Int32  `tfsdk:"route_retry_time"`
	RouteWaitingJobs       types.Bool   `tfsdk:"route_waiting_jobs"`
	Started                types.Bool   `tfsdk:"started"`
}

func (m queueModel) ToPbsQueue() pbsclient.PbsQueue {
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
		ResourcesAssigned:      m.ResourcesAssigned.ValueStringPointer(),
		ResourcesAvailable:     m.ResourcesAvailable.ValueStringPointer(),
		ResourcesDefault:       m.ResourcesDefault.ValueStringPointer(),
		ResourcesMax:           m.ResourcesMax.ValueStringPointer(),
		ResourcesMin:           m.ResourcesMin.ValueStringPointer(),
		RouteDestinations:      m.RouteDestinations.ValueStringPointer(),
		RouteHeldJobs:          m.RouteHeldJobs.ValueBoolPointer(),
		RouteLifetime:          m.RouteLifetime.ValueInt32Pointer(),
		RouteRetryTime:         m.RouteRetryTime.ValueInt32Pointer(),
		RouteWaitingJobs:       m.RouteWaitingJobs.ValueBoolPointer(),
	}

	return queue
}

func createQueueModel(queue pbsclient.PbsQueue) queueModel {
	model := queueModel{
		Name:      types.StringValue(queue.Name),
		Enabled:   types.BoolValue(queue.Enabled),
		Started:   types.BoolValue(queue.Started),
		QueueType: types.StringValue(queue.QueueType),
	}

	if queue.AclGroupEnable != nil {
		model.AclGroupEnable = types.BoolValue(*queue.AclGroupEnable)
	}
	if queue.AclGroups != nil {
		model.AclGroups = types.StringValue(*queue.AclGroups)
	}
	if queue.AclHostEnable != nil {
		model.AclHostEnable = types.BoolValue(*queue.AclHostEnable)
	}
	if queue.AclHosts != nil {
		model.AclHosts = types.StringValue(*queue.AclHosts)
	}
	if queue.AclUserEnable != nil {
		model.AclUserEnable = types.BoolValue(*queue.AclUserEnable)
	}
	if queue.AclUsers != nil {
		model.AclUsers = types.StringValue(*queue.AclUsers)
	}
	if queue.AltRouter != nil {
		model.AltRouter = types.StringValue(*queue.AltRouter)
	}
	if queue.BackfillDepth != nil {
		model.BackfillDepth = types.Int32Value(int32(*queue.BackfillDepth))
	}
	if queue.CheckpointMin != nil {
		model.CheckpointMin = types.Int32Value(int32(*queue.CheckpointMin))
	}
	if queue.DefaultChunk != nil {
		model.DefaultChunk = types.StringValue(*queue.DefaultChunk)
	}
	if queue.FromRouteOnly != nil {
		model.FromRouteOnly = types.BoolValue(*queue.FromRouteOnly)
	}
	if queue.KillDelay != nil {
		model.KillDelay = types.Int32Value(int32(*queue.KillDelay))
	}
	if queue.MaxArraySize != nil {
		model.MaxArraySize = types.Int32Value(int32(*queue.MaxArraySize))
	}
	if queue.MaxGroupRes != nil {
		model.MaxGroupRes = types.Int32Value(int32(*queue.MaxGroupRes))
	}
	if queue.MaxGroupResSoft != nil {
		model.MaxGroupResSoft = types.Int32Value(int32(*queue.MaxGroupResSoft))
	}
	if queue.MaxGroupRun != nil {
		model.MaxGroupRun = types.Int32Value(int32(*queue.MaxGroupRun))
	}
	if queue.MaxGroupRunSoft != nil {
		model.MaxGroupRunSoft = types.Int32Value(int32(*queue.MaxGroupRunSoft))
	}
	if queue.MaxQueuable != nil {
		model.MaxQueuable = types.Int32Value(int32(*queue.MaxQueuable))
	}
	if queue.MaxQueued != nil {
		model.MaxQueued = types.StringValue(*queue.MaxQueued)
	}
	if queue.MaxQueuedRes != nil {
		model.MaxQueuedRes = types.StringValue(*queue.MaxQueuedRes)
	}
	if queue.MaxRun != nil {
		model.MaxRun = types.StringValue(*queue.MaxRun)
	}
	if queue.MaxRunRes != nil {
		model.MaxRunRes = types.StringValue(*queue.MaxRunRes)
	}
	if queue.MaxRunResSoft != nil {
		model.MaxRunResSoft = types.StringValue(*queue.MaxRunResSoft)
	}
	if queue.MaxRunSoft != nil {
		model.MaxRunSoft = types.StringValue(*queue.MaxRunSoft)
	}
	if queue.MaxRunning != nil {
		model.MaxRunning = types.Int32Value(int32(*queue.MaxRunning))
	}
	if queue.MaxUserRes != nil {
		model.MaxUserRes = types.StringValue(*queue.MaxUserRes)
	}
	if queue.MaxUserResSoft != nil {
		model.MaxUserResSoft = types.StringValue(*queue.MaxUserResSoft)
	}
	if queue.MaxUserRun != nil {
		model.MaxUserRun = types.Int32Value(int32(*queue.MaxUserRun))
	}
	if queue.MaxUserRunSoft != nil {
		model.MaxUserRunSoft = types.Int32Value(int32(*queue.MaxUserRunSoft))
	}
	if queue.NodeGroupKey != nil {
		model.NodeGroupKey = types.StringValue(*queue.NodeGroupKey)
	}
	if queue.Partition != nil {
		model.Partition = types.StringValue(*queue.Partition)
	}
	if queue.Priority != nil {
		model.Priority = types.Int32Value(int32(*queue.Priority))
	}
	if queue.QueuedJobsThreshold != nil {
		model.QueuedJobsThreshold = types.StringValue(*queue.QueuedJobsThreshold)
	}
	if queue.QueuedJobsThresholdRes != nil {
		model.QueuedJobsThresholdRes = types.StringValue(*queue.QueuedJobsThresholdRes)
	}
	if queue.ResourcesAssigned != nil {
		model.ResourcesAssigned = types.StringValue(*queue.ResourcesAssigned)
	}
	if queue.ResourcesAvailable != nil {
		model.ResourcesAvailable = types.StringValue(*queue.ResourcesAvailable)
	}
	if queue.ResourcesDefault != nil {
		model.ResourcesDefault = types.StringValue(*queue.ResourcesDefault)
	}
	if queue.ResourcesMax != nil {
		model.ResourcesMax = types.StringValue(*queue.ResourcesMax)
	}
	if queue.ResourcesMin != nil {
		model.ResourcesMin = types.StringValue(*queue.ResourcesMin)
	}
	if queue.RouteDestinations != nil {
		model.RouteDestinations = types.StringValue(*queue.RouteDestinations)
	}
	if queue.RouteHeldJobs != nil {
		model.RouteHeldJobs = types.BoolValue(*queue.RouteHeldJobs)
	}
	if queue.RouteLifetime != nil {
		model.RouteLifetime = types.Int32Value(int32(*queue.RouteLifetime))
	}
	if queue.RouteRetryTime != nil {
		model.RouteRetryTime = types.Int32Value(int32(*queue.RouteRetryTime))
	}
	if queue.RouteWaitingJobs != nil {
		model.RouteWaitingJobs = types.BoolValue(*queue.RouteWaitingJobs)
	}

	return model
}
