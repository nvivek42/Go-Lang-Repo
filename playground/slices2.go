package main

import "fmt"

func main() {
	arr := [4]string{"vivek", "narole", "kunal", "bura"}
	a := arr[0:3]
	fmt.Println("a: ", a, "arr: ", arr)
	a[0] = "kunal"
	a[2] = "vivek"
	fmt.Println("a: ", a, "arr: ", arr)
}
