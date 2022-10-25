package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Model string
	Price float64
}

// func main() {

// 	// carro := Car{"Punto", "Fiat"}
// 	// carro.Andar()

// 	// result, err := soma(20, 2)
// 	// fmt.Println(result, err)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	generateCars()
// 	e := echo.New()
// 	e.GET("/cars", getCars)
// 	e.POST("/cars", createCar)
// 	e.Logger.Fatal(e.Start(":9876"))

// }

var cars []Car

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func generateCars() {
	cars = append(cars, Car{Name: "Uno", Model: "Fiat", Price: 8000})
	cars = append(cars, Car{Name: "Strada", Model: "Fiat", Price: 18000})
	cars = append(cars, Car{Name: "Argo", Model: "Fiat", Price: 80000})
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	saveCar(*car)
	// cars = append(cars, *car)
	return c.JSON(200, cars)
}

func (c Car) Andar() {
	fmt.Println("O carro", c.Name, "está andando")
}

func soma(a int, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("Soma maior que 10")
	}
	return a + b, nil
}

func saveCar(car Car) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, model, price) VALUES ($1, $2, $3)")

	// fmt.Println(err) erro com a senha do mysql

	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price, car.Model)
	if err != nil {
		return err
	}
	return nil

}
