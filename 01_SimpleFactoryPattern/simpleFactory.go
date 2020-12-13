package main

import "fmt"

type Shape interface {
	Draw()
}

//Circle 圆形类
type Circle struct {}

//Square 正方形类
type Square struct {}

//Rectangle 矩形类
type Rectangle struct {}

//实例化Circle类
func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Draw() {
	fmt.Println("Circle Draw() method.")
}

//实例化Square类
func NewSquare() *Square {
	return &Square{}
}

func (s *Square) Draw() {
	fmt.Println("Square Draw() method.")
}

//实例化Rectangle类
func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle Draw() methmod.")
}

//创建工厂类
type Factory struct {}

//实例化工厂类
func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetShape(shapetype string) Shape{
	switch shapetype {
	case "Circle":
		return NewCircle()
	case "Square":
		return NewSquare()
	case "Rectangle":
		return NewRectangle()
	default:
		return nil
	}
}