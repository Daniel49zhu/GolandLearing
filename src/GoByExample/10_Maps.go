package main

import "fmt"

// Maps are Go’s built-in associative data type
// (sometimes called hashes or dicts in other languages).
func main() {
	//To create an empty map,
	// use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)
	//Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)
	//Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))
	//The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)
	//获取key值返回的第二参数代表key是否存在
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	//You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
