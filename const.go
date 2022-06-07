package main

import(
	"fmt"
	"math"
)

const s string = "const"

func main(){
	bob := "red"

	fmt.Println(s)
	fmt.Println(bob)

	const n = 5000000

	const d = 3e20 /n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}