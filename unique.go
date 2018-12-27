package sliceutil

import (
	"reflect"
)

// Unique leaves only unique items in slice. The slice is modified in-place.
// Slice items must have comparable types.
// If slice is not a slice pointer, Unique panics.
func Unique(slicePtr interface{}) {
	sliceVal := reflect.ValueOf(slicePtr)
	if sliceVal.Kind() == reflect.Ptr {
		sliceVal = sliceVal.Elem()
	}
	mp := make(map[interface{}]bool)

	Filter(slicePtr, func(i int) bool {
		v := sliceVal.Index(i).Interface()
		exists, _ := mp[v]
		mp[v] = true
		return !exists
	})
}
