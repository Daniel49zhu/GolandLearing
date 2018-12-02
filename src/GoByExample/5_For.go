package main

import "fmt"

func main() {
	// for is Go’s only looping construct.
	i := 1
	//single condition
	for i <= 3 {
		fmt.Print(i, " ")
		i = i + 1
	}
	//A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Print(j, " ")
	}
	//for without a condition will loop repeatedly
	for {
		fmt.Print("loop ")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, " ")
	}
}
