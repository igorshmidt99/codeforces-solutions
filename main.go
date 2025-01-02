package main

import (
	"fmt"
	"slices"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	slices.SortFunc[[]int, int](sl, func(a, b int) int { return b - a })
	fmt.Printf("%v\n", sl)
}
