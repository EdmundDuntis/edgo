package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

func main() {
	deferExample()
	b()
	_, _ = func1("Go")

	//t:=time.Now()
	//res:=fibonacci(50)
	//d:=time.Since(t)
	//fmt.Println(res,d)
	t1:=time.Now()
	res:=fibonacci1(40)
	d1:=time.Since(t1)
	fmt.Println(res,d1)


	fmt.Println(strings.IndexFunc("中abc文def",IsAscii))

}

func deferExample() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

// trace为支持链式调用需要返回入参形式相同的返回值
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	// 链式调用
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	// 链式调用
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

// 使用 defer 语句来记录函数的参数与返回值
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}


var fibres [100]int
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
		fibres[n] = 1
	} else {
		fiba, fibb := fibres[n-1], fibres[n-2]
		if fiba == 0 {
			fiba = fibonacci(n - 1)
			fibres[n-1] = fiba
		}
		if fibb == 0 {
			fibb = fibonacci(n - 2)
			fibres[n-2] = fibb
		}
		res = fiba + fibb
	}
	return
}


func fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}

func IsAscii(c rune) bool {
	if c > 255 {
		return false
	}
	return true
}
