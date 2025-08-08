package provider

import (
	"strings"
	"testing"
)

func TestTestAccResourceName(t *testing.T) {
	// Test basic functionality
	name1 := testAccResourceName("test")
	name2 := testAccResourceName("test")
	
	// Names should be different
	if name1 == name2 {
		t.Errorf("Expected different names, got same: %s", name1)
	}
	
	// Names should have correct format
	if !strings.Contains(name1, "test_") {
		t.Errorf("Expected name to contain prefix, got: %s", name1)
	}
	
	// Test length constraints
	longPrefix := "verylongprefixname"
	longName := testAccResourceName(longPrefix)
	if len(longName) > 15 {
		t.Errorf("Name too long for PBS: %s (length: %d)", longName, len(longName))
	}
	
	// Test that it produces reasonable output
	for i := 0; i < 10; i++ {
		name := testAccResourceName("hook")
		t.Logf("Generated name %d: %s", i, name)
		if len(name) > 15 {
			t.Errorf("Name %s is too long (%d chars)", name, len(name))
		}
		if !strings.HasPrefix(name, "hook_") {
			t.Errorf("Name %s doesn't have expected prefix", name)
		}
	}
}
