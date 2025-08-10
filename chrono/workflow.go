package chrono

import (
	"fmt"
)

// Step represents a single action in a workflow
type Step struct {
	Name string
	Run  func() error
}

// Workflow represents a series of steps to execute
type Workflow struct {
	Name  string
	Steps []Step
}

// Run executes all steps in the workflow sequentially
func (w *Workflow) Run() error {
	fmt.Printf("Starting workflow: %s\n", w.Name)
	for i, s := range w.Steps {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(w.Steps), s.Name)
		if err := s.Run(); err != nil {
			return fmt.Errorf("step %q failed: %w", s.Name, err)
		}
	}
	fmt.Printf("Workflow %s completed\n", w.Name)
	return nil
}
