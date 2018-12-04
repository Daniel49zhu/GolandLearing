package main

import "fmt"

//Go supports methods defined on struct types.
type rect struct {
	width, height int
}

//method在接收时可以自动在value和指针之间切换
//传value时会进行复制，因此不会改变原来的值
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}
	fmt.Println("perim:", r.perim())
	fmt.Println("area:", r.area())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
