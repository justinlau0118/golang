// package main

// import (
// 	"fmt"
// )

// func main() {

// 	sum := 0

// 	for i := 1; i <= 9; i++ {

// 		for j := 1; j <= i; j++ {

// 			sum = j * i
// 			fmt.Print(j, "x", i, "=", sum, "\t")
// 		}
// 		fmt.Println()

// 	}

// }
package main

import (
	"fmt"
)

func main() {

for i := 1; i <=9;i++ {

   for j := 1; j<=i;j++{

      fmt.Print(j,"x",i,"=",i*j,"\t")

}
fmt.Println()
}
}