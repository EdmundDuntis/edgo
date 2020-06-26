package glib

import "fmt"

// TypeInspect 类型查看
func TypeInspect(i interface{}) string {
	rst := ""
	switch t := i.(type) {
	default:
		rst = fmt.Sprintf("%T", t)
	}
	return rst
}


