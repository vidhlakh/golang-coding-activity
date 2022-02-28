package sorting

import "fmt"

func Mergesort() {
	items := []int{9, 4, 3, 6, 1, 2, 10}
	//ditemstitems2 := []int{2, 10, 5, 7, 8}
	res := mergeSortfunc(items)
	fmt.Println(res)

}
func mergeSortfunc(items []int) []int {
	if len(items) < 2 {
		return items
	}

	left := mergeSortfunc(items[:len(items)/2])
	right := mergeSortfunc(items[len(items)/2:])
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	return merge(left, right)
}

// merge itemsctuitemsl code
func merge(left []int, right []int) []int {
	i := 0
	j := 0
	final := []int{}
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}

	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	fmt.Println("final so far:", final)
	return final
}
