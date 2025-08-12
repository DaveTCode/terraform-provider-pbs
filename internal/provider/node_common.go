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
	PNames             types.String            `tfsdk:"pnames"`
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
		Name: m.Name.ValueString(),
	}

	// Set pointer fields using utility functions for null checking
	SetStringPointerIfNotNull(m.Comment, &node.Comment)
	SetStringPointerIfNotNull(m.CurrentAoe, &node.CurrentAoe)
	SetStringPointerIfNotNull(m.CurrentEoe, &node.CurrentEoe)
	SetInt32PointerIfNotNull(m.InMultiNodeHost, &node.InMultiNodeHost)
	SetStringPointerIfNotNull(m.Mom, &node.Mom)
	SetBoolPointerIfNotNull(m.NoMultinodeJobs, &node.NoMultinodeJobs)
	SetStringPointerIfNotNull(m.Partition, &node.Partition)
	SetStringPointerIfNotNull(m.PNames, &node.PNames)
	SetInt32PointerIfNotNull(m.Port, &node.Port)
	SetBoolPointerIfNotNull(m.PowerOffEligible, &node.PowerOffEligible)
	SetBoolPointerIfNotNull(m.PowerProvisioning, &node.PowerProvisioning)
	SetInt32PointerIfNotNull(m.Priority, &node.Priority)
	SetBoolPointerIfNotNull(m.ProvisionEnable, &node.ProvisionEnable)
	SetStringPointerIfNotNull(m.Queue, &node.Queue)
	SetBoolPointerIfNotNull(m.ResvEnable, &node.ResvEnable)

	// Convert ResourcesAvailable map, excluding 'host' and 'vnode' keys
	// This is a bit hacky but the host and vnode attributes aren't something you can
	// set on the resources_available so we don't want terraform getting confused about
	// what it's managing
	node.ResourcesAvailable = ConvertTypesStringMapFiltered(m.ResourcesAvailable, []string{"host", "vnode"})

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
