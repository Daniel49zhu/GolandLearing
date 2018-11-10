package echo2

import (
	"fmt"
	"os"
)

func main3() {
	s,sep := "",""
	for _,arg := range os.Args[1:] {
		s+=sep+arg
		sep = " "
	}
	fmt.Println(os.Args[1])
	fmt.Println(s)
}