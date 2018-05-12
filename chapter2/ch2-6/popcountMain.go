package main

import (
	"fmt"
	"os"
	"strconv"

	"./popcount"
)

func main() {
	u, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Printf("v1 pop counted %d = %d\n", u, popcount.PopCount1(u))
	fmt.Printf("v2 pop counted %d = %d\n", u, popcount.PopCount2(u))
	fmt.Printf("v3 pop counted %d = %d\n", u, popcount.PopCount3(u))
	fmt.Printf("v4 pop counted %d = %d\n", u, popcount.PopCount4(u))
}
