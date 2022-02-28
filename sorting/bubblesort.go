package sorting

import "fmt"

// []int { }

func Bubblesort() {
	data := []int{9, 4, 3, 6, 1, 2, 10}
	data2 := []int{2, 10, 5, 7, 8}
	chan1 := make(chan []int, 0)
	chan2 := make(chan []int, 0)
	//chan3 := make(chan []int, 0)
	go sort(data, chan1)
	res := <-chan1
	fmt.Print(res)
	go sort(data2, chan2)
	res = <-chan2
	fmt.Print(res)

	//result = sort(data2)
	//fmt.Print(result)
}
func sort(data []int, chan1 chan []int) {
	//var result []int
	//
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
				//fmt.Println("Data iteration", data)
			}
		}
	}
	//fmt.Println("data", data)
	chan1 <- data
	//return data
}
