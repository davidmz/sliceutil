package sliceutil

// Disposer is a function that performs cleanup uperations over the
// removed slice items.
//
// Some sliceutil functions (they are takes slice
// pointer instead of slice itself) can remove items from the slice.
// These functions always fill the freed slice slots by the zero values
// to prevent the memory leaks, so normally you don't have to worry about it.
// But sometimes it may be necessary to perform some additional cleanup
// procedures and this is what disposers are for.
//
// The disposer takes an integer index in the slice of the item being removed.
// You can obtain slice item as slice[index] in disposer.
// Note that it is not the same index that item was have before the function run!
// This index and the corresponding slice item are only exists during the
// disposer call.
type Disposer func(index int)
