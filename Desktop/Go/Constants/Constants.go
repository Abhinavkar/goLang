package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 50000
	const d = 330e20 / n
	fmt.Println(n)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
