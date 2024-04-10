package main

import "github.com/labstack/echo/v4"

type Car struct {
	Name string
	Price float64
}

var cars []Car

func CreateCars (){
	cars = append(cars, Car{Name: "Ferrari", Price: 3000000})
	cars = append(cars, Car{Name: "Mercedes", Price: 4000000})
	cars = append(cars, Car{Name: "Porsche", Price: 5000000})
}

func main (){
	CreateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":8000"))
}

func getCars (c echo.Context) error {
	return c.JSON(200, cars)
}