package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPbsSizeValidator_ValidCases(t *testing.T) {
	testValidator := PbsSize()
	ctx := context.Background()

	validCases := []string{
		"",          // empty string (defaults to bytes)
		"0",         // zero
		"1",         // just integer (defaults to bytes)
		"1024",      // integer only
		"1b",        // bytes
		"1w",        // words
		"1kb",       // kilobytes
		"1kw",       // kilowords
		"1mb",       // megabytes
		"1mw",       // megawords
		"1gb",       // gigabytes
		"1gw",       // gigawords
		"1tb",       // terabytes
		"1tw",       // terawords
		"1pb",       // petabytes
		"1pw",       // petawords
		"1KB",       // uppercase suffix
		"1GB",       // uppercase suffix
		"1024mb",    // larger numbers
		"999999999", // large integer
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
				t.Errorf("Expected valid size '%s' to pass validation, but got errors: %v", testCase, resp.Diagnostics)
			}
		})
	}
}

func TestPbsSizeValidator_InvalidCases(t *testing.T) {
	testValidator := PbsSize()
	ctx := context.Background()

	invalidCases := []struct {
		input string
		desc  string
	}{
		{"-1", "negative number"},
		{"-5mb", "negative with suffix"},
		{"abc", "non-numeric"},
		{"1.5gb", "decimal number"},
		{"1xyz", "invalid suffix"},
		{"1byte", "invalid suffix (too long)"},
		{"1k", "incomplete suffix"},
		{"gb", "missing integer"},
		{"1 gb", "space in value"},
		{"1_gb", "underscore in value"},
	}

	for _, testCase := range invalidCases {
		t.Run("invalid_"+testCase.input+"_"+testCase.desc, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("test"),
				ConfigValue: types.StringValue(testCase.input),
			}
			resp := &validator.StringResponse{}

			testValidator.ValidateString(ctx, req, resp)

			if !resp.Diagnostics.HasError() {
				t.Errorf("Expected invalid size '%s' (%s) to fail validation, but it passed", testCase.input, testCase.desc)
			}
		})
	}
}

func TestPbsSizeValidator_Description(t *testing.T) {
	testValidator := PbsSize()
	ctx := context.Background()

	desc := testValidator.Description(ctx)
	if desc == "" {
		t.Error("Description should not be empty")
	}

	markdownDesc := testValidator.MarkdownDescription(ctx)
	if markdownDesc == "" {
		t.Error("MarkdownDescription should not be empty")
	}
}
