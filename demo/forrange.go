package main

import (
	"fmt"
)

func main() {

	g := "波妞妞"

	for i, v := range g {
		fmt.Printf("%T", g, i, v)

	}

}
