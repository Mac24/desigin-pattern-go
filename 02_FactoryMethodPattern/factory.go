package main

import "fmt"

//动物类
type animal struct {}

//具体的动物
type dog struct {
	animal
}

type cat struct {
	animal
}

type pig struct {
	animal
}

//动物都有speak()方法(抽象产品)
type Speaker interface {
	Speak()
}

func (d *dog) Speak() {
	fmt.Println("我是dog, 汪汪汪")
}

func (c *cat) Speak() {
	fmt.Println("我是cat, 喵喵喵")
}

func (p *pig) Speak() {
	fmt.Println("我是pig, 呼呼呼")
}

//抽象工厂接口
type AnimalFactory interface {
	GetAnimal() Speaker
}

//生产具体动物的工厂
type DogFactory struct {}
type CatFactory struct {}
type PigFactory struct {}

func (d *DogFactory) GetAnimal() Speaker {
	return &dog{}
}

func (c *CatFactory) GetAnimal() Speaker {
	return &cat{}
}

func (p *PigFactory) GetAnimal() Speaker {
	return &pig{}
}


