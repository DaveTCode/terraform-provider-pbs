package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type uniqueFlagsValidator struct{}

func (v uniqueFlagsValidator) Description(_ context.Context) string {
	return "each flag character must be unique"
}

func (v uniqueFlagsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v uniqueFlagsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	s := req.ConfigValue.String()
	seen := map[rune]bool{}
	for _, ch := range s {
		if seen[ch] {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path, v.Description(ctx), s,
			))
			return
		}
		seen[ch] = true
	}
}

// UniqueFlags returns a validator that ensures each character in the flag string is unique.
func UniqueFlags() validator.String {
	return uniqueFlagsValidator{}
}
