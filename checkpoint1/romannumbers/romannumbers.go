package main

import (
	"fmt"
	"os"
)

func main() {
	const r1, r5, r10, r50, r100, r500, r1000 = 1, 5, 10, 50, 100, 500, 1000
	num := os.Args[1]
	fmt.Printf("%d + %d = %d", r5, r5, r10)
	fmt.Printf(num)

}
