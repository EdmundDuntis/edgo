package main

import "fmt"

func demo(recoverHighestPanicAtFirst bool) {
	fmt.Println("====================")
	defer func() {
		if !recoverHighestPanicAtFirst{
			// 恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer func() {
			//  恢复恐慌2
			fmt.Println("恐慌", recover(), "被恢复了")
		}()
		if recoverHighestPanicAtFirst {
			//  恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer fmt.Println("现在有两个恐慌共存")
		panic(2)
	}()
	panic(1)
}

func main() {
	demo(true)
	demo(false)
}