package main

import (
	"fmt"
)

type sds []byte

func structConvInterface(s *sds) interface{} {
	var inter interface{}
	inter = s
	return inter
}

func main() {
	//1000 0001
	fmt.Println(1024 * 1024 * 512)
	fmt.Println(536870911 > 1024*1024*512)
}
