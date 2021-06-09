package main

import "fmt"

func main() {

	for {
	score := 0.0
	fmt.Println("plz input score：")
	fmt.Scan(&score)
	switch {
	case score <= 100 && score >= 90:
		fmt.Println("优秀", score)
	case score < 90 && score >= 80:
		fmt.Println("良好", score)
	case score < 80 && score >= 70:
		fmt.Println("及格", score)
	default:
		fmt.Println("加油", score)
}
    if score < 70{
		break
	}
	}
}
