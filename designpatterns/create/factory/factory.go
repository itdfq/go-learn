package main

import "fmt"

type CarFactory interface {
	product() string
}

type BmwCar struct {
}
type BenzCar struct {
}

func (b BenzCar) product() string {
	fmt.Println("Benz车")
	return "Benz"
}

func (b BmwCar) product() string {
	fmt.Println("BmwCar车")
	return "BmwCar"
}

func main() {
	bmw := getCar("BmwCar")
	Benz := getCar("BmwCar")
	bmw.product()
	Benz.product()
}

func getCar(band string) CarFactory {
	if band == "BmwCar" {
		return new(BmwCar)
	}
	if band == "BenzCar" {
		return new(BenzCar)
	}
	return nil
}
