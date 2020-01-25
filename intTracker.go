package main

import (

	"fmt"

	"github.com/dyk0/intTracker/stringparse"
)

func main() {
	if stringparse.Parse("vgo a") {
		fmt.Println("passed")
	}else{
		fmt.Println("failed")
	}
}
