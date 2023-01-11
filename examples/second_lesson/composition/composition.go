package main

import "fmt"

func main() {
	c1 := Concrete1{}
	c2 := Concrete2{}
	c1.Name = "c1"
	//c2.Base.Name = "c2"
	c1.SayHello()
	c2.SayHello()
}

// Swimming 会游泳的
type Swimming interface {
	Swim()
}

type Duck interface {
	// 鸭子是会游泳的，所以这里组合了它
	Swimming
}

type Base struct {
	Name string
}

type Concrete1 struct {
	Base
}

type Concrete2 struct {
	*Base
}

func (c Concrete1) SayHello() {
	// c.Name 直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)

	// 调用了被组合的
	c.Base.SayHello()
}

func (b *Base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}
