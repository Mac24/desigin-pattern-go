package main

import "fmt"

//抽象产品TV
type AbstractTV interface {
	watch()
}

//具体产品
type HairTV struct {}

type MideaTV struct {}

func (ht *HairTV) watch() {
	fmt.Println("watch HairTV, HairTV watch() method")
}
func (mt *MideaTV) watch() {
	fmt.Println("watch MideaTV, MideaTV watch() method")
}

//抽象产品Air
type AbstractAir interface {
	use()
}

//具体产品
type HairAir struct {}

type MideaAir struct {}

func (ha *HairAir) use() {
	fmt.Println("use HairAir, HairAir use() method")
}

func (ma *MideaAir) use() {
	fmt.Println("use MideaAir, MideaAir use() method")
}

//抽象工厂
type AbstractFactory interface {
	createTV(tvName string) AbstractTV
	createAir(airName string) AbstractAir
}

type HairFactory struct {}

type MideaFactory struct {}

func (h *HairFactory) createTV(tvName string) AbstractTV{
	fmt.Printf("HairFactory create %v\n", tvName)
	return &HairTV{}

}

func (h *HairFactory) createAir(airName string) AbstractAir{
	fmt.Printf("HairFactory create %v\n", airName)
	return &HairAir{}
}

func (m *MideaFactory) createTV(tvName string) AbstractTV {
	fmt.Printf("MideaFactory create %v\n", tvName)
	return &MideaTV{}
}

func (m *MideaFactory) createAir(airName string) AbstractAir{
	fmt.Printf("MideaFactory create %v\n", airName)
	return &MideaAir{}
}

type FactoryProducer struct {}

func (fp *FactoryProducer) GetFactory(factoryName string) AbstractFactory {
	switch factoryName {
	case "HairFactory":
		return &HairFactory{}
	case "MideaFactory":
		return &MideaFactory{}
	default:
		return nil
	}
}

func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}
