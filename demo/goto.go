package main

import "fmt"

func main() {

	sum := 0
	index := 1

START:

	if index > 100 {
		goto END
	}
	sum += index
	index++
	goto START

END:
	fmt.Println(sum)

}
