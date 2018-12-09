package main

import "os"

func main() {
	//panic用于快速失败
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
