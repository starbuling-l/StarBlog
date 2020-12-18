package logging

import (
	"fmt"
	"time"

	"github.com/starbuling-l/StarBlog/pkg/setting"
)

func GetLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSavePath,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

/*func GetLogFileFullPath() string {
	prefixPath := GetLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filename string) *os.File {
	_, err := os.Stat(filename)
	switch {
	case os.IsExist(err):
		mkdir()
	case os.IsPermission(err):
		log.Fatalf("Permission%v", err)
	}
	handle, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("fail to open file %v", err)
	}
	return handle
}

func mkdir() {
	dir, _ := os.Getwd()
	if err := os.MkdirAll(dir+"/"+GetLogFilePath(), os.ModePerm); err != nil {
		panic(err)
	}
}
*/
