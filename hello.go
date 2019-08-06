package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer func() {
		log.Println("returned")
	}()
	fmt.Println("hello edmund")
	time.Sleep(time.Second)
	// fmt.Println(os.Environ())
}
