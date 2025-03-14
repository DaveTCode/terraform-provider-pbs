package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pbsResourceModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
	Flag types.String `tfsdk:"flag"`
}

func (m pbsResourceModel) ToPbsResource() pbsclient.PbsResource {
	return pbsclient.PbsResource{
		Name: m.Name.ValueString(),
		Type: m.Type.ValueString(),
		Flag: m.Flag.ValueStringPointer(),
	}
}

func createPbsResoureModel(r pbsclient.PbsResource) pbsResourceModel {
	model := pbsResourceModel{
		Name: types.StringValue(r.Name),
		Type: types.StringValue(r.Type),
		Flag: types.StringPointerValue(r.Flag),
	}

	return model
}
