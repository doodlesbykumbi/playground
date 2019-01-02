package main

import (
	"fmt"
	"time"
	"example/pkg"
)

func main() {
	for {
		fmt.Println("this is fun!", pkg.SOME_CONST)

		time.Sleep(time.Second * 1)
	}
}
