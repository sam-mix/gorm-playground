package util

import (
	"reflect"
)

func ConvertStructWithTags(i interface{}, tags map[string]reflect.StructTag) interface{} {
	value := reflect.Indirect(reflect.ValueOf(i))

	t := value.Type()
	fields := make([]reflect.StructField, 0)
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i))
		if tag, ok := tags[t.Field(i).Name]; ok {
			fields[i].Tag = tag
		}
	}

	newType := reflect.StructOf(fields)
	newValue := value.Convert(newType)
	return newValue.Interface()
}
