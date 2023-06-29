package main

type I interface {
	Foo()
}

type A struct{}

func (a A) Foo() {}

type B struct{}

func (b B) Foo() {}

func instantiated[X I](x X) {
	x.Foo()
}

var (
	a A
	b B
)

func main() {
	instantiated[A](a) // static call
	instantiated[B](b) // static call

	local[C]().Foo()

	lambda[A]()()()
}

func local[X I]() I {
	var x X
	return x
}

type C struct{}

func (c C) Foo() {}

func lambda[X I]() func() func() {
	return func() func() {
		var x X
		return x.Foo
	}
}
