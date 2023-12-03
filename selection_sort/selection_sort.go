package main

import "fmt"

func findSmallest(arr []int) int {
	smallest_value := arr[0]
	smallest_index := 0
	for index, value := range arr {
		if value < smallest_value {
			smallest_index = index
		}
	}
	return smallest_index
}

func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for index := range arr {
		smallest := findSmallest(arr)
		newArr[index] = arr[smallest]
		fmt.Println("arr[:smallest]: ", arr[:smallest])
		fmt.Println("arr[smallest+1:]: ", arr[smallest+1:])

		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func main() {
	s := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(s))
}
