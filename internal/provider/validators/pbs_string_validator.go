package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	pbsStringFieldRegex = regexp.MustCompile(`^[-_a-zA-Z0-9 !"#$%´()*+,./:;<=>?@[\\\]^_'{|}~]*$`)
)

type pbsStringValidator struct{}

func (v pbsStringValidator) Description(_ context.Context) string {
	return "Any character, including the space character. Only one of the two types of quote characters, \" or ', may appear in any given value. Values:[_a-zA-Z0-9][[-_a-zA-Z0-9 ! \" # $ % ´ ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ' { | } ~] ...] String resource values are case-sensitive.  No limit on length."
}

func (v pbsStringValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v pbsStringValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Empty strings are allowed
	if value == "" {
		return
	}

	// Check character pattern
	if !pbsStringFieldRegex.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid PBS String Field",
			fmt.Sprintf("String contains invalid characters for PBS field. "+
				"PBS string fields can only contain "+
				"[-_a-zA-Z0-9 !\"#$%%´()*+,./:;<=>?@[\\]^_'{|}~]. "+
				"Got: %s", value),
		)
		return
	}

	// Check quote character constraint - only one type of quote may appear
	hasDoubleQuote := strings.Contains(value, "\"")
	hasSingleQuote := strings.Contains(value, "'")

	if hasDoubleQuote && hasSingleQuote {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid PBS String Field",
			fmt.Sprintf("String cannot contain both single and double quotes. "+
				"Only one type of quote character (\" or ') may appear in any given value. "+
				"Got: %s", value),
		)
	}
}

// PbsString returns a validator that validates PBS string field requirements.
func PbsString() validator.String {
	return pbsStringValidator{}
}
