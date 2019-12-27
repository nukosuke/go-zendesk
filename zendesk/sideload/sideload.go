// sideload Allows for sideload support when calling the zendesk api.
// For more information about sideloading see: https://developer.zendesk.com/rest_api/docs/support/side_loading
package sideload

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

type SideLoader interface {
	Key() string
	Unmarshal([]byte) error
}

type sideLoader struct {
	value    interface{}
	key      string
	jsonPath string
}

func (s *sideLoader) Key() string {
	return s.key
}

func (s *sideLoader) Unmarshal(b []byte) error {
	v := gjson.ParseBytes(b).Get(s.jsonPath)
	if !v.Exists() {
		return fmt.Errorf("could not find %s in result %s", s.jsonPath, string(b))
	}

	return json.Unmarshal([]byte(v.Raw), s.value)
}

// In some instances sideloading results in extra objects being added to the api response.
// The IncludeObject sideloader should be used in such cases
func IncludeObject(key string, v interface{}) SideLoader {
	return Include(key, key, v)
}

// Include is a function used to initialize a Sideloader.
// it takes 3 parameters
// key: the string to be included in the query string
// path: where the object can be found in the result body for more information see https://github.com/tidwall/gjson
// v: an object that the result can be unmarshalled into.
func Include(key string, path string, v interface{}) SideLoader {
	return &sideLoader{
		value:    v,
		key:      key,
		jsonPath: path,
	}
}
