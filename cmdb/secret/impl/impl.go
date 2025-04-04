package impl

import (
	"github.com/infraboard/devops/cmdb/secret"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	ioc.Controller().Registry(&SecretServiceImpl{})
}

var _ secret.Service = (*SecretServiceImpl)(nil)

type SecretServiceImpl struct {
	ioc.ObjectImpl
	col *mongo.Collection
}

func (s *SecretServiceImpl) Name() string {
	return secret.AppName
}

func (s *SecretServiceImpl) Init() error {
	// 定义使用的集合
	s.col = ioc_mongo.DB().Collection("secrets")
	return nil
}
