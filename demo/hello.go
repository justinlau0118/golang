package main

import "fmt"

/* func main() {
	中文 := "hello，你好，中文"
	fmt.Println(中文, &中文)
} */

func main() {
	var arr1 [2]string
	arr1 = [2]string{"1", "2"}
	fmt.Println(arr1)

	var slice1 []string
	slice1 = []string{"a", "b", "c"}
	slice1 = append(slice1, "x")
	fmt.Println(slice1)

	var slice2 []string
	slice2 = make([]string, 0, 3)
	slice2 = append(slice2, "x", "y", "z", "m")
	fmt.Println(slice2)

	var slice4 []int
	slice4 = make([]int, 0, 3)
	slice4 = append(slice4, 1, 2, 3)
	fmt.Println(slice4, len(slice4))

	var slice5 []int
	slice5 = []int{1, 2, 3, 4}
	slice6 := slice5[0:2]
	b, f := "a", 1
	fmt.Println(slice6, b, f)

	var slice7 []int
	copy(slice7, slice6)
	fmt.Println(slice7)
	var m1 map[int]string
	fmt.Println(m1)
	var m2 = map[int]string{01: "xiaogao"}
	m2[02] = "boniuniu"
	m2[02] = "piniuniu"
	// delete(m2, 01)
	fmt.Println(len(m2))

	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for i, v := range slice5 {
		fmt.Println(i, v)
	}
}
