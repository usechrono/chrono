package chrono

// registry stores workflows by name
var registry = make(map[string]*Workflow)

// Register adds a workflow to the registry
func Register(w *Workflow) {
	registry[w.Name] = w
}

// GetWorkflow retrieves a workflow by name
func GetWorkflow(name string) (*Workflow, bool) {
	w, ok := registry[name]
	return w, ok
}

// List returns the names of all registered workflows
func List() []string {
	names := make([]string, 0, len(registry))
	for k := range registry {
		names = append(names, k)
	}
	return names
}
