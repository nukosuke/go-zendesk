package zendesk

import (
	"encoding/json"
	"reflect"
	"testing"
)

const collaboratorListJSON = `[562,"someone@example.com",{"name":"SomeoneElse","email":"else@example.com"}]`

func TestCanBeUnmarshalled(t *testing.T) {
	c := &Collaborators{}
	err := c.UnmarshalJSON([]byte(collaboratorListJSON))
	if err != nil {
		t.Fatalf("Unmarshal returned an error %v", err)
	}

	list := c.List()
	if len(list) != 3 {
		t.Fatalf("Collaborators %v did not have the correct length when unmarshaled", c)
	}

	for _, i := range list {
		switch i.(type) {
		case string:
		case int64:
		case Collaborator:
			//do nothing
		default:
			t.Fatalf("List contained an unexpected type %v", reflect.TypeOf(i))
		}
	}
}

func TestCanBeMarshalled(t *testing.T) {
	c := &Collaborators{}
	err := c.UnmarshalJSON([]byte(collaboratorListJSON))
	if err != nil {
		t.Fatalf("Unmarshal returned an error %v", err)
	}

	out, err := c.MarshalJSON()
	if err != nil {
		t.Fatalf("Marshal returned an error %v", err)
	}

	if string(out) != collaboratorListJSON {
		t.Fatalf("Json output %s did not match expected output %s", out, collaboratorListJSON)
	}
}

func TestCanBeRemarshalled(t *testing.T) {
	var src, dst Collaborators
	marshalled, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("Marshalling returned an error %v", err)
	}
	err = json.Unmarshal(marshalled, &dst)
	if err != nil {
		t.Fatalf("Unmarshal returned an error %v", err)
	}

	if !reflect.DeepEqual(src, dst) {
		t.Fatalf("remarshalling is inconsistent")
	}
}
