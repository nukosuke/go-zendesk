package zendesk

import "fmt"

type OptionsError struct {
	opts interface{}
}

func (e *OptionsError) Error() string {
	return fmt.Sprintf("invalid options: %v", e.opts)
}
