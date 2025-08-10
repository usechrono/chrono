package chrono

import (
	"fmt"
)

// Step represents a single action in a workflow
type Step struct {
	Name string
	Run  func(params map[string]interface{}) error
}

// Workflow represents a series of steps to execute
type Workflow struct {
	Name  string
	Steps []Step
}

// Run executes all steps in the workflow sequentially
func (w *Workflow) Run(params map[string]interface{}) {
	fmt.Printf("Starting workflow: %s\n", w.Name)
	for i, step := range w.Steps {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(w.Steps), step.Name)
		if err := step.Run(params); err != nil {
			fmt.Printf("    ❌ Error: %v\n", err)
			break
		}
	}
	fmt.Printf("Workflow %s completed\n", w.Name)
}
