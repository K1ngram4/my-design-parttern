package builder

// Product 是产品
type Product struct {
	partA string
	partB string
	partC string
}

func (p *Product) PartA() string {
	return p.partA
}

func (p *Product) SetPartA(partA string) {
	p.partA = partA
}

func (p *Product) PartB() string {
	return p.partB
}

func (p *Product) SetPartB(partB string) {
	p.partB = partB
}

func (p *Product) PartC() string {
	return p.partC
}

func (p *Product) SetPartC(partC string) {
	p.partC = partC
}

// Builder 是一个抽象建造者
type Builder interface {
	// 创建新产品
	newProduct()
	buildPartA()
	buildPartB()
	buildPartC()
	// 返回产品对象
	getResult() interface{}
}

// Director 是指挥者
type Director struct {
	builder Builder
}

func (d *Director)Construct() interface{} {
	d.builder.newProduct()
	d.builder.buildPartA()
	d.builder.buildPartB()
	d.builder.buildPartC()
	return d.builder.getResult()
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// ConcreteBuilder 是具体的建造者
type ConcreteBuilder struct {
	product *Product
}

func (cb *ConcreteBuilder)newProduct() {
	cb.product = &Product{}
}

func (cb *ConcreteBuilder)buildPartA() {
	cb.product.SetPartA("A部分已构建")
}
func (cb *ConcreteBuilder) buildPartB() {
	cb.product.SetPartB("B部分已构建")
}

func (cb *ConcreteBuilder)buildPartC() {
	cb.product.SetPartC("C部分已构建")
}
// 返回产品对象
func (cb *ConcreteBuilder)getResult() interface{} {
	return cb.product
}



