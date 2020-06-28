package main

import "fmt"
import "github.com/mattn/go-sqlite3"

type Saiyan struct {
	Name string
}

type movable interface {
	move()
}

type walk struct{}

func (w *walk) move() {
	fmt.Println("walk")
}

type run struct{}

func (w *run) move() {
	fmt.Println("run")
}

func main() {
	_, _ = packages()
	pointerOrValue()
	interfaceTest()
}

func interfaceTest() {
	w := &walk{}
	r := &run{}
	moving(w)
	moving(r)
}

func moving(m movable)  {
	m.move()
}

func packages() (int, error) {
	return fmt.Println(sqlite3.Version())
}

func pointerOrValue() {
	a := make([]*Saiyan, 10)
	//或者
	//b := make([]*Saiyan, 10)
	a0 := &Saiyan{"a0"}
	a1 := &Saiyan{"a1"}
	a[0] = a0
	a[1] = a1
	fmt.Println(a[0], a[1])
	fmt.Println(a0, a1)
	a[0] = a0
	a[1] = a1
	a0.Name = "x"
	fmt.Println(a[0], a[1])
	fmt.Println(a0, a1)
}
