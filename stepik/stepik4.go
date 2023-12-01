// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var a, b int

// 	fmt.Scan(&a, &b)

// 	res := 0

// 	for i := a; i <= b; i++ {
// 		res += i

// 	}
// 	fmt.Println(res)
// }

package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 1.0; i <= 10.0; i++ {
		res := math.Pow(i, 2)
		fmt.Println(res)
	}

}
