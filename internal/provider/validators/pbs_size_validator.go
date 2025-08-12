package provider

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	// PBS size validation regex
	// Format: <integer>[<suffix>]
	// Suffix can be: b, w, kb, kw, mb, mw, gb, gw, tb, tw, pb, pw (case insensitive).
	pbsSizeRegex = regexp.MustCompile(`^(\d+)([bw]|[kmgtpKMGTP][bwBW])?$`)
)

// pbsSizeValidator validates that a string conforms to PBS size attribute requirements.
type pbsSizeValidator struct{}

// Description returns a description of the validator suitable for logging and error messages.
func (v pbsSizeValidator) Description(_ context.Context) string {
	return "value must be a valid PBS size attribute"
}

// MarkdownDescription returns a markdown description of the validator suitable for documentation.
func (v pbsSizeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString performs the validation.
func (v pbsSizeValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Empty strings are allowed (defaults to bytes)
	if value == "" {
		return
	}

	// Check basic format
	if !pbsSizeRegex.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid PBS Size Attribute",
			fmt.Sprintf("Size must be in format <integer>[<suffix>] where suffix can be one of: "+
				"b, w (bytes/words), kb, kw (kilobytes/kilowords), mb, mw (megabytes/megawords), "+
				"gb, gw (gigabytes/gigawords), tb, tw (terabytes/terawords), pb, pw (petabytes/petawords). "+
				"Got: %s", value),
		)
		return
	}

	// Parse the value to validate the integer part
	if err := validatePbsSizeFormat(value); err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid PBS Size Attribute",
			fmt.Sprintf("Invalid size format: %s. %s", value, err.Error()),
		)
	}
}

// validatePbsSizeFormat validates the PBS size format more thoroughly.
func validatePbsSizeFormat(value string) error {
	// Normalize to lowercase for suffix checking
	lowerValue := strings.ToLower(value)

	// Find where the suffix starts (first non-digit character)
	suffixStart := -1
	for i, r := range value {
		if r < '0' || r > '9' {
			suffixStart = i
			break
		}
	}

	var integerPart string
	var suffix string

	if suffixStart == -1 {
		// No suffix, just integer
		integerPart = value
		suffix = ""
	} else {
		integerPart = value[:suffixStart]
		suffix = lowerValue[suffixStart:]
	}

	// Validate integer part
	if integerPart == "" {
		return fmt.Errorf("missing integer part")
	}

	// Parse integer to ensure it's valid
	intVal, err := strconv.ParseInt(integerPart, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid integer part: %s", integerPart)
	}

	// Integer must be non-negative
	if intVal < 0 {
		return fmt.Errorf("size must be non-negative, got: %d", intVal)
	}

	// Validate suffix if present
	if suffix != "" {
		validSuffixes := []string{"b", "w", "kb", "kw", "mb", "mw", "gb", "gw", "tb", "tw", "pb", "pw"}
		valid := slices.Contains(validSuffixes, suffix)
		if !valid {
			return fmt.Errorf("invalid suffix '%s', must be one of: %s", suffix, strings.Join(validSuffixes, ", "))
		}
	}

	return nil
}

// PbsSize returns a validator that validates PBS size attribute requirements.
func PbsSize() validator.String {
	return pbsSizeValidator{}
}
