package pbsclient

import (
	"testing"
)

func TestEscapeStringForQmgr(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		desc     string
	}{
		{
			input:    "simple_string",
			expected: `"simple_string"`,
			desc:     "simple string should use double quotes",
		},
		{
			input:    "string with spaces",
			expected: `"string with spaces"`,
			desc:     "string with spaces should use double quotes",
		},
		{
			input:    "string'with'single'quotes",
			expected: `"string'with'single'quotes"`,
			desc:     "string with single quotes should use double quotes",
		},
		{
			input:    `string"with"double"quotes`,
			expected: `'string"with"double"quotes'`,
			desc:     "string with double quotes should use single quotes",
		},
		{
			input:    "",
			expected: `""`,
			desc:     "empty string should use double quotes",
		},
		{
			input:    "normal-string_123",
			expected: `"normal-string_123"`,
			desc:     "alphanumeric with dashes/underscores should use double quotes",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := escapeStringForQmgr(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s for input %q", tc.expected, result, tc.input)
			}
		})
	}
}

func TestGenerateUpdateStringAttributeCommand(t *testing.T) {
	testCases := []struct {
		oldValue *string
		newValue *string
		expected []string
		desc     string
	}{
		{
			oldValue: nil,
			newValue: stringPtr("simple_value"),
			expected: []string{`/opt/pbs/bin/qmgr -c 'set queue test_queue enabled="simple_value"'`},
			desc:     "new string value with double quotes",
		},
		{
			oldValue: nil,
			newValue: stringPtr(`value"with"quotes`),
			expected: []string{`/opt/pbs/bin/qmgr -c 'set queue test_queue enabled='value"with"quotes''`},
			desc:     "new string value with double quotes should use single quotes",
		},
		{
			oldValue: nil,
			newValue: stringPtr("value'with'single'quotes"),
			expected: []string{`/opt/pbs/bin/qmgr -c 'set queue test_queue enabled="value'with'single'quotes"'`},
			desc:     "new string value with single quotes should use double quotes",
		},
		{
			oldValue: stringPtr("old_value"),
			newValue: nil,
			expected: []string{`/opt/pbs/bin/qmgr -c 'unset queue test_queue enabled'`},
			desc:     "unsetting a value",
		},
		{
			oldValue: stringPtr("old_value"),
			newValue: stringPtr("old_value"),
			expected: []string{},
			desc:     "no change should return empty slice",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := generateUpdateStringAttributeCommand("queue", "test_queue", "enabled", tc.oldValue, tc.newValue)
			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d commands, got %d", len(tc.expected), len(result))
				return
			}
			for i, cmd := range result {
				if cmd != tc.expected[i] {
					t.Errorf("Expected command %d to be %q, got %q", i, tc.expected[i], cmd)
				}
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
