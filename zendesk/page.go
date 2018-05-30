package zendesk

// Page is base struct for resource pagination
type Page struct {
	PreviousPage *string `json:"previous_page"`
	NextPage     *string `json:"next_page"`
	Count        int64   `json:"count"`
}

func (p Page) HasPrev() bool {
	return (p.NextPage != nil)
}

func (p Page) HasNext() bool {
	return (p.NextPage != nil)
}
