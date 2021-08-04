package createModel

import (
	"my-design-parttern/createModel/factory_method"
	"testing"
)

func TestSingleDog(t *testing.T) {
	var expect1 = "Golang摸鱼"
	var expect2 = "Java划水"
	var factory factory_method.SingleDogFactory
	factory = factory_method.GolangProgrammerFactory{}
	var golangProgrammer = factory.Create()
	golangProgrammer.Study("Golang")
	var res = golangProgrammer.Typing()
	if res != expect1 {
		t.Fatalf("TestSingleDog()failed,expect:%s,acturelly:%s", expect1, res)
	}

	factory = factory_method.JavaProgrammerFactory{}
	var javaProgrammer = factory.Create()
	javaProgrammer.Study("Java")
	res = javaProgrammer.Typing()
	if res != expect2 {
		t.Fatalf("TestSingleDog()failed,expect:%s,acturelly:%s", expect2, res)
	}
}
