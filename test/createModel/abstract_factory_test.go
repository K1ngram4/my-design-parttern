package createModel

import "my-design-parttern/createModel/abstract_factory"

func saveMainAndDetail(factory abstract_factory.DaoFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func ExampleRDBFactory() {
	var factory abstract_factory.DaoFactory
	factory = &abstract_factory.RDBDaoFactory{}
	saveMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXMLFactory() {
	var factory abstract_factory.DaoFactory
	factory = &abstract_factory.XMLFactory{}
	saveMainAndDetail(factory)
	// Output:
	// XML main save
	// XML detail save
}
