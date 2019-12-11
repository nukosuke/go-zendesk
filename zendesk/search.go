package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
)

// SearchOptions are the options that can be provided to the search API
//
// ref: https://developer.zendesk.com/rest_api/docs/support/search#available-parameters
type SearchOptions struct {
	PageOptions
	Query     string `url:"query"`
	SortBy    string `url:"sort_by,omitempty"`
	SortOrder string `url:"sort_order,omitempty"`
}

type SearchAPI interface {
	Search(ctx context.Context, opts *SearchOptions) (SearchResults, Page, error)
}

type SearchResults struct {
	results []interface{}
}

func (r *SearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.results)
}

func (r *SearchResults) UnmarshalJSON(b []byte) error  {
	var (
		results []interface{}
		tmp     []json.RawMessage
	)

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	for _, v := range tmp {
		// TODO: check with @nukosuke to see if we should try to speed this up using github.com/tidwall/gjson
		m := make(map[string]interface{})
		err := json.Unmarshal(v, &m)
		if err != nil {
			return err
		}

		t, ok := m["result_type"].(string)
		if !ok {
			return fmt.Errorf("could not assert result type to string. json was: %v", v)
		}

		var value interface{}

		switch t {
		case "group":
			var g Group
			err = json.Unmarshal(v, &g)
			value = g
		case "ticket":
			var t Ticket
			err = json.Unmarshal(v, &t)
			value = t
		case "user":
			var u User
			err = json.Unmarshal(v, &u)
			value = u
		case "organization":
			var o Organization
			err = json.Unmarshal(v, &o)
			value = o
		case "topic":
			var t Topic
			err = json.Unmarshal(v, &t)
			value = t
		default:
			err = fmt.Errorf("value of result was an unsupported type %s", t)
		}

		if err != nil {
			return err
		}

		results= append(results, value)
	}

	r.results = results
	return nil
}

// GetTriggers fetch trigger list
//
// ref: https://developer.zendesk.com/rest_api/docs/support/triggers#getting-triggers
func (z *Client) Search(ctx context.Context, opts *SearchOptions) (SearchResults, Page, error) {
	var data struct {
		Results SearchResults `json:"results"`
		Page
	}

	if opts == nil {
		return SearchResults{}, Page{}, &OptionsError{opts}
	}

	u, err := addOptions("/search.json", opts)
	if err != nil {
		return SearchResults{}, Page{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return SearchResults{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return SearchResults{}, Page{}, err
	}

	return data.Results, data.Page, nil
}

// String return string formatted for Search results
func (r *SearchResults) String() string {
	return fmt.Sprintf("%v", r.results)
}

// List return internal array in Search Results
func (r *SearchResults) List() []interface{} {
	return r.results
}
