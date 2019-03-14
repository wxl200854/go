package main

import (
	"fmt"
)

type Human interface {
	sayhi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	name  string
	group string
}

type Mystr string

func (s *Student) sayhi() {
	fmt.Println("Student: s = ", *s)
}

func (t *Teacher) sayhi() {
	fmt.Println("Teacher: t = ", *t)
}

func (s *Mystr) sayhi() {
	fmt.Println("Mystr: s = ", *s)
}

func main() {
	var i Human
	s := &Student{"mike", 15}
	i = s
	i.sayhi()

	t := &Teacher{"lili", "nihao"}
	i = t
	i.sayhi()

	// var s0 Mystr = "welcome"
	s0 := Mystr("welcome")
	i = &s0
	i.sayhi()
}
