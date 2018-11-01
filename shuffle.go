package sliceutil

import (
	"fmt"
	"math/rand"
	"reflect"
)

// Shuffle shuffles Go slice in-place.
// If slice is not a slice, Shuffle panics.
func Shuffle(slice interface{}) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(fmt.Sprintf("function called with non-slice value of type %T", slice))
	}
	swap := reflect.Swapper(slice)
	len := val.Len()

	for i := 0; i < len; i++ {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
