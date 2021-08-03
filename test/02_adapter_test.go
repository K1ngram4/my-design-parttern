package test

import (
	"my-design-parttern/02_adapter"
	"testing"
)

func TestCoffeeMachine(t *testing.T) {
	var expect = "run with 220V"
	var a = _2_adapter.NewChinaElectricAdapter()
	var coffeeMachine = _2_adapter.NewAmericanElectricWithAdapter(a)
	var ret = coffeeMachine.RunWith110V()
	if ret != expect {
		t.Fatalf("TestCoffeeMachine() failed,expect : %s ,actually : %s", expect, ret)
	}
}
