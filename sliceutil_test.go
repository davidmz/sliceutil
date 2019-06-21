package sliceutil_test

import (
	"fmt"
	"math/rand"

	"github.com/davidmz/sliceutil"
)

func ExampleFilter() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	sliceutil.Filter(&numbers, func(i int) bool { return numbers[i]%2 == 0 })
	fmt.Println(numbers)
	// Output: [4 8 16 42]
}

func ExampleFilter_nil() {
	numbers := []int(nil)
	sliceutil.Filter(&numbers, func(i int) bool { return numbers[i]%2 == 0 })
	fmt.Println(numbers == nil)
	// Output: true
}

func ExampleFilter_disposer() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	sliceutil.Filter(
		&numbers,
		func(i int) bool { return numbers[i]%2 == 0 },
		func(i int) { fmt.Println("Disposed:", numbers[i]) },
	)
	fmt.Println(numbers)
	// Output:
	// Disposed: 23
	// Disposed: 15
	// [4 8 16 42]
}

func ExampleShuffle() {
	rand.Seed(42)
	numbers := []int{4, 8, 15, 16, 23, 42}
	sliceutil.Shuffle(numbers)
	fmt.Println(numbers)
	// Output: [4 42 16 23 15 8]
}

func ExampleReverse() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	sliceutil.Reverse(numbers)
	fmt.Println(numbers)
	// Output: [42 23 16 15 8 4]
}

func ExampleSome() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	found := sliceutil.Some(numbers, func(i int) bool { return numbers[i] == 15 })
	fmt.Println(found)
	// Output: true
}

func ExampleSome_notFound() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	found := sliceutil.Some(numbers, func(i int) bool { return numbers[i] == 17 })
	fmt.Println(found)
	// Output: false
}

func ExampleDiff3() {
	a := []int{4, 4, 8, 15, 16, 23, 42}
	b := []int{5, 9, 15, 16, 23, 48, 48}
	aa, ab, bb := sliceutil.Diff3(a, b)
	fmt.Println(aa, ab, bb)
	// Output: [4 8 42] [15 16 23] [5 9 48]
}

func ExampleUnique() {
	numbers := []int{4, 8, 15, 8, 15, 16, 23, 42, 23}
	sliceutil.Unique(&numbers)
	fmt.Println(numbers)
	// Output: [4 8 15 16 23 42]

}
