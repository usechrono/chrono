package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/usechrono/chrono/chrono"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: chrono run <workflow> [--key value]...")
		os.Exit(1)
	}

	workflowName := os.Args[2]

	// Dynamic params: parse --key value pairs
	params := make(map[string]interface{})
	fs := flag.NewFlagSet(workflowName, flag.ExitOnError)
	user := fs.String("user", "", "User name")
	email := fs.String("email", "", "Email address")
	fs.Parse(os.Args[3:])

	if *user != "" {
		params["user"] = *user
	}
	if *email != "" {
		params["email"] = *email
	}

	if wf, ok := chrono.GetWorkflow(workflowName); ok {
		wf.Run(params)
	} else {
		fmt.Printf("Workflow %s not found\n", workflowName)
	}
}
