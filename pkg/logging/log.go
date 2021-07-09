package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

type FileLog struct {

}

func (FileLog) Create() Loger {
	return &FileLog{}
}

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (FileLog) Setup()  {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func (FileLog) Debug(v ...interface{})  {
	SetPrefix(DEBUG)
	logger.Println(v)
}

func (FileLog) Info(v ...interface{})  {
	SetPrefix(INFO)
	logger.Println(v)
}

func (FileLog) Warn(v ...interface{})  {
	SetPrefix(WARNING)
	logger.Println(v)
}

func (FileLog) Error(v ...interface{})  {
	SetPrefix(ERROR)
	logger.Println(v)
}

func (FileLog) Fatal(v ...interface{})  {
	SetPrefix(FATAL)
	logger.Println(v)
}

func SetPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}






















