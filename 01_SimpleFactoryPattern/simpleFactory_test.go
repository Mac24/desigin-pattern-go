package main

import "testing"

func Test(t *testing.T) {
	t.Run("Shape Factory:", SimpleFactoryTest)
}

func SimpleFactoryTest(t *testing.T) {
	f := NewFactory()

	circle := f.GetShape("Circle")
	if circle != nil {
		circle.Draw()
	}

	square := f.GetShape("Square")
	if square != nil {
		square.Draw()
	}

	rectangle := f.GetShape("Rectangle")
	if rectangle != nil {
		rectangle.Draw()
	}
}