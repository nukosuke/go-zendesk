package zendesk

import "fmt"

// OptionsError is an error type for invalid option argument.
type OptionsError struct {
	opts interface{}
}

func (e *OptionsError) Error() string {
	return fmt.Sprintf("invalid options: %v", e.opts)
}
