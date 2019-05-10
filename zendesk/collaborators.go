package zendesk

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Collaborator struct {
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Collaborators struct {
	collaborators []interface{}
}

func (c *Collaborators) String() string {
	return fmt.Sprintf("%v", c.collaborators)
}

func (c *Collaborators) List() []interface{} {
	return c.collaborators
}

func (c *Collaborators) Append(i interface{}) error  {
	switch e := i.(type) {
	case string:
		c.collaborators = append(c.collaborators, e)
	case Collaborator:
		c.collaborators = append(c.collaborators, e)
	case int64:
		c.collaborators = append(c.collaborators, e)
	case map[string]interface{}:
		// This might be better suited in UnmarshalJSON
		collab := Collaborator{}
		name, ok := e["name"]
		if !ok {
			return fmt.Errorf("map %v did not contain a name value", e)
		}

		collab.Name, ok = name.(string)
		if !ok {
			return fmt.Errorf("type of name %v was not string", name)
		}

		email, ok := e["email"]
		if !ok {
			return fmt.Errorf("map %v did not contain an email value", e)
		}

		collab.Email, ok = email.(string)
		if !ok {
			return fmt.Errorf("type of email %v was not string", name)
		}
		c.collaborators = append(c.collaborators, collab)
	default:
		return fmt.Errorf("unsupported collaborator type %v", reflect.TypeOf(i))
	}

	return nil
}

func (c *Collaborators) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.collaborators)
}

func (c *Collaborators)  UnmarshalJSON(b []byte) error  {
	var tmpCollaborators []interface{}
	newCollaborators := Collaborators{}
	err :=  json.Unmarshal(b, &tmpCollaborators)
	if err != nil {
		return err
	}

	for _, i := range tmpCollaborators {
		var err error
		switch e := i.(type) {
		case float64:
			err = newCollaborators.Append(int64(e))
		default:
			err = newCollaborators.Append(i)
		}

		if err !=  nil {
			 return err
		}
	}


	// possibly validate that there aren't unexpected types in the slice
	c.collaborators = newCollaborators.List()
	return nil
}
