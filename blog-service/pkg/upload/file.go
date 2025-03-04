package upload

import (
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/pkg/util"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// 获取文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// 获取扩展名
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	//return os.IsNotExist(err)
	return errors.Is(err, fs.ErrNotExist)
}

// 检查是否包含后缀名
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)

	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}
	return false
}

// 检查最大大小
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, fs.ErrPermission)
}

// 创建保存路径
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.Mkdir(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// 保存文件

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
