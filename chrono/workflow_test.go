package chrono

import (
	"testing"
)

func TestWelcomeWorkflow(t *testing.T) {
	wf, ok := GetWorkflow("welcome")
	if !ok {
		t.Fatal("welcome workflow not found")
	}

	params := map[string]interface{}{
		"user": "Tester",
	}

	// We just want to ensure Run does not error out
	// Since Run doesn't return error, no error to catch here.
	wf.Run(params)
}

func TestValidateWorkflow(t *testing.T) {
	wf, ok := GetWorkflow("validate")
	if !ok {
		t.Fatal("validate workflow not found")
	}

	// Test case 1: Missing email param - expect error output
	params1 := map[string]interface{}{}
	wf.Run(params1)

	// Test case 2: Provided email param - should succeed
	params2 := map[string]interface{}{
		"email": "test@example.com",
	}
	wf.Run(params2)
}
