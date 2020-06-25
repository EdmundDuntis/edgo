package main

import (
	"fmt"
	"github.com/EdmundDuntis/edgo/pkg/glib"
)

var phrase func(phraseName string)=func(phraseName string){
	fmt.Println("================================",phraseName,"================================")
}

func main() {
	phrase("arrayOrSlice")
	arrayOrSlice()
	phrase("rangeString")
	rangeString()
	phrase("mapKeyNotExist")
	mapKeyNotExist()
	phrase("variableParamter")
	variableParamter(1,2,3,4,5)

}

// variableParamter 变参函数
func variableParamter(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func mapKeyNotExist() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	// 存在k2
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// 删除之后不存在k2
	delete(m, "k2")
	_, prs = m["k2"]
	fmt.Println("prs:", prs)
}

func rangeString() {
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
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
