package main

import (
	"fmt"
)

func main() {
	var c, p, f int
	fmt.Scan(&c, &p, &f)

	if p >= c*f {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}