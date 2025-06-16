package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

var mu sync.Mutex

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description:
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//
	ext := filepath.Ext(file.Filename)
	//
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	//
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	//
	mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	//
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filepath := global.GVA_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() //
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() //  defer

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() //  defer

	_, copyErr := io.Copy(out, f) // （）
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description:
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	//  key
	if key == "" {
		return errors.New("key")
	}

	//  key
	if strings.Contains(key, "..") || strings.ContainsAny(key, `\/:*?"<>|`) {
		return errors.New("key")
	}

	p := filepath.Join(global.GVA_CONFIG.Local.StorePath, key)

	//
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return errors.New("")
	}

	//
	mu.Lock()
	defer mu.Unlock()

	err := os.Remove(p)
	if err != nil {
		return errors.New(": " + err.Error())
	}

	return nil
}
