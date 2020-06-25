package main

import (
	"fmt"
	"github.com/EdmundDuntis/edgo/pkg/glib"
)

func main() {
	arrayOrSlice()
}

func arrayOrSlice() {
	// 数组，类型推断
	a := [...]string{"a", "b", "c"}
	fmt.Println(glib.TypeInspect(a))
	// 切片
	s1 := []string{"a", "b", "c"}
	fmt.Println(glib.TypeInspect(s1))
	s2 := make([]string, 3)
	fmt.Println(glib.TypeInspect(s2))
	fmt.Println(cap(s2), len(s2))
	// 切片，长度和容量
	s3 := make([]string, 3, 6)
	fmt.Println(glib.TypeInspect(s3))
	fmt.Println(cap(s3), len(s3))
}
