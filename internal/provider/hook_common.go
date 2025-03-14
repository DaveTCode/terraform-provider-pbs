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

func (m pbsHookModel) ToPbsHook() pbsclient.PbsHook {
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

	model.Type = types.StringPointerValue(h.Type)
	model.Alarm = types.Int32PointerValue(h.Alarm)
	model.Debug = types.BoolPointerValue(h.Debug)
	model.Enabled = types.BoolPointerValue(h.Enabled)
	model.Event = types.StringPointerValue(h.Event)
	model.FailAction = types.StringPointerValue(h.FailAction)
	model.Freq = types.Int32PointerValue(h.Freq)
	model.Order = types.Int32PointerValue(h.Order)
	model.User = types.StringPointerValue(h.User)

	return model
}
