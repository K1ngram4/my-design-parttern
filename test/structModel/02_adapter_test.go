package structModel

import (
	"my-design-parttern/structModel/adapter"
	"testing"
)

func TestCoffeeMachine(t *testing.T) {
	var expect = "run with 220V"
	var a = adapter.NewChinaElectricAdapter()
	var coffeeMachine = adapter.NewAmericanElectricWithAdapter(a)
	var ret = coffeeMachine.RunWith110V()
	if ret != expect {
		t.Fatalf("TestCoffeeMachine() failed,expect : %s ,actually : %s", expect, ret)
	}
}
