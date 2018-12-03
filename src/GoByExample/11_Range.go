package main

import "fmt"

func main() {
	// range iterates over elements in a variety of data structures.

	//range on arrays and slices
	nums := []int{2, 3, 4}
	sum := 0
	//可以用_来省略第一个参数index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k, _ := range kvs {
		fmt.Println("key:", k)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, byte(c))
	}

}
