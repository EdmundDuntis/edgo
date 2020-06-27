package main

import "fmt"

type Saiyan struct {
	Name    string
}

func main() {
	a := make([]Saiyan, 10)
	//或者
	//b := make([]*Saiyan, 10)

	a0:=Saiyan{"a0"}
	a1:=Saiyan{"a1"}
	a[0]=a0
	a[1]=a1
	fmt.Println(a[0],a[1])
	fmt.Println(a0,a1)
	a[0]=a0
	a[1]=a1
	a0.Name="x"
	fmt.Println(a[0],a[1])
	fmt.Println(a0,a1)
}