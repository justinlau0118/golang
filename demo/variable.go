package main

import "fmt"

var v0 int
var v1 int = 1
var v2 = 2
var v3, v4 int = 3, 4
var v5, v6 = "5", 6

func main() {
	// var (
	// 	v7 int = 2 + 5
	// 	v8 int = v2 + v6
	// )
	// fmt.Println(v0, v1, v2, v3, v4, v5, v6, v7, v8)
	// var my_name,d_name string = "lt","bn"
	// fmt.Printf("my name is %s\n and %s\n",my_name,d_name)
	var my_name string
	fmt.Print("plz input ur name ")
	fmt.Scan(&my_name)
	fmt.Printf("hello %s",my_name)
}
