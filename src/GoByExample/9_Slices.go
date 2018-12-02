package main

import "fmt"

func main() {
	// Slices are a key data type in Go,
	// giving a more powerful interface to sequences than arrays.
	//通过make创建一个slice
	s := make([]string, 3)
	fmt.Println("emp:", s) //outout: emp: [  ]
	//We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) //output: set: [a b c]
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//内置的append方法可以动态增长slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) //output: apd: [a b c d e f]
	//Slices can also be copy’d.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)
	//Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	//We can declare and initialize a variable for slice
	// in a single line as well.
	//不同于Array，[]中不指定长度
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//Slices can be composed into multi-dimensional data structures.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
