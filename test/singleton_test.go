package test

import (
	"my-design-parttern/03_singleton"
	"sync"
	"testing"
)

func TestNewWife(t *testing.T) {
	//我想拥有迪丽热巴
	dilireba := _3_singleton.GetInstance()
	//我还想拥有古力娜扎
	gulinazha := _3_singleton.GetInstance()
	//其实还是迪丽热巴
	if dilireba != gulinazha {
		t.Fatal("instance is not equal")
	}
}

func TestParallelNewWife(t *testing.T) {
	var controlChan = make(chan struct{})
	const num = 100
	var wg = sync.WaitGroup{}
	var wifes = [num]_3_singleton.WifeInterface{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(ii int) {
			<-controlChan
			wifes[ii] = _3_singleton.GetInstance()
			wg.Done()
		}(i)
	}

	close(controlChan)
	wg.Wait()
	//一百个迪丽热巴，都是同一个人，很恐怖
	for i := 1; i < num; i++ {
		if wifes[i] != wifes[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
