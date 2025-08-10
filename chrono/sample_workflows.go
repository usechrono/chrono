package chrono

import "fmt"

// init registers some test workflows automatically
func init() {
	Register(&Workflow{
		Name: "welcome",
		Steps: []Step{
			{
				Name: "Greet User",
				Run: func(params map[string]interface{}) error {
					user, ok := params["user"].(string)
					if !ok || user == "" {
						user = "Guest"
					}
					fmt.Printf("Hello, %s! Welcome to Chrono.\n", user)
					return nil
				},
			},
		},
	})

	Register(&Workflow{
		Name: "validate",
		Steps: []Step{
			{
				Name: "Check Email Param",
				Run: func(params map[string]interface{}) error {
					email, ok := params["email"].(string)
					if !ok || email == "" {
						return fmt.Errorf("email param is required")
					}
					fmt.Printf("Email provided: %s\n", email)
					return nil
				},
			},
		},
	})
}
