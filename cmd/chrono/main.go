package main

import (
	"fmt"
	"os"

	"github.com/usechrono/chrono/chrono"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: chrono run <workflow-name>")
		os.Exit(1)
	}

	cmd := os.Args[1]
	name := os.Args[2]

	// Example workflow for testing
	chrono.Register(&chrono.Workflow{
		Name: "send-welcome",
		Steps: []chrono.Step{
			{Name: "create-user-record", Run: func() error {
				fmt.Println("    -> creating user record...")
				return nil
			}},
			{Name: "send-email", Run: func() error {
				fmt.Println("    -> sending welcome email...")
				return nil
			}},
		},
	})

	switch cmd {
	case "run":
		wf, ok := chrono.GetWorkflow(name)
		if !ok {
			fmt.Printf("Workflow %q not found\n", name)
			os.Exit(1)
		}
		if err := wf.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}
}
