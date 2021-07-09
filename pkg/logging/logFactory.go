package logging

var Logor Loger

type Loger interface {
	Setup()
	Info(v ...interface{})
	//Error(v ...interface{})
}

type LogerFactory interface {
	Create() Loger
}

func Setup(factory LogerFactory) {
	Logor = factory.Create()
	Logor.Setup()
}

func NewLog() Loger {
	return Logor
}

