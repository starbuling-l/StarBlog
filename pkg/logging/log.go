package logging

import (
	"fmt"
	"github.com/starbuling-l/StarBlog/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := GetLogFilePath()
	fileName := GetLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err:%v", err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)

}

func Info(v ...interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}

func Warning(v ...interface{}) {
	SetPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	SetPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	logger.Fatal(v)
}

func SetPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s] [%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
