package main

import "testing"

func Test(t *testing.T){
	t.Run("AnimalFactory test", FactoryTest)
}

func FactoryTest(t *testing.T) {
	d := DogFactory{}
	d.GetAnimal().Speak()

	c := CatFactory{}
	c.GetAnimal().Speak()

	p := PigFactory{}
	p.GetAnimal().Speak()
}