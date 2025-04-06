package k8s

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	request1 "github.com/infraboard/mcube/v2/pb/request"
	"github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/types"
)

const (
	AppName    = "k8s"
	SECRET_KEY = "23gs6gxHrz1kNEvshRmunkXbwIiaEcYfh+EMu+e9ewA="
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	QueryCluster(context.Context, *QueryClusterRequest) (types.Set[*Cluster], error)
	DescribeCluster(context.Context, *DescribeClusterRequest) (*Cluster, error)
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*Cluster, error)
}

type QueryClusterRequest struct {
	// 资源范围
	Scope *resource.Scope `json:"scope"`
	// 资源标签过滤
	Filters []*resource.LabelRequirement `json:"filters"`
	// 分页参数
	Page *request.PageRequest `json:"page"`
	// 关键字参数
	Keywords string `json:"keywords"`
	// 集群所属厂商
	Vendor string `json:"vendor"`
	// 集群所属地域
	Region string `json:"region"`
	// 集群Id列表
	ClusterIds []string `json:"cluster_ids"`
}

type DescribeClusterRequest struct {
	// Cluster id
	Id string `json:"id"`
}

type UpdateClusterRequest struct {
	// Cluster id
	Id string `json:"id"`
	// 更新模式
	UpdateMode request1.UpdateMode `json:"update_mode"`
	// 更新人
	UpdateBy string `json:"update_by"`
	// 更新时间
	UpdateAt int64 `json:"update_at"`
	// 更新的书本信息
	Spec *CreateClusterRequest `json:"spec"`
}

type DeleteClusterRequest struct {
	// Cluster id
	Id string `json:"id"`
}
