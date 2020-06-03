package main

import "fmt"

func panicTest(p bool) {
	defer checkPanic()
	if p {
		panic("panic requested")
	}
}

func checkPanic() {
	if r := recover(); r != nil {
		fmt.Println("A Panic wa captured, message :", r)
	}
}

func main() {
	panicTest(true)
	fmt.Println("hello world")
	i := 1
	for true {
		i++
		if i == 10 {
			break
		}
	}
}
