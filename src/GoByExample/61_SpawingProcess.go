package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("ipconfig")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">ipconfig")
	fmt.Println(string(dateOut))
}
