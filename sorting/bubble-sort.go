//bubble sort 
package main

import "fmt"

func main() {

	arr := []int{31, 8, 6, 54, 95, 84, 71, 67}

	arraylen := len(arr) - 1

	for i := 0; i < arraylen; i++ {

		for j := 0; j < arraylen; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}

		}

	}

	fmt.Println(arr)

}
