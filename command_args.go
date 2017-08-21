package main

import (
	"fmt"
	"os"
)

func main() {
	var s,seq string
	for i:=1;i<len(os.Args);i++ {
		s += seq + os.Args[i]
		seq = " "
	}

	var s2,seq2 string

	for _,arg := range os.Args[0:] {
		//中间用空格 隔开
		s2 += seq2 + arg
		seq2 = " "
	}



	fmt.Println(s,s2)
}


