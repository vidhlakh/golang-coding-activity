package countbits

import (
	"fmt"
	"strconv"
)

// Input : n = 6
// Output : 2
// Binary representation of 6 is 110 and has 2 set bits

// Input : n = 13
// Output : 3
// Binary representation of 13 is 1101 and has 3 set bits
func Countbits(input int64) int64 {
	if input < 0 {
		return 0
	}
	dict := make(map[string]int64)
	binaryInput := strconv.FormatInt(input, 2)
	for _, s := range binaryInput {
		dict[string(s)] += 1
	}
	fmt.Println("Dict :", dict["1"])
	return dict["1"]
}

// Brian Kernighan Algorithm
func CountbitsMethod2(n int64) int64 {
	var count int64
	fmt.Println("Binary rep :", strconv.FormatInt(n, 2))
	for n > 0 {
		n &= (n - 1)
		fmt.Println("n:", n)
		count += 1
	}
	fmt.Println("Count of set bits:", count)
	return count
}
