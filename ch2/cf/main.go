package main

import (
	"fmt"
	"os"
	"strconv"

)

func main() {

	for _ , arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg,64)
		if err != nil{
			fmt.Fprintf(os.Stderr,"cf:%v\n",err)
		}

		f :=
	}

}
