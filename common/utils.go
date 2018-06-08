package common

import (
	"reflect"
)

func Struct2Map(obj interface{}) (data map[string] interface{}, err error) {
    data = make(map[string] interface{})
    keys := reflect.TypeOf(obj)
    values := reflect.ValueOf(obj)
    for i := 0; i < keys.NumField(); i++ {
      data[keys.Field(i).Name] = values.Field(i).Interface()
    }
    err = nil
    return
}
