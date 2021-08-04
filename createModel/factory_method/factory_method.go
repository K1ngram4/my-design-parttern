package factory_method

// SingleDog 是一个程序员接口类型
type SingleDog interface {
	Study(string)
	Typing() string
}

// SingleDogFactory 是一个程序员梦工厂，职责是创造程序员
type SingleDogFactory interface {
	Create() SingleDog
}

// Skill 每个程序员都需要一个技能，那就是学习，具体学习啥技能，就给他标记成啥程序员
type Skill struct {
	flag string
}

// Study 实现了Study函数
func (s *Skill) Study(f string) {
	s.flag = f
}

// GolangProgrammerFactory 是专门创造golang程序员的工厂
type GolangProgrammerFactory struct{}

func (GolangProgrammerFactory) Create() SingleDog {
	return &GolangProgrammer{
		&Skill{},
	}
}

// GolangProgrammer 是golang程序员，他会学习和写代码，学习技能使用组合(相当于继承)Skill
type GolangProgrammer struct {
	*Skill
}

func (g *GolangProgrammer) Typing() string {
	return g.flag + "摸鱼"
}

// JavaProgrammerFactory 同样的方式再来一个Java的程序员梦工厂
type JavaProgrammerFactory struct{}

func (JavaProgrammerFactory) Create() SingleDog {
	return &JavaProgrammer{
		&Skill{},
	}
}

type JavaProgrammer struct {
	*Skill
}

func (j *JavaProgrammer) Typing() string {
	return j.flag + "划水"
}
