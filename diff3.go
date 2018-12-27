package sliceutil

import (
	"fmt"
	"reflect"
)

// Diff3 calculates difference between two slices and return three
// new slices: onlyA with items belongs only to A slice, AandB with items common to A and B
// and onlyB with items belongs only to B. A and B must be slices of the same type and
// their items must have comparable types.
func Diff3(sliceA, sliceB interface{}) (onlyA, AandB, onlyB interface{}) {
	aVal := reflect.ValueOf(sliceA)
	bVal := reflect.ValueOf(sliceB)
	if aVal.Kind() != reflect.Slice && aVal.Kind() != reflect.Array || aVal.Type() != bVal.Type() {
		panic(fmt.Sprintf("function called with incompatible types: %T, %T", sliceA, sliceB))
	}
	itemType := aVal.Type().Elem()
	if !itemType.Comparable() {
		panic(fmt.Sprintf("elements of %T is not comparable", sliceA))
	}

	aMap := reflect.MakeMap(reflect.MapOf(itemType, reflect.TypeOf(true)))
	bMap := reflect.MakeMap(reflect.MapOf(itemType, reflect.TypeOf(true)))

	for i := 0; i < aVal.Len(); i++ {
		aMap.SetMapIndex(aVal.Index(i), reflect.ValueOf(true))
	}
	for i := 0; i < bVal.Len(); i++ {
		bMap.SetMapIndex(bVal.Index(i), reflect.ValueOf(true))
	}

	aaVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)
	abVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)
	bbVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)

	for i := 0; i < aVal.Len(); i++ {
		v := aVal.Index(i)
		if bMap.MapIndex(v).IsValid() {
			abVal = reflect.Append(abVal, v)
		} else {
			aaVal = reflect.Append(aaVal, v)
		}
	}
	for i := 0; i < bVal.Len(); i++ {
		v := bVal.Index(i)
		if !aMap.MapIndex(v).IsValid() {
			bbVal = reflect.Append(bbVal, v)
		}
	}

	aaPtr := reflect.New(aaVal.Type())
	aaPtr.Elem().Set(aaVal)
	Unique(aaPtr.Interface())

	abPtr := reflect.New(abVal.Type())
	abPtr.Elem().Set(abVal)
	Unique(abPtr.Interface())

	bbPtr := reflect.New(bbVal.Type())
	bbPtr.Elem().Set(bbVal)
	Unique(bbPtr.Interface())

	onlyA = aaPtr.Elem().Interface()
	AandB = abPtr.Elem().Interface()
	onlyB = bbPtr.Elem().Interface()
	return
}
