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
	if !m.Type.IsNull() {
		hook.Type = m.Type.ValueStringPointer()
	}
	if !m.Alarm.IsNull() {
		hook.Alarm = m.Alarm.ValueInt32Pointer()
	}
	if !m.Debug.IsNull() {
		hook.Debug = m.Debug.ValueBoolPointer()
	}
	if !m.Enabled.IsNull() {
		hook.Enabled = m.Enabled.ValueBoolPointer()
	}
	if !m.Event.IsNull() {
		hook.Event = m.Event.ValueStringPointer()
	}
	if !m.FailAction.IsNull() {
		hook.FailAction = m.FailAction.ValueStringPointer()
	}
	if !m.Freq.IsNull() {
		hook.Freq = m.Freq.ValueInt32Pointer()
	}
	if !m.Order.IsNull() {
		hook.Order = m.Order.ValueInt32Pointer()
	}
	if !m.User.IsNull() {
		hook.User = m.User.ValueStringPointer()
	}

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
