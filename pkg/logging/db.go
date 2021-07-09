package logging

import "log"

type DbLog struct {

}

func (DbLog) Create() Loger {
	return &DbLog{}
}

func (DbLog) Setup()  {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func (DbLog) Info(v ...interface{})  {
	SetPrefix(ERROR)
	logger.Println(v)
}
