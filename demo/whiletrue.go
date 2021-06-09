package main

import "fmt"

func main() {

	sum := 0

	index := 1

	for {

		sum += index
		index++
		if index > 100 {
			break
		}

	}
	fmt.Println("1到100相加之和是：", sum)
}
