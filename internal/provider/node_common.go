package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pbsNodeModel struct {
	ID                 types.String            `tfsdk:"id"`
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
	PowerOffEligible   types.Bool              `tfsdk:"poweroff_eligible"`
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
		ID:   types.StringValue(h.Name), // Use name as ID
		Name: types.StringValue(h.Name),
	}

	model.Comment = types.StringPointerValue(h.Comment)
	model.CurrentAoe = types.StringPointerValue(h.CurrentAoe)
	model.CurrentEoe = types.StringPointerValue(h.CurrentEoe)
	model.InMultiNodeHost = types.Int32PointerValue(h.InMultiNodeHost)
	model.Mom = types.StringPointerValue(h.Mom)
	model.NoMultinodeJobs = types.BoolPointerValue(h.NoMultinodeJobs)
	model.Partition = types.StringPointerValue(h.Partition)
	model.PNames = types.StringPointerValue(h.PNames)
	model.Port = types.Int32PointerValue(h.Port)
	model.PowerOffEligible = types.BoolPointerValue(h.PowerOffEligible)
	model.PowerProvisioning = types.BoolPointerValue(h.PowerProvisioning)
	model.Priority = types.Int32PointerValue(h.Priority)
	model.ProvisionEnable = types.BoolPointerValue(h.ProvisionEnable)
	model.Queue = types.StringPointerValue(h.Queue)
	model.ResvEnable = types.BoolPointerValue(h.ResvEnable)

	// Only set ResourcesAvailable if there are actually resources to set
	// (excluding host and vnode which are PBS internal)
	hasResources := false
	for k := range h.ResourcesAvailable {
		if k != "host" && k != "vnode" {
			hasResources = true
			break
		}
	}

	if hasResources {
		model.ResourcesAvailable = make(map[string]types.String)
		for k, v := range h.ResourcesAvailable {
			// This is a bit hacky but the host and vnode attributes arent something you can
			// set on the resources_available so we don't want terraform getting confused about
			// what it's managing
			if k != "host" && k != "vnode" {
				model.ResourcesAvailable[k] = types.StringValue(v)
			}
		}
	}
	// If hasResources is false, leave ResourcesAvailable as nil (null in Terraform)

	return model
}
