package singleton

import "sync"

var (
	w    *wife
	once sync.Once
)

type WifeInterface interface {
	DoSth() string
}

type wife struct{}

func (w *wife) DoSth() string {
	return "feel good"
}

func GetInstance() WifeInterface {
	once.Do(func() {
		w = &wife{}
	})
	return w
}
