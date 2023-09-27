package provider

import (
	logger "github.com/sirupsen/logrus"
	"io"
	"manatee-publish/internal/model"
	"os"
	"path/filepath"
)

type FileProvider struct {
	model.OssConfig
}

func (o *FileProvider) GetOssConfigID() uint64 {
	return o.OssConfig.ID
}

func (p *FileProvider) Auth() error {
	return nil
}

func (p *FileProvider) Close() error {
	return nil
}

func (p *FileProvider) PutObject(objName string, reader io.Reader) error {
	fp := filepath.Join(p.Bucket, p.Region, p.SubDir, objName)
	dir, _ := filepath.Split(fp)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logger.Errorln(err)
			return err
		}
	}
	f, err := os.Create(fp)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, reader)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func (p *FileProvider) GetObject(objName string) (io.ReadCloser, error) {
	// 从本地读取
	f, err := os.Open(filepath.Join(p.Bucket, p.Region, p.SubDir, objName))
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return f, nil
}
