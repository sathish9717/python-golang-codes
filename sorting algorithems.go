// insertion sort

package main

import "fmt"

func main() {
	l := []int{6,3,8,9,5,10}
	InsertionSort(l)
	fmt.Println(l)

}

func InsertionSort(arr []int) {
	
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
// bubblesort
package main

import "fmt"

func main() {
	list := []int{3, 7, 5, 9, 5, 2, 7}
	bblSort(list)
	fmt.Println(list)

}

func bblSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var a bool
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				a = true
			}
		}
		if !a {
			break
		}
	}

}
//mergesort
package main

import "fmt"

func main() {
	l := []int{3, 7, 5, 9, 5, 2, 7}
	MergeSort(l)
	fmt.Println(l)

}
func MergeSort(arr []int) (arr2 []int) {
	if len(arr) < 2 {
		return arr
	}
	a := MergeSort(arr[:len(arr)/2])
	b := MergeSort(arr[len(arr)/2:])
	
	s := Merge(a, b)
	for i, v := range s {
		arr[i] = v
	}
	return arr

}
func Merge(a []int, b []int) []int {
	L := 0
	R := 0
	arr := []int{}
	for L < len(a) && R < len(b) {
		if a[L] < b[R] {
			arr = append(arr, a[L])
			L++
		} else {
			arr = append(arr, b[R])
			R++
		}
	}
	if L < len(arr) {
		arr = append(arr, a[L:]...)
	}
	if R < len(arr) {
		arr = append(arr, b[R:]...)
	}
	return arr
  

}
//selectionsort
package main

import "fmt"

func main() {
	list := []int{3, 7, 5, 9, 5, 2, 7}
	SelectionSort(list)
	fmt.Println(list)

}
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}


