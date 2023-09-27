package provider

import (
	"io"
	"manatee-publish/internal/model"
	"strings"
)

type OsManager interface {
	Auth() error
	Close() error
	PutObject(objName string, reader io.Reader) error
	GetObject(objName string) (io.ReadCloser, error)
	GetOssConfigID() uint64
}

func GetProvider(c *model.OssConfig) (m OsManager) {
	switch strings.ToLower(c.Type) {
	case "minio":
		m = &Minio{
			OssConfig: *c,
		}
	case "ks3":
		m = &KS3{
			OssConfig: *c,
		}
	case "obs":
		m = &OBS{
			OssConfig: *c,
		}
	case "oss":
		m = &OSS{
			OssConfig: *c,
		}
	case "azure":
		m = &Azure{
			OssConfig: *c,
		}
	default:
		m = &FileProvider{
			OssConfig: *c,
		}
	}
	return
}
