package main

import (
	"fmt"
	"time"
)

func main() {
	tlog := timing()
	defer tlog()
	fmt.Println("hello edmund.")
}

func timing() func() {
	start := time.Now()
	return func() {
		fmt.Println("time cost", time.Since(start))
	}
}
