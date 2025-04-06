package impl

import (
	"github.com/infraboard/devops/mpaas/k8s"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	ioc.Controller().Registry(&K8sServiceImpl{})
}

var _ k8s.Service = (*K8sServiceImpl)(nil)

type K8sServiceImpl struct {
	ioc.ObjectImpl
	col *mongo.Collection
}

func (s *K8sServiceImpl) Name() string {
	return k8s.AppName
}

func (s *K8sServiceImpl) Init() error {
	// 定义使用的集合
	s.col = ioc_mongo.DB().Collection("k8s")
	return nil
}
