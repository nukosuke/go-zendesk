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
	value interface{}
	key string
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

func Include(key string, path string, v interface{}) SideLoader {
	return &sideLoader{
		value: v,
		key:   key,
		jsonPath: path,
	}
}
