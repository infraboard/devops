package impl

import (
	"github.com/infraboard/devops/cmdb/resource"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

var _ resource.Service = (*impl)(nil)

// 业务具体实现
type impl struct {
	// 继承模版
	ioc.ObjectImpl

	// 模块子Logger
	log *zerolog.Logger

	//
	col *mongo.Collection
}

// 对象名称
func (i *impl) Name() string {
	return resource.AppName
}

// 初始化
func (i *impl) Init() error {
	// 对象
	i.log = log.Sub(i.Name())
	i.log.Debug().Msgf(ioc_mongo.Get().Database)
	i.col = ioc_mongo.DB().Collection("resources")
	return nil
}
