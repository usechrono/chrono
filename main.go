package main

import "fmt"

func main() {
	var version string = "v0.1" // explicit type
	count := 3                  // inferred type (shortcut)

	fmt.Println(version, count)

}
