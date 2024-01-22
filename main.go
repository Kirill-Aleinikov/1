package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&b)
	var numbers []int
	for i := 0; i < b; i++ {
		fmt.Scan(&a)
		numbers = append(numbers, a)
	}
	sum(numbers...)
}

func sum(arg ...int) {
	var total = 0
	for _, value := range arg {
		total += value
	}
	fmt.Println(total)
}
