package missingnumber

import (
	"fmt"
)

func MissingNumber() {
	input := []int{1, 2, 3, 4, 6, 7, 8, 9}
	res := missingMethod2(input)
	fmt.Println("Missing number is :", res)
}

func missing(input []int) (res int) {
	i := input[0]
	for j := 0; j < len(input); j++ {
		if i == input[j] {
			i++
		} else {
			res = i
		}
	}
	return res
}

// Method is to find sum of n natural numbers , sum of input array, find the differnece to get the element
// sum of n natural nu : n(n+1)/2
func missingMethod2(input []int) (res int) {
	n := input[len(input)-1]
	actualSum := n * (n + 1) / 2
	fmt.Println("Auctual sum", actualSum)
	sum := 0
	for _, i := range input {
		sum += i
	}
	fmt.Println("Got sum:", sum)
	res = actualSum - sum
	return res
}
