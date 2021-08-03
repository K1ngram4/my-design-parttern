package test

import (
	"my-design-parttern/00_simple_factory"
	"testing"
)

func TestNewTesla(t *testing.T) {
	var expect = "tesla"
	var teslaNo1 = _0_simple_factory.NewCar(expect)
	if teslaNo1 == nil {
		t.Fatal("TestNewTesla() failed,teslaNo1 is nil")
		return
	}
	var typeTeslaNo1 = teslaNo1.GetCarType()
	if expect != typeTeslaNo1 {
		t.Fatalf("TestNewTesla() failed,expect : %s ,actually : %s", expect, typeTeslaNo1)
	}
}
