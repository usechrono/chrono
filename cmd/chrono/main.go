package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/usechrono/chrono/chrono"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: chrono run <workflow> [--key value]...")
		os.Exit(1)
	}

	workflowName := os.Args[2]

	params := make(map[string]interface{})

	// Parse --key value pairs from os.Args starting at index 3
	args := os.Args[3:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")
			if i+1 < len(args) {
				value := args[i+1]
				params[key] = value
				i++ // skip next since it's the value
			} else {
				fmt.Printf("Missing value for param %s\n", key)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Unexpected argument: %s\n", arg)
			os.Exit(1)
		}
	}

	if wf, ok := chrono.GetWorkflow(workflowName); ok {
		wf.Run(params)
	} else {
		fmt.Printf("Workflow %s not found\n", workflowName)
	}
}
