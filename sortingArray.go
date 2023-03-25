package sortingArray

import "fmt"

// major upgrade adalah perubahan yg mejebabkan error atau perubahan dratis pada user yg menggunakan package, sebuah contoh pada parameter
func InsectionSort(jmlh_data int, numbers ...int) {
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println("=======proccess", jmlh_data, "=======")
	fmt.Println(numbers)
	fmt.Println("=======done=======")
}
