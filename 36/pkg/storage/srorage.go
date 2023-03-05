package storage

import (
	"fmt"
	"reflect"
)

type Storage []interface{}

func New() Storage {
	var s []interface{}
	return s
}

func (s *Storage) Put(m interface{}) {
	*s = append(*s, m)
}

func (s Storage) Get() {
	for _, structure := range s {
		v := getValue(structure, "Text")
		fmt.Println(v)
	}
}

func getValue(structure interface{}, key string) interface{} {
	var result interface{}
	if reflect.TypeOf(structure).Elem().Kind() == reflect.Struct {
		elem := reflect.ValueOf(structure).Elem()
		field := elem.FieldByName(key)
		if field.IsValid() {
			result = field.Interface()
		}
	}
	return result
}
