package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var AssetVersionService = newAssetVersionService()

type assetVersionService struct {
	common.BaseService[*model.AssetVersion]
}

func newAssetVersionService() *assetVersionService {
	s := new(assetVersionService)
	s.BaseService = common.BaseService[*model.AssetVersion]{
		Service: s,
	}
	return s
}

func (*assetVersionService) Model() *model.AssetVersion {
	return new(model.AssetVersion)
}
