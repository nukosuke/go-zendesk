// sideload Allows for sideload support when calling the zendesk api.
// For more information about sideloading see: https://developer.zendesk.com/rest_api/docs/support/side_loading
package sideload

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
)

type SideLoader interface {
	Key() string
}

type ExtraObjectSideloader interface {
	SideLoader
	AppendToStruct([]reflect.StructField) []reflect.StructField
	IsAssignable() bool
	SetValue(reflect.Value)
}

type simpleInclude string

func (s simpleInclude) Key() string {
	return string(s)
}

// In some instances sideloading results in extra fields being added to the main object returned by the api
// The Include sideloader should be used in such cases
func Include(key string) SideLoader  {
	return simpleInclude(key)
}

type sideLoader struct {
	value interface{}
	key string
}

func (s *sideLoader) Key() string {
	return s.key
}

func (s *sideLoader) AppendToStruct(sf []reflect.StructField) []reflect.StructField {
	return append(sf, reflect.StructField{
		Name: strcase.ToCamel(s.key),
		Tag:  reflect.StructTag(fmt.Sprintf("json:%s", s.key)),
		Type: reflect.TypeOf(s.value),
	})
}

func (s *sideLoader) SetValue(v reflect.Value) {
	v.Elem().FieldByName(strcase.ToCamel(s.key)).Set(reflect.ValueOf(s.value))
}

func (s *sideLoader) IsAssignable() bool {
	return reflect.ValueOf(s.value).Kind() == reflect.Ptr
}

// In some instances sideloading results in extra objects being added to the api response.
// The IncludeObject sideloader should be used in such cases
func IncludeObject(key string, v interface{}) SideLoader {
	return &sideLoader{
		value: v,
		key:   key,
	}
}



