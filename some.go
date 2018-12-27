package sliceutil

import (
	"fmt"
	"reflect"
)

// Some finds the first item in slice which satisfies the predicate
// and returns true if it was found or false otherwise.
// If slice is not a slice or an array, Some panics.
func Some(slice interface{}, predicate func(int) bool) bool {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice && sliceVal.Kind() != reflect.Array {
		panic(fmt.Sprintf("function called with non-slice value of type %T", slice))
	}

	len := sliceVal.Len()
	for i := 0; i < len; i++ {
		if predicate(i) {
			return true
		}
	}
	return false
}
