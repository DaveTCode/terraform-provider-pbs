package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SetStringPointerIfNotNull sets a string pointer field if the types.String is not null
func SetStringPointerIfNotNull(field types.String, target **string) {
	if !field.IsNull() {
		*target = field.ValueStringPointer()
	}
}

// SetBoolPointerIfNotNull sets a bool pointer field if the types.Bool is not null
func SetBoolPointerIfNotNull(field types.Bool, target **bool) {
	if !field.IsNull() {
		*target = field.ValueBoolPointer()
	}
}

// SetInt32PointerIfNotNull sets an int32 pointer field if the types.Int32 is not null
func SetInt32PointerIfNotNull(field types.Int32, target **int32) {
	if !field.IsNull() {
		*target = field.ValueInt32Pointer()
	}
}

// SetInt64PointerIfNotNull sets an int64 pointer field if the types.Int64 is not null
func SetInt64PointerIfNotNull(field types.Int64, target **int64) {
	if !field.IsNull() {
		*target = field.ValueInt64Pointer()
	}
}

// ConvertTypesStringMap converts a map[string]types.String to map[string]string
func ConvertTypesStringMap(source map[string]types.String) map[string]string {
	result := make(map[string]string)
	for k, v := range source {
		result[k] = v.ValueString()
	}
	return result
}

// ConvertTypesStringMapIfNotEmpty converts a map[string]types.String to map[string]string only if the source is not empty
func ConvertTypesStringMapIfNotEmpty(source map[string]types.String, target *map[string]string) {
	if len(source) > 0 {
		*target = make(map[string]string)
		for k, v := range source {
			(*target)[k] = v.ValueString()
		}
	}
}

// ConvertTypesStringMapFiltered converts a map[string]types.String to map[string]string, excluding specified keys
func ConvertTypesStringMapFiltered(source map[string]types.String, excludeKeys []string) map[string]string {
	result := make(map[string]string)
	excludeSet := make(map[string]bool)
	for _, key := range excludeKeys {
		excludeSet[key] = true
	}

	for k, v := range source {
		if !excludeSet[k] {
			result[k] = v.ValueString()
		}
	}
	return result
}
