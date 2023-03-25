package sortingArray

import "fmt"

func InsectionSort(numbers ...int) {
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println(numbers)

}
