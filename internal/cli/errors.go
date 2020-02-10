package cli

import "fmt"

type InvalidArgsError struct {
	InvalidArg string
}

func (e *InvalidArgsError) Error() string {
	return fmt.Sprintf("Invalid argument: %s\n", e.InvalidArg)
}