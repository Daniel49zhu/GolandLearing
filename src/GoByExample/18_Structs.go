package main

import "fmt"

//Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	//Access struct fields with a dot.
	fmt.Println(sp.age) //50

	sp.age = 51
	fmt.Println(s) //51
}
