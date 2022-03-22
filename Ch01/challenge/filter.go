package main

import (
	"fmt"
)

// filter returns a slice with only the values that pred(val) returned true
func filter(pred func(int) bool, values []int) []int {

	//create a slice of length 0 and capacity same as original slice
	var filteredValues []int;
	filteredValues = make([]int, 0, len(values));

	// start by iterating over the values in slice "values"
	for _, v := range	values {
		// apply filter
		if pred(v) == true {
			// if filter matches, add to the returned values
			filteredValues = append(filteredValues, v);
		}
	}

	return filteredValues; // FIXME
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(filter(isOdd, values)) // [1 3 5 7]
}
