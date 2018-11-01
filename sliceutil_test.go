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
