package main

import "fmt"

func zeroval(ival int) {
	fmt.Println(&ival)
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println(iptr)
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i, " address", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)
	//he &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("pointer", &i)
}

//output:
//0xc00004a098
//zeroval: 1
//0xc00004a080
//zeroptr 0
//pointer 0xc00004a080
