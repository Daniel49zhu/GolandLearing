package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("src/GoByExample/testRead.txt")
	check(err)
	fmt.Print(string(dat))
}
