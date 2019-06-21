// Package sliceutil is a set of tiny but useful functions designed for in-place slice operations.
package sliceutil

import (
	"fmt"
	"reflect"
)

// Filter filters Go slice in-place.
// If slice is not a slice pointer, Filter panics.
// Custom disposers of the removing items can be used, see description
// of the Disposer type.
func Filter(slicePtr interface{}, predicate func(int) bool, disposers ...Disposer) {
	val := reflect.ValueOf(slicePtr)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		panic(fmt.Sprintf("function called with non-slice pointer value of type %T", slicePtr))
	}

	sliceVal := val.Elem()
	swap := reflect.Swapper(sliceVal.Interface())
	len := sliceVal.Len()

	last := 0
	for i := 0; i < len; i++ {
		if predicate(i) {
			if last < i {
				swap(i, last)
			}
			last++
			continue
		}
	}

	zero := reflect.Zero(sliceVal.Type().Elem())
	for i := last; i < len; i++ {
		for _, disposer := range disposers {
			disposer(i)
		}
		sliceVal.Index(i).Set(zero)
	}

	sliceVal.SetLen(last)
	sliceVal.SetCap(last)
}
