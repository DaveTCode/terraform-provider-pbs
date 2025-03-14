package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pbsNodeModel struct {
	Comment            types.String            `tfsdk:"comment"`
	CurrentAoe         types.String            `tfsdk:"current_aoe"`
	CurrentEoe         types.String            `tfsdk:"current_eoe"`
	InMultiNodeHost    types.Int32             `tfsdk:"in_multi_node_host"`
	Mom                types.String            `tfsdk:"mom"`
	Name               types.String            `tfsdk:"name"`
	NoMultinodeJobs    types.Bool              `tfsdk:"no_multinode_jobs"`
	Partition          types.String            `tfsdk:"partition"`
	PNames             types.String            `tfsdk:"p_names"`
	Port               types.Int32             `tfsdk:"port"`
	PowerOffEligible   types.Bool              `tfsdk:"power_off_eligible"`
	PowerProvisioning  types.Bool              `tfsdk:"power_provisioning"`
	Priority           types.Int32             `tfsdk:"priority"`
	ProvisionEnable    types.Bool              `tfsdk:"provision_enable"`
	Queue              types.String            `tfsdk:"queue"`
	ResourcesAvailable map[string]types.String `tfsdk:"resources_available"`
	ResvEnable         types.Bool              `tfsdk:"resv_enable"`
}

func (m pbsNodeModel) ToPbsNode() pbsclient.PbsNode {
	node := pbsclient.PbsNode{
		Comment:           m.Comment.ValueStringPointer(),
		CurrentAoe:        m.CurrentAoe.ValueStringPointer(),
		CurrentEoe:        m.CurrentEoe.ValueStringPointer(),
		InMultiNodeHost:   m.InMultiNodeHost.ValueInt32Pointer(),
		Mom:               m.Mom.ValueStringPointer(),
		Name:              m.Name.ValueString(),
		NoMultinodeJobs:   m.NoMultinodeJobs.ValueBoolPointer(),
		Partition:         m.Partition.ValueStringPointer(),
		PNames:            m.PNames.ValueStringPointer(),
		Port:              m.Port.ValueInt32Pointer(),
		PowerOffEligible:  m.PowerOffEligible.ValueBoolPointer(),
		PowerProvisioning: m.PowerProvisioning.ValueBoolPointer(),
		Priority:          m.Priority.ValueInt32Pointer(),
		ProvisionEnable:   m.ProvisionEnable.ValueBoolPointer(),
		Queue:             m.Queue.ValueStringPointer(),
		ResvEnable:        m.ResvEnable.ValueBoolPointer(),
	}

	node.ResourcesAvailable = make(map[string]string)
	for k, v := range m.ResourcesAvailable {
		// This is a bit hacky but the host and vnode attributes arent something you can
		// set on the resources_available so we don't want terraform getting confused about
		// what it's managing
		if k != "host" && k != "vnode" {
			node.ResourcesAvailable[k] = v.ValueString()
		}
	}

	return node
}

func createPbsNodeModel(h pbsclient.PbsNode) pbsNodeModel {
	model := pbsNodeModel{
		Name: types.StringValue(h.Name),
	}

	if h.Comment != nil {
		model.Comment = types.StringValue(*h.Comment)
	}
	if h.CurrentAoe != nil {
		model.CurrentAoe = types.StringValue(*h.CurrentAoe)
	}
	if h.CurrentEoe != nil {
		model.CurrentEoe = types.StringValue(*h.CurrentEoe)
	}
	if h.InMultiNodeHost != nil {
		model.InMultiNodeHost = types.Int32Value(*h.InMultiNodeHost)
	}
	if h.Mom != nil {
		model.Mom = types.StringValue(*h.Mom)
	}
	if h.NoMultinodeJobs != nil {
		model.NoMultinodeJobs = types.BoolValue(*h.NoMultinodeJobs)
	}
	if h.Partition != nil {
		model.Partition = types.StringValue(*h.Partition)
	}
	if h.PNames != nil {
		model.PNames = types.StringValue(*h.PNames)
	}
	if h.Port != nil {
		model.Port = types.Int32Value(*h.Port)
	}
	if h.PowerOffEligible != nil {
		model.PowerOffEligible = types.BoolValue(*h.PowerOffEligible)
	}
	if h.PowerProvisioning != nil {
		model.PowerProvisioning = types.BoolValue(*h.PowerProvisioning)
	}
	if h.Priority != nil {
		model.Priority = types.Int32Value(*h.Priority)
	}
	if h.ProvisionEnable != nil {
		model.ProvisionEnable = types.BoolValue(*h.ProvisionEnable)
	}
	if h.Queue != nil {
		model.Queue = types.StringValue(*h.Queue)
	}
	if h.ResvEnable != nil {
		model.ResvEnable = types.BoolValue(*h.ResvEnable)
	}

	model.ResourcesAvailable = make(map[string]types.String)
	for k, v := range h.ResourcesAvailable {
		// This is a bit hacky but the host and vnode attributes arent something you can
		// set on the resources_available so we don't want terraform getting confused about
		// what it's managing
		if k != "host" && k != "vnode" {
			model.ResourcesAvailable[k] = types.StringValue(v)
		}
	}

	return model
}
