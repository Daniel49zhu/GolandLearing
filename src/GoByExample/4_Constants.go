package main

import (
	"fmt"
	"math"
)

//A const statement can appear anywhere a var statement can.
const s string = "constant"

func main() {
	fmt.Println(s)

	//const可以是任意类型任意精度

	const n = 500000000000000
	const d = 3e20 / n
	fmt.Println(d)
	//const的数字没有具体类型除非他被显式转换
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
