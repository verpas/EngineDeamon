package gotest

import "fmt"

func Print(i int) (ret bool) {
	fmt.Print("test")
	ret = false
	if i == 0 {
		ret = true
	}
	return
}
