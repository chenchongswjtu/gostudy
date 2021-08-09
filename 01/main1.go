package main

import "fmt"

type Optioner interface {
	apply(*option)
}

type option struct {
	name string
	age  int
}

type funcOption struct {
	f func(*option)
}

func (fo *funcOption) apply(o *option) {
	fo.f(o)
}

func newFuncOption(f func(*option)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithName(name string) Optioner {
	return newFuncOption(func(o *option) {
		o.name = name
	})
}

func WithAge(age int) Optioner {
	return newFuncOption(func(o *option) {
		o.age = age
	})
}

func main() {
	var oi option
	var ops []Optioner
	ops = append(ops, WithName("tom"), WithAge(1))

	for _, o := range ops {
		o.apply(&oi)
	}

	fmt.Println(oi)
}
