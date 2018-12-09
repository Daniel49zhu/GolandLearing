package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("D://demo.txt")
	//在此处定义的closeFile方法，会在main方法结束之后执行
	// This will be executed at the end of the enclosing function (main), after writeFile has finished.
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "Hello World\r\n --from 42_Defer.go")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
