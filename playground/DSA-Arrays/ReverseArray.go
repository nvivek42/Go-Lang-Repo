package main

import "fmt"

func Reversearray(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		u := len(arr) - 1 - i
		arr[i], arr[u] = arr[u], arr[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("before reversing: ", arr)
	Reversearray(arr)
	fmt.Println("after reversing: ", arr)
}
