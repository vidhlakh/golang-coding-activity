package smallestcommonnumber

import (
	"fmt"
	"sort"
)

func smallestcommonnumber(a, b, c []int) int {
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("C:", c)
	i := 0
	k := 0
	j := 0
	N := len(a)
	M := len(b)
	P := len(c)
	for i < M && j < N && k < P {
		fmt.Println("a,b,c", a[i], b[j], c[k])
		if a[i] == b[j] && a[i] == c[k] {
			return a[i]
		} else if a[i] > b[j] {
			j = j + 1
		} else if a[i] > c[k] {
			k = k + 1
		} else {
			i = i + 1
		}

	}
	return -1
}
func SmallestCommonfunc() {
	a := []int{1, 4, 12, 34, 19, 28, 34, 9}
	b := []int{23, 45, 6, 7, 8, 9, 34}
	c := []int{3, 5, 7, 9, 34, 19}
	result := smallestcommonnumber(a, b, c)
	fmt.Println("Result:", result)
}

// i = 0,j = 0, k = 0
// 1 ==6 == 3 ?
