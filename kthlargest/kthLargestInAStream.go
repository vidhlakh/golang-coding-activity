package kthlargest

import "fmt"

func Kthlargestfunc() {
	input := []int{10, 20, 11, 70, 50, 40, 100, 5}
	kthlargestSimple(input)
}
func kthlargestSimple(input []int) (output int) {
	m := &MaxHeap{}
	for _, ele := range input {
		m.Insert(ele)
		fmt.Println("HEap", m)
	}
	fmt.Println("Max heap:", m)
	for i := 0; i <= 5; i++ {
		m.Extract()
		fmt.Println("Extract:", m)
	}
	return output
}
