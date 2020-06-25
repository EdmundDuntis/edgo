package main

import "fmt"

func main() {
	messages := make(chan int)
	go func() {
		for n := 0; n < 10; n++ {
			messages <- n
		}
		close(messages)
	}()

/*	for {
		msg, more := <-messages
		if more {
			fmt.Println(msg)
		} else {
			break
		}
	}*/

	for msg := range messages {
		fmt.Println(msg)
	}
}
