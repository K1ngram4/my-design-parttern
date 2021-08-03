package _1_facade

const (
	registerSuccess    = "register success"
	paySuccess         = "pay success"
	getMedicineSuccess = "get medicine success"
)

//接待员
type Receptionist interface {
	//接待员帮我们挂号，缴费，取药
	HelpSeeDoctor() string
}

//接待员李华
type ReceptionistLi struct {
	//李华有三个成员，分别是挂号、缴费、取药
	//这里用接口类型，而不是结构体,如果外部需要使用这些细节功能，则可以直接调用，比如张三只要拿个药，则只要GetMedicine接口和函数就足够
	r Register
	p Pay
	m GetMedicine
}

//接待员处理主函数
func (r *ReceptionistLi) HelpSeeDoctor() string {
	//(1)首先挂号
	if r.r.Register() != registerSuccess {
		return ""
	}
	//(2)挂号成功，付款
	if r.p.Pay() != paySuccess {
		return ""
	}
	//(3)付款成功，取药
	if r.m.GetMedicine() != getMedicineSuccess {
		return ""
	}
	//(4)返回成功
	return "ok"
}

type Register interface {
	Register() string
}

func NewRegister() Register {
	return &r{}
}

type r struct{}

func (r *r) Register() string {
	return registerSuccess
}

type Pay interface {
	Pay() string
}

func NewPay() Pay {
	return &p{}
}

type p struct{}

func (p *p) Pay() string {
	return paySuccess
}

type GetMedicine interface {
	GetMedicine() string
}

func NewGetMedicine() GetMedicine {
	return &gM{}
}

type gM struct{}

func (gm *gM) GetMedicine() string {
	return getMedicineSuccess
}

// 对外暴露构造方法
func NewReceptionist() Receptionist {
	return &ReceptionistLi{
		r: NewRegister(),
		p: NewPay(),
		m: NewGetMedicine(),
	}
}
