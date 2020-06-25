package main

import (
	"fmt"
	"github.com/EdmundDuntis/edgo/pkg/glib"
)

func main() {
	t := glib.TypeInspect("abc")
	fmt.Println(t)
	t = glib.TypeInspect(1)
	fmt.Println(t)
}
