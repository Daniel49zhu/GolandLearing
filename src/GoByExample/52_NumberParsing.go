package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	//the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	i1, _ := strconv.ParseInt("1001", 2, 64)
	i2, _ := strconv.ParseInt("1001", 10, 64)
	i3, _ := strconv.ParseInt("1001", 16, 64)
	fmt.Println(i1) //9
	fmt.Println(i2) //1001
	fmt.Println(i3) //4097

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
