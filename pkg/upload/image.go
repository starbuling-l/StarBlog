package upload

import (
	"fmt"
	"github.com/starbuling-l/StarBlog/pkg/file"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimPrefix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "\\" + GetImagePath() + name
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImage check if the file exists
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		return false
		log.Printf("file.GetSize err:%v", err)
	}
	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err:%v", err)
	}

	err = file.IsNotExistMKDir(dir + "\\" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMKDir err:%v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission err:%v", err)
	}
	return nil
}
