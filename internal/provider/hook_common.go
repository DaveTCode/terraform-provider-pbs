package provider

import (
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pbsHookModel struct {
	ID         types.String `tfsdk:"id"`
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
	hook := pbsclient.PbsHook{
		Name: m.Name.ValueString(),
	}

	// Only set pointer fields if the value is not null
	SetStringPointerIfNotNull(m.Type, &hook.Type)
	SetInt32PointerIfNotNull(m.Alarm, &hook.Alarm)
	SetBoolPointerIfNotNull(m.Debug, &hook.Debug)
	SetBoolPointerIfNotNull(m.Enabled, &hook.Enabled)
	SetStringPointerIfNotNull(m.Event, &hook.Event)
	SetStringPointerIfNotNull(m.FailAction, &hook.FailAction)
	SetInt32PointerIfNotNull(m.Freq, &hook.Freq)
	SetInt32PointerIfNotNull(m.Order, &hook.Order)
	SetStringPointerIfNotNull(m.User, &hook.User)

	return hook
}

func createPbsHookModel(h pbsclient.PbsHook) pbsHookModel {
	model := pbsHookModel{
		ID:   types.StringValue(h.Name), // Use name as ID
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
