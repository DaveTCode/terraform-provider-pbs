package provider

import (
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SetStringPointerIfNotNull sets a string pointer field if the types.String is a
// known, non-null value. Unknown values are skipped so that attributes flagged for
// unset (whose planned value is unknown) leave the target nil and produce a qmgr unset.
func SetStringPointerIfNotNull(field types.String, target **string) {
	if !field.IsNull() && !field.IsUnknown() {
		*target = field.ValueStringPointer()
	}
}

// SetBoolPointerIfNotNull sets a bool pointer field if the types.Bool is a known, non-null value.
func SetBoolPointerIfNotNull(field types.Bool, target **bool) {
	if !field.IsNull() && !field.IsUnknown() {
		*target = field.ValueBoolPointer()
	}
}

// SetInt32PointerIfNotNull sets an int32 pointer field if the types.Int32 is a known, non-null value.
func SetInt32PointerIfNotNull(field types.Int32, target **int32) {
	if !field.IsNull() && !field.IsUnknown() {
		*target = field.ValueInt32Pointer()
	}
}

// SetInt64PointerIfNotNull sets an int64 pointer field if the types.Int64 is a known, non-null value.
func SetInt64PointerIfNotNull(field types.Int64, target **int64) {
	if !field.IsNull() && !field.IsUnknown() {
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
		if !planFields[i].UserField.IsNull() && !planFields[i].UserField.IsUnknown() {
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
		if !stateFields[i].UserField.IsNull() && !stateFields[i].UserField.IsUnknown() &&
			!updatedFields[i].NormalizedField.IsNull() && !updatedFields[i].NormalizedField.IsUnknown() {
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
