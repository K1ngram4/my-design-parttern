package structModel

import (
	"my-design-parttern/structModel/facade"
	"testing"
)

//可以让李华帮我看病
func TestSeeDoctor(t *testing.T) {
	var expect = "ok"
	//创建一个李华
	var lihua = facade.NewReceptionist()
	//帮我看病
	var ret = lihua.HelpSeeDoctor()
	if ret != expect {
		t.Fatalf("TestSeeDoctor() failed,expect : %s ,actually : %s", expect, ret)
	}
}

//有时我知道自己只是感冒了，只需要去拿药，不需要请李华
func TestGetMedicine(t *testing.T) {
	var expect = "get medicine success"
	var gm = facade.NewGetMedicine()
	var ret = gm.GetMedicine()
	if ret != expect {
		t.Fatalf("TestGetMedicine() failed,expect : %s ,actually : %s", expect, ret)
	}
}
