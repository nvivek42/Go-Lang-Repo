package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(msg, ": ", i)
	}
}
func main() {
	go f("direct")
	go f("goRoutines")
	go func(msg string) {
		for i := 0; i < 15; i++ {
			fmt.Println(msg)
		}
	}("async func")
	time.Sleep(time.Second)

	fmt.Println("after goroutines call")

}
