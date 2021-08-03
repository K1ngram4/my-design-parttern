package _0_simple_factory

import (
	"fmt"
	"reflect"
)

type Car interface {
	GetCarType() string
}

func NewCar(carType string) Car {
	switch carType {
	case "bmw":
		return &bmw{}
	case "tesla":
		return &tesla{}
	default:
		fmt.Println("unknown carType")
		return nil
	}
}

type bmw struct{}

func (c *bmw) GetCarType() string {
	return reflect.TypeOf(*c).Name()
}

type tesla struct{}

func (t *tesla) GetCarType() string {
	return reflect.TypeOf(*t).Name()
}
