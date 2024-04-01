package main

import "fmt"

type car struct {
	brand string
	model string
	year  int
}

func (c car) getBrand() string {
	return c.brand
}

type motocycle struct {
	brand string
	year  int
}

func (m motocycle) getBrand() string {
	return m.brand
}

type person struct {
	firstName string
	age       int
	vehicle   vehicle
}

type vehicle interface {
	getBrand() string
}

func (p *person) drove() {
	//p.firstName = "Wellington" (muda somente aqui sem o pointer)
	p.firstName = "Wellington"
	fmt.Println(p.firstName, "drove", p.vehicle.getBrand())
}

func main() {
	ford := car{
		brand: "ford",
		model: "Ka",
		year:  2021,
	}

	yamaha := motocycle{
		brand: "yamaha",
		year:  2020,
	}

	wesley := person{
		firstName: "wesley",
		age:       20,
		vehicle:   yamaha,
	}

	wesley.drove()
	fmt.Println("car's brand:", ford.brand)

	//fmt.Println("First Name:", wesley.firstName, "Car:", wesley.car.model)
}
