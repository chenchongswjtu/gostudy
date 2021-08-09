// 函数选项模式
package main

import "fmt"

type Option interface {
	apply(*options)
}

type options struct {
	name string
	age  int
}

type funcOption func(*options)

func (f funcOption) apply(os *options) {
	f(os)
}

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

func main() {
	var optIns options
	var opts []Option
	opts = append(opts, withAge(1), withName("tom"))
	for _, o := range opts {
		o.apply(&optIns)
	}

	fmt.Println(optIns)
}
