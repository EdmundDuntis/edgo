package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
	"time"
)

func main() {

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	deferExample()
	b()
	_, _ = func1("Go")

	t := time.Now()
	for i:=0;i<9000;i++{
		fibonacci(i)
	}

	where()

	d := time.Since(t)
	fmt.Println(d)

	t1 := time.Now()
	for i:=0;i<9000;i++{
		fibonacci1(i)
	}
	d1 := time.Since(t1)
	fmt.Println(d1)

	fmt.Println(strings.IndexFunc("中abc文def", IsAscii))

	addBmp, addJpeg := MakeAddSuffix(".bmp"), MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file"), addJpeg("file"))

}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
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

var fibres [10000]int

func fibonacci(n int) (res int) {
	if fibres[n] != 0 {
		res = fibres[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibres[n] = res
	return
}

func fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func IsAscii(c rune) bool {
	if c > 255 {
		return false
	}
	return true
}



