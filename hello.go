package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				j := i * i
				fmt.Println(i, j)
				time.Sleep(time.Second)
				runtime.Gosched()
			}
		}(i)
	}

	select {}

}
