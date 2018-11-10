package echo1

import (
	"fmt"
	"os"
)

func main1() {
	var s , sep string;
	sep = "\r\n"
	for i := 1;i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}