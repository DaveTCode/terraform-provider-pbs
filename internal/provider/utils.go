package provider

import (
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SetStringPointerIfNotNull sets a string pointer field if the types.String is not null.
func SetStringPointerIfNotNull(field types.String, target **string) {
	if !field.IsNull() {
		*target = field.ValueStringPointer()
	}
}

// SetBoolPointerIfNotNull sets a bool pointer field if the types.Bool is not null.
func SetBoolPointerIfNotNull(field types.Bool, target **bool) {
	if !field.IsNull() {
		*target = field.ValueBoolPointer()
	}
}

// SetInt32PointerIfNotNull sets an int32 pointer field if the types.Int32 is not null.
func SetInt32PointerIfNotNull(field types.Int32, target **int32) {
	if !field.IsNull() {
		*target = field.ValueInt32Pointer()
	}
}

// SetInt64PointerIfNotNull sets an int64 pointer field if the types.Int64 is not null.
func SetInt64PointerIfNotNull(field types.Int64, target **int64) {
	if !field.IsNull() {
		val := field.ValueInt64()
		*target = &val
	}
}

// ConvertTypesStringMap converts a map[string]types.String to map[string]string.
func ConvertTypesStringMap(source map[string]types.String) map[string]string {
	result := make(map[string]string)
	for k, v := range source {
		result[k] = v.ValueString()
	}
	return result
}

// ConvertTypesStringMapIfNotEmpty converts a map[string]types.String to map[string]string only if the source is not empty.
func ConvertTypesStringMapIfNotEmpty(source map[string]types.String, target *map[string]string) {
	if len(source) > 0 {
		*target = make(map[string]string)
		for k, v := range source {
			(*target)[k] = v.ValueString()
		}
	}
}

// ConvertTypesStringMapFiltered converts a map[string]types.String to map[string]string, excluding specified keys.
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

// normalizeCommaSeparatedString splits a comma-separated string, sorts the items, and rejoins them.
func normalizeCommaSeparatedString(value string) string {
	if value == "" {
		return ""
	}

	items := strings.Split(value, ",")
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}
	sort.Strings(items)
	return strings.Join(items, ",")
}

// AclFieldPair represents a pair of user field and normalized field for ACL preservation.
type AclFieldPair struct {
	UserField       types.String
	NormalizedField types.String
}

// preserveUserAclFormats preserves user-provided ACL field formats from plan.
func preserveUserAclFormats(planFields, resultFields []AclFieldPair) {
	if len(planFields) != len(resultFields) {
		return // Safety check
	}

	for i := range planFields {
		if !planFields[i].UserField.IsNull() {
			resultFields[i].UserField = planFields[i].UserField
		}
	}
}

// preserveUserAclFormatsFromState preserves user-provided ACL field formats from state when semantically equivalent.
func preserveUserAclFormatsFromState(stateFields, updatedFields []AclFieldPair) {
	if len(stateFields) != len(updatedFields) {
		return // Safety check
	}

	for i := range stateFields {
		if !stateFields[i].UserField.IsNull() && !updatedFields[i].NormalizedField.IsNull() {
			userFormat := stateFields[i].UserField.ValueString()
			pbsFormat := updatedFields[i].NormalizedField.ValueString()

			if normalizeCommaSeparatedString(userFormat) == normalizeCommaSeparatedString(pbsFormat) {
				updatedFields[i].UserField = stateFields[i].UserField
			}
		}
	}
}

// addNormalizedAclField sets the normalized version of an ACL field if the source is not nil.
func addNormalizedAclField(source *string, target *types.String) {
	if source != nil {
		*target = types.StringValue(normalizeCommaSeparatedString(*source))
	}
}

// convertStringMapToTypesStringMap converts a map[string]string to map[string]types.String.
func convertStringMapToTypesStringMap(source map[string]string) map[string]types.String {
	elements := make(map[string]types.String)
	for k, v := range source {
		elements[k] = types.StringValue(v)
	}
	return elements
}
