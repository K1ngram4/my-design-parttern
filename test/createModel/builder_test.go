package createModel

import (
	"my-design-parttern/createModel/builder"
	"testing"
)

func TestBuilder(t *testing.T) {
	b:= &builder.ConcreteBuilder{}
	director := builder.NewDirector(b)
	product := director.Construct()
	var expectA = "A部分已构建"
	var expectB = "B部分已构建"
	var expectC = "C部分已构建"
	var productImpl,_ = product.(*builder.Product)
	if expectA != productImpl.PartA() {
		t.Errorf("expectA: %s ,actually : %s ",expectA,productImpl.PartA())
	}
	if expectB != productImpl.PartB() {
		t.Errorf("expectB: %s ,actually : %s ",expectB,productImpl.PartB())
	}
	if expectC != productImpl.PartC() {
		t.Errorf("expectC: %s ,actually : %s ",expectC,productImpl.PartC())
	}
}
