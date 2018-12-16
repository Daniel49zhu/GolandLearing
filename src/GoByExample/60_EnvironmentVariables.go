package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Getenv("path"))
	fmt.Println(os.Getenv("JAVA_HOME"))
	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
