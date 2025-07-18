package upload

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var MinioClient *Minio // ，

type Minio struct {
	Client *minio.Client
	bucket string
}

func GetMinio(endpoint, accessKeyID, secretAccessKey, bucketName string, useSSL bool) (*Minio, error) {
	if MinioClient != nil {
		return MinioClient, nil
	}
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL, // Set to true if using https
	})
	if err != nil {
		return nil, err
	}
	// bucket
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			// log.Printf("We already own %s\n", bucketName)
		} else {
			return nil, err
		}
	}
	MinioClient = &Minio{Client: minioClient, bucket: bucketName}
	return MinioClient, nil
}

func (m *Minio) UploadFile(file *multipart.FileHeader) (filePathres, key string, uploadErr error) {
	f, openError := file.Open()
	// mutipart.File to os.File
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	filecontent := bytes.Buffer{}
	_, err := io.Copy(&filecontent, f)
	if err != nil {
		global.GVA_LOG.Error("", zap.Any("err", err.Error()))
		return "", "", errors.New(", err:" + err.Error())
	}
	f.Close() //  defer

	//
	ext := filepath.Ext(file.Filename)
	filename := utils.MD5V([]byte(strings.TrimSuffix(file.Filename, ext))) + ext
	if global.GVA_CONFIG.Minio.BasePath == "" {
		filePathres = "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename
	} else {
		filePathres = global.GVA_CONFIG.Minio.BasePath + "/" + time.Now().Format("2006-01-02") + "/" + filename
	}

	// 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Upload the file with PutObject
	info, err := m.Client.PutObject(ctx, global.GVA_CONFIG.Minio.BucketName, filePathres, &filecontent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		global.GVA_LOG.Error("minio", zap.Any("err", err.Error()))
		return "", "", errors.New("minio, err:" + err.Error())
	}
	return global.GVA_CONFIG.Minio.BucketUrl + "/" + info.Key, filePathres, nil
}

func (m *Minio) DeleteFile(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Delete the object from MinIO
	err := m.Client.RemoveObject(ctx, m.bucket, key, minio.RemoveObjectOptions{})
	return err
}
