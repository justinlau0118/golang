package main

import (
	"fmt"
	"testing"
)

func TestNinenine(t *testing.T) {

	sum := 0

	for i := 1; i <= 9; i++ {

		for j := 1; j <= i; j++ {

			sum = j * i
			fmt.Print(j, "x", i, "=", sum, "\t")
		}
		fmt.Println()

	}

}
