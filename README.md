# sliceutil

Package sliceutil is a set of tiny but useful functions designed for in-place slice operations.

See full documentation on [GoDoc](http://godoc.org/github.com/davidmz/sliceutil)

License: Unlicensed

## Package functions

### In-place operations
```go
// Filter filters Go slice in-place.
// If slice is not a slice pointer, Filter panics.
func Filter(slicePtr interface{}, predicate func(int) bool)
```

```go
// Reverse reverts Go slice in-place.
// If slice is not a slice, Reverse panics.
func Reverse(slice interface{})
```

```go
// Shuffle shuffles Go slice in-place.
// If slice is not a slice, Shuffle panics.
func Shuffle(slice interface{})
```

```go
// Unique leaves only unique items in slice. The slice is modified in-place.
// If slice is not a slice pointer, Unique panics.
func Unique(slicePtr interface{})
```

### Operations that keeps arguments untouched

```go
// Some finds the first item in slice which satisfies the predicate
// and returns true if it was found or false otherwise.
// If slice is not a slice or an array, Some panics.
func Some(slice interface{}, predicate func(int) bool) bool
```

```go
// Diff3 calculates difference between two slices and return three
// new slices: onlyA with items belongs only to A slice, AandB with items common to A and B
// and onlyB with items belongs only to B. A and B must be slices of the same type and
// their items must be comparable.
func Diff3(sliceA, sliceB interface{}) (onlyA, AandB, onlyB interface{})
```
