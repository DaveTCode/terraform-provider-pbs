package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// fakePrivateState is an in-memory implementation of the private state getter and
// setter interfaces used to unit test the one-shot unset tracking helpers.
type fakePrivateState struct {
	data map[string][]byte
}

func (f *fakePrivateState) GetKey(_ context.Context, key string) ([]byte, diag.Diagnostics) {
	return f.data[key], nil
}

func (f *fakePrivateState) SetKey(_ context.Context, key string, value []byte) diag.Diagnostics {
	if f.data == nil {
		f.data = map[string][]byte{}
	}
	f.data[key] = value
	return nil
}

func TestAppliedServerUnsets_RoundTrip(t *testing.T) {
	ctx := context.Background()
	private := &fakePrivateState{}

	if diags := writeAppliedServerUnsets(ctx, private, []string{"scheduler_iteration", "default_queue"}); diags.HasError() {
		t.Fatalf("unexpected diagnostics writing applied unsets: %v", diags)
	}

	applied, diags := readAppliedServerUnsets(ctx, private)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics reading applied unsets: %v", diags)
	}
	if len(applied) != 2 {
		t.Fatalf("expected 2 applied unsets, got %d", len(applied))
	}
	for _, name := range []string{"scheduler_iteration", "default_queue"} {
		if _, ok := applied[name]; !ok {
			t.Errorf("expected %q to be tracked as applied", name)
		}
	}
}

func TestAppliedServerUnsets_Prune(t *testing.T) {
	ctx := context.Background()
	private := &fakePrivateState{}

	// Initially track two, then rewrite with only one: the dropped name must be
	// re-armed (no longer reported as applied).
	if diags := writeAppliedServerUnsets(ctx, private, []string{"scheduler_iteration", "default_queue"}); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if diags := writeAppliedServerUnsets(ctx, private, []string{"scheduler_iteration"}); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	applied, diags := readAppliedServerUnsets(ctx, private)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if _, ok := applied["default_queue"]; ok {
		t.Errorf("expected default_queue to be re-armed after prune")
	}
	if _, ok := applied["scheduler_iteration"]; !ok {
		t.Errorf("expected scheduler_iteration to remain applied")
	}
}

func TestAppliedServerUnsets_EmptyAndNil(t *testing.T) {
	ctx := context.Background()

	// A nil private state reads as empty without error.
	applied, diags := readAppliedServerUnsets(ctx, nil)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if len(applied) != 0 {
		t.Errorf("expected no applied unsets from nil private state, got %d", len(applied))
	}

	// An unset private state (no key written) reads as empty.
	applied, diags = readAppliedServerUnsets(ctx, &fakePrivateState{})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if len(applied) != 0 {
		t.Errorf("expected no applied unsets from empty private state, got %d", len(applied))
	}
}

func TestIsUnsettableServerAttribute(t *testing.T) {
	unsettable := []string{"scheduler_iteration", "default_queue", "managers", "max_run_res", "acl_users"}
	for _, name := range unsettable {
		if !isUnsettableServerAttribute(name) {
			t.Errorf("expected %q to be unsettable", name)
		}
	}

	notUnsettable := []string{"id", "name", "unset_attributes", "acl_users_normalized", "acl_hosts_normalized"}
	for _, name := range notUnsettable {
		if isUnsettableServerAttribute(name) {
			t.Errorf("expected %q to not be unsettable", name)
		}
	}
}
