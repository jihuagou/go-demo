package logging

import "sync"

var (
	Logor Loger
	once sync.Once
)

type Loger interface {
	Setup()
	Info(v ...interface{})
	//Error(v ...interface{})
}

type LogerFactory interface {
	Create() Loger
}

func Setup(factory LogerFactory) {
	once.Do(func() {
		Logor = factory.Create()
		Logor.Setup()
	})

}

func NewLog() Loger {
	return Logor
}

