package impl

import (
	"context"
	"time"

	"github.com/infraboard/devops/mpaas/k8s"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/types"
	k8s_client "github.com/infraboard/mpaas/provider/k8s"
)

// CreateCluster implements k8s.Service.
func (s *K8sServiceImpl) CreateCluster(ctx context.Context, in *k8s.CreateClusterRequest) (*k8s.Cluster, error) {
	ins, err := k8s.NewCluster(in)
	if err != nil {
		return nil, exception.NewBadRequest("validate create cluster error, %s", err)
	}

	// 连接集群检查状态
	s.checkStatus(ins)
	if err := ins.IsAlive(); err != nil {
		return nil, err
	}

	// 加密
	err = ins.EncryptedKubeConf()
	if err != nil {
		return nil, err
	}

	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *K8sServiceImpl) checkStatus(ins *k8s.Cluster) {
	client, err := k8s_client.NewClient(ins.KubeConfig)
	if err != nil {
		ins.Status.Message = err.Error()
		return
	}

	if ctx := client.CurrentContext(); ctx != nil {
		ins.Id = ctx.Cluster
		ins.ServerInfo.AuthUser = ctx.AuthInfo
	}

	if k := client.CurrentCluster(); k != nil {
		ins.ServerInfo.Server = k.Server
	}

	// 检查凭证是否可用
	ins.Status.CheckAt = time.Now().Unix()
	v, err := client.ServerVersion()
	if err != nil {
		ins.Status.IsAlive = false
		ins.Status.Message = err.Error()
	} else {
		ins.Status.IsAlive = true
		ins.ServerInfo.Version = v
	}
}

// DeleteCluster implements k8s.Service.
func (s *K8sServiceImpl) DeleteCluster(context.Context, *k8s.DeleteClusterRequest) (*k8s.Cluster, error) {
	panic("unimplemented")
}

// DescribeCluster implements k8s.Service.
func (s *K8sServiceImpl) DescribeCluster(ctx context.Context, in *k8s.DescribeClusterRequest) (*k8s.Cluster, error) {
	ins, err := s.get(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if err := ins.DecryptedKubeConf(); err != nil {
		return nil, err
	}
	return ins, nil
}

// QueryCluster implements k8s.Service.
func (s *K8sServiceImpl) QueryCluster(context.Context, *k8s.QueryClusterRequest) (types.Set[*k8s.Cluster], error) {
	panic("unimplemented")
}

// UpdateCluster implements k8s.Service.
func (s *K8sServiceImpl) UpdateCluster(context.Context, *k8s.UpdateClusterRequest) (*k8s.Cluster, error) {
	panic("unimplemented")
}
