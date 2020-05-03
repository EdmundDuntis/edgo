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
	fmt.Println("hello edmund duntis")
	time.Sleep(time.Second)
}
