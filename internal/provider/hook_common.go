package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pbsHookModel struct {
	Name       types.String `tfsdk:"name"`
	Type       types.String `tfsdk:"type"`
	Alarm      types.Int32  `tfsdk:"alarm"`
	Debug      types.Bool   `tfsdk:"debug"`
	Enabled    types.Bool   `tfsdk:"enabled"`
	Event      types.String `tfsdk:"event"`
	FailAction types.String `tfsdk:"fail_action"`
	Freq       types.Int32  `tfsdk:"freq"`
	Order      types.Int32  `tfsdk:"order"`
	User       types.String `tfsdk:"user"`
}

func (m pbsHookModel) ToPbsResource() pbsclient.PbsHook {
	return pbsclient.PbsHook{
		Name:       m.Name.ValueString(),
		Type:       m.Type.ValueStringPointer(),
		Alarm:      m.Alarm.ValueInt32Pointer(),
		Debug:      m.Debug.ValueBoolPointer(),
		Enabled:    m.Enabled.ValueBoolPointer(),
		Event:      m.Event.ValueStringPointer(),
		FailAction: m.FailAction.ValueStringPointer(),
		Freq:       m.Freq.ValueInt32Pointer(),
		Order:      m.Order.ValueInt32Pointer(),
		User:       m.User.ValueStringPointer(),
	}
}

func createPbsHookModel(h pbsclient.PbsHook) pbsHookModel {
	model := pbsHookModel{
		Name: types.StringValue(h.Name),
	}

	if h.Type != nil {
		model.Type = types.StringValue(*h.Type)
	}
	if h.Alarm != nil {
		model.Alarm = types.Int32Value(*h.Alarm)
	}
	if h.Debug != nil {
		model.Debug = types.BoolValue(*h.Debug)
	}
	if h.Enabled != nil {
		model.Enabled = types.BoolValue(*h.Enabled)
	}
	if h.Event != nil {
		model.Event = types.StringValue(*h.Event)
	}
	if h.FailAction != nil {
		model.FailAction = types.StringValue(*h.FailAction)
	}
	if h.Freq != nil {
		model.Freq = types.Int32Value(*h.Freq)
	}
	if h.Order != nil {
		model.Order = types.Int32Value(*h.Order)
	}
	if h.User != nil {
		model.User = types.StringValue(*h.User)
	}

	return model
}
