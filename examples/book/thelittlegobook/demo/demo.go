package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
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

	errProcess()

	e := deferTest()
	fmt.Println("ee:", e)

}

func deferTest() (err error) {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	//file = nil
	defer func() {
		err = file.Close()
	}()
	return
}

func errProcess() {
	s := "4x"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

	if err := process(2); err != nil {
		fmt.Println("2:", err)
	}
	if err := process(0); err != nil {
		fmt.Println("0", err)
	}
}
func process(count int) error {
	if count < 1 {
		return errors.New("invalid count")
	}
	return nil
}

func interfaceTest() {
	w := &walk{}
	r := &run{}
	moving(w)
	moving(r)
}

func moving(m movable) {
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
