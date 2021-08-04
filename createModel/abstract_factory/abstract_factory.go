package abstract_factory

import "fmt"

// OrderMainDao 抽象产品1
type OrderMainDao interface {
	SaveOrderMain()
}

// OrderDetailDao 抽象产品2
type OrderDetailDao interface {
	SaveOrderDetail()
}

//DaoFactory 抽象工厂
type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

// RDBDaoFactory 具体工厂1
type RDBDaoFactory struct{}

func (RDBDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (RDBDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

// RDBMainDao 具体产品1
type RDBMainDao struct{}

func (RDBMainDao) SaveOrderMain() {
	fmt.Println("rdb main save")
}

// RDBDetailDao 具体产品2
type RDBDetailDao struct{}

func (RDBDetailDao) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

// XMLFactory 具体工厂2
type XMLFactory struct{}

func (XMLFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (XMLFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}

// XMLMainDao 具体产品1
type XMLMainDao struct{}

func (XMLMainDao) SaveOrderMain() {
	fmt.Println("XML main save")
}

// XMLDetailDao 具体产品2
type XMLDetailDao struct{}

func (XMLDetailDao) SaveOrderDetail() {
	fmt.Println("XML detail save")
}
