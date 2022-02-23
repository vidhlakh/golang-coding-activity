package binaryOperators

import "fmt"

// << o fill left shift
// >> signed right shift
// 000 0
// 001 1
// 010 2
// 011 3
// 100 4
// 101 5
// 110 6
// 111 7
// 1000 8
// 011 << 1 -> 110 6
// 011 >> 1 -> 001 1
// 1011 << 1 -> 0110
// 1101 << 1 -> 1010
func BinaryOperators() {
	var x uint = 3 // 011  //0000 1001
	var y uint = 7 // 111 //0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)

	// Bit wise Left shift multiplication
	fmt.Println("6 << 3", -6<<3)
	// Bitwise Right shift division
	fmt.Println("48 >> 3", -48>>3)
}
