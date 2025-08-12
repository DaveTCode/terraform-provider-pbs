package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPbsStringValidator_ValidCases(t *testing.T) {
	testValidator := PbsString()
	ctx := context.Background()

	validCases := []string{
		"",                       // empty string
		"validstring",            // basic valid string
		"_underscore_start",      // underscore at start
		"test123",                // alphanumeric
		"test-string_with.chars", // allowed special chars
		"string'with'quotes",     // single quotes only
		`string"with"quotes`,     // double quotes only
		" starts_with_space",     // starts with space (now allowed)
		"-starts_with_dash",      // starts with dash (now allowed)
		".starts_with_dot",       // starts with dot (now allowed)
	}

	for _, testCase := range validCases {
		t.Run("valid_"+testCase, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("test"),
				ConfigValue: types.StringValue(testCase),
			}
			resp := &validator.StringResponse{}

			testValidator.ValidateString(ctx, req, resp)

			if resp.Diagnostics.HasError() {
				t.Errorf("Expected valid string '%s' to pass validation, but got errors: %v", testCase, resp.Diagnostics)
			}
		})
	}
}

func TestPbsStringValidator_InvalidCases(t *testing.T) {
	testValidator := PbsString()
	ctx := context.Background()

	invalidCases := []string{
		`both'and"quotes`,       // both quote types
		"string\nwith\nnewline", // newline character
		"string\twith\ttab",     // tab character
	}

	for _, testCase := range invalidCases {
		t.Run("invalid_"+testCase, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("test"),
				ConfigValue: types.StringValue(testCase),
			}
			resp := &validator.StringResponse{}

			testValidator.ValidateString(ctx, req, resp)

			if !resp.Diagnostics.HasError() {
				t.Errorf("Expected invalid string '%s' to fail validation, but it passed", testCase)
			}
		})
	}
}
