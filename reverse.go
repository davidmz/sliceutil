package sliceutil

import (
	"fmt"
	"reflect"
)

// Reverse reverts Go slice in-place.
// If slice is not a slice, Reverse panics.
func Reverse(slice interface{}) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(fmt.Sprintf("function called with non-slice value of type %T", slice))
	}
	swap := reflect.Swapper(slice)
	len := val.Len()

	for i := len/2 - 1; i >= 0; i-- {
		j := len - 1 - i
		swap(i, j)
	}
}
