package adapter

//美国电器，可以在110V电压运行
type AmericanElectric interface {
	RunWith110V() string
}

//返回一个带有适配器的美国电器--咖啡机
func NewAmericanElectricWithAdapter(adapter ChinaElectric) AmericanElectric {
	return &coffeeMachine{adapter}
}

//咖啡机结构体，实际上包括了一个适配器，这个适配器实现ChinaElectric接口，可以在220V电压运行
type coffeeMachine struct {
	ChinaElectric
}

//修改函数，使得能够在220V环境下运行
func (a *coffeeMachine) RunWith110V() string {
	return a.RunWith220V()
}

//中国电器，可以在220V电压运行
type ChinaElectric interface {
	RunWith220V() string
}

//返回一个适配器，可以在220V电压下运行
func NewChinaElectricAdapter() ChinaElectric {
	return &chinaElectricAdapter{}
}

type chinaElectricAdapter struct{}

func (c *chinaElectricAdapter) RunWith220V() string {
	return "run with 220V"
}
