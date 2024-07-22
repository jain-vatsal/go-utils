package utils

import (
	"reflect"
)

func IsListSorted(list interface{}) bool {
	v := reflect.ValueOf(list)

	if v.Kind() != reflect.Slice {
		panic("IsListSorted: input is not a slice")
	}

	length := v.Len()
	if length <= 1 {
		return true
	}

	for i := 1; i < length; i++ {
		if !less(v.Index(i-1), v.Index(i)) {
			return false
		}
	}
	return true
}

func less(a, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() <= b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() <= b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() <= b.Float()
	case reflect.String:
		return a.String() <= b.String()
	default:
		panic("IsListSorted: unsupported element type")
	}
}
