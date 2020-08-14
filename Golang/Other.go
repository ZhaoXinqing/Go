package main

import (
	"fmt"
)

type Person interface {
	SayHello()
}

type Girl struct {
	Sex string
}

type Boy struct {
	Sex string
}

func (this *Girl) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}

func (this *Boy) SayHello() {
	fmt.Println("Hi, I am a " + this.Sex)
}

func main() {
	g := &Girl{"girl"}
	b := &Boy{"boy"}

	p := map[int]Person{}
	p[0] = g
	p[1] = b

	for _, v := range p {
		v.SayHello()
	}
}
