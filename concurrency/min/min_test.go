package min

import "fmt"

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
		104, 33, 25, 12, 4,
		56, 14, 26, 79, 25,
		13, 10, 2, 999, 1024,
	}))
	// Output: 2
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
		104, 33, 25, 12, 4,
		56, 14, 26, 79, 25,
		13, 10, 2, 999, 1024,
	}, 4))
	// Output: 2
}
