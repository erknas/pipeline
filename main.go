package main

import "fmt"

// gen converts list of integers into channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

// sq recieve integers from channel and returns channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {
	for num := range sq(gen(1, 2, 3, 4, 5, 6, 7, 8, 9)) {
		fmt.Println(num) // 1 4 9 16 25 36 49 64 81
	}
}
