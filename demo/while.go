package main

import  "fmt"

func main(){

sum := 0

index := 1

for  index <= 100 {

	sum += index
	index ++

}
	fmt.Println("1到100相加之和是：",sum)
}