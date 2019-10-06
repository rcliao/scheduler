package processor

import "fmt"

// ProcessFunc defines the func type for process func
type ProcessFunc func()

// Debug func prints the data into console
func Debug(data string) ProcessFunc {
	return func() {
		fmt.Printf("DEBUG func got data: %s", data)
	}
}
