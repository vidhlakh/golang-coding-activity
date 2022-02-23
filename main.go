package main

import (
	"fmt"

	"github.com/vidhlakh/golang-coding-activity/binaryOperators"
	"github.com/vidhlakh/golang-coding-activity/countbits"
	"github.com/vidhlakh/golang-coding-activity/kthlargest"
	"github.com/vidhlakh/golang-coding-activity/reversewords"
)

func main() {
	fmt.Println("HEllo")
	//result := Countbits(6)
	res2 := countbits.CountbitsMethod2(100)
	fmt.Println("Number of set bits: ", res2)
	binaryOperators.BinaryOperators()
	reversewords.Reversefunc()
	kthlargest.Kthlargestfunc()
	//smallestcommonnumber.SmallestCommonfunc()
}
