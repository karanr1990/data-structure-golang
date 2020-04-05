//binary search program

package main

import "fmt"

func main() {
	arr := []int{6, 8, 31, 54, 67, 71, 84, 95}

	fmt.Printf("%v", arr)
	var searchElement int
	fmt.Scanf("%d", &searchElement)
	fmt.Println("Read number", searchElement, "from stdin")

	var mid int
	start := 0

	arrlength := len(arr) - 1
	end := arrlength

	for start <= end {
		mid = (start + end) / 2

		if searchElement == arr[mid] {
			fmt.Printf("%d found at location %d.\n", searchElement, mid+1)
			break
		} else if searchElement > arr[mid] {
			start = mid + 1
		} else if searchElement < arr[mid] {
			end = mid - 1
		}

	}

	if start > end {
		fmt.Printf("Not found! %d isn't present in the list.\n", searchElement)
	}

}
