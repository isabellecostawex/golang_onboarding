package main

import "github.com/labstack/echo/v4"

type Car struct {
	Name string
	Price float64
}

var cars []Car

func generateCars (){
	cars = append(cars, Car{Name: "Ferrari", Price: 3000000})
	cars = append(cars, Car{Name: "Mercedes", Price: 4000000})
	cars = append(cars, Car{Name: "Porsche", Price: 5000000})
}

func main (){
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars (c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar (c echo.Context) error{
	car := new(Car)
	if err != c.Bind(car); err != nil {
		retur err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}
