// 函数选项模式
package main

import "fmt"

type options struct {
	name string
	age  int
}

type funcOption func(*options)

func withName(name string) funcOption {
	return func(o *options) {
		o.name = name
	}
}

func withAge(age int) funcOption {
	return func(o *options) {
		o.age = age
	}
}

type option interface {
	apply(o *options)
}

func (f funcOption) apply(o *options) {
	f(o)
}

func main() {
	var optionInst options
	var options []option
	options = append(options, withAge(1), withName("tom"))
	for _, o := range options {
		o.apply(&optionInst)
	}

	fmt.Println(optionInst)
}
