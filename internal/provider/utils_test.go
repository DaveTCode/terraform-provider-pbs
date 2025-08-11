package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestConvertTypesStringMap(t *testing.T) {
	// Test simple conversion
	source := map[string]types.String{
		"key1": types.StringValue("value1"),
		"key2": types.StringValue("value2"),
	}

	result := ConvertTypesStringMap(source)

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result["key1"] != "value1" {
		t.Errorf("Expected 'value1', got %s", result["key1"])
	}

	if result["key2"] != "value2" {
		t.Errorf("Expected 'value2', got %s", result["key2"])
	}
}

func TestConvertTypesStringMapIfNotEmpty(t *testing.T) {
	// Test empty map - should not modify target
	var target map[string]string
	source := map[string]types.String{}

	ConvertTypesStringMapIfNotEmpty(source, &target)

	if target != nil {
		t.Errorf("Expected target to remain nil for empty source")
	}

	// Test non-empty map
	source = map[string]types.String{
		"key1": types.StringValue("value1"),
	}

	ConvertTypesStringMapIfNotEmpty(source, &target)

	if target == nil {
		t.Errorf("Expected target to be initialized")
	}

	if len(target) != 1 {
		t.Errorf("Expected 1 item, got %d", len(target))
	}

	if target["key1"] != "value1" {
		t.Errorf("Expected 'value1', got %s", target["key1"])
	}
}

func TestConvertTypesStringMapFiltered(t *testing.T) {
	source := map[string]types.String{
		"key1":  types.StringValue("value1"),
		"host":  types.StringValue("hostvalue"),
		"vnode": types.StringValue("vnodevalue"),
		"key2":  types.StringValue("value2"),
	}

	result := ConvertTypesStringMapFiltered(source, []string{"host", "vnode"})

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result["key1"] != "value1" {
		t.Errorf("Expected 'value1', got %s", result["key1"])
	}

	if result["key2"] != "value2" {
		t.Errorf("Expected 'value2', got %s", result["key2"])
	}

	if _, exists := result["host"]; exists {
		t.Errorf("Expected 'host' to be filtered out")
	}

	if _, exists := result["vnode"]; exists {
		t.Errorf("Expected 'vnode' to be filtered out")
	}
}

func TestSetStringPointerIfNotNull(t *testing.T) {
	var target *string

	// Test null value - should not modify target
	nullField := types.StringNull()
	SetStringPointerIfNotNull(nullField, &target)

	if target != nil {
		t.Errorf("Expected target to remain nil for null field")
	}

	// Test non-null value
	nonNullField := types.StringValue("test")
	SetStringPointerIfNotNull(nonNullField, &target)

	if target == nil {
		t.Errorf("Expected target to be set")
	}

	if *target != "test" {
		t.Errorf("Expected 'test', got %s", *target)
	}
}
