package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 0
	opp := 1
	rand.Seed(time.Now().Unix()) //设置随机数种子(只需要执行一次)
	randnum := (rand.Intn(100))
	for opp <= 5 {

		fmt.Println("plz inpu number,five chances: ", randnum)
		fmt.Scan(&num)

		opp++

		if num > randnum {

			fmt.Println("wrong,gt")
		} else if num < randnum {

			fmt.Println("wrong,lt")

		} else {
			fmt.Println("bingo")
			return
		}

	}
	fmt.Println("already no chance,bad luck")
}
