package main

import (
	"fmt"
	"testing"
)

func Test (t *testing.T) {
	t.Run("AbstractFactory test", AbstractFactoryTest)
}

func AbstractFactoryTest(t *testing.T) {
	f := NewFactoryProducer()
	h := f.GetFactory("HairFactory")

	ha := h.createAir("HairAir")
	if ha != nil {
		ha.use()
	} else {
		fmt.Println("Not HairFactory")
	}

	ht := h.createTV("HairTV")
	if ht != nil {
		ht.watch()
	} else {
		fmt.Println("Not HairFactory")
	}

	m := f.GetFactory("MideaFactory")
	ma := m.createAir("MideaAir")
	if ma != nil {
		ma.use()
	} else {
		fmt.Println("Not MideaFactory")
	}

	mt := m.createTV("MideaTV")
	if mt != nil {
		mt.watch()
	} else {
		fmt.Println("Not MideaFactory")
	}
}
