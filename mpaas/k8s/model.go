package k8s

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/infraboard/mcube/v2/crypto/cbc"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	k8s_client "github.com/infraboard/mpaas/provider/k8s"
)

type Cluster struct {
	Id string `json:"id" bson:"_id"`
	// 录入时间
	CreateAt int64 `json:"create_at" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `json:"update_at" bson:"update_at"`
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by"`
	// 集群相关信息
	ServerInfo ServerInfo `json:"server_info" bson:",inline"`

	// 集群定义信息
	CreateClusterRequest `bson:",inline"`

	// 集群状态
	Status Status `json:"status" bson:",inline"`

	isEncrypted bool
}

func (r *Cluster) String() string {
	return pretty.ToJSON(r)
}

func (r *Cluster) SetIsEncrypted(v bool) {
	r.isEncrypted = v
}

func (r *Cluster) EncryptedKubeConf() error {
	if r.isEncrypted {
		return nil
	}

	key, err := base64.StdEncoding.DecodeString(SECRET_KEY)
	if err != nil {
		return err
	}

	cipherText, err := cbc.MustNewAESCBCCihper(key).Encrypt([]byte(r.KubeConfig))
	if err != nil {
		return err
	}
	r.KubeConfig = base64.StdEncoding.EncodeToString(cipherText)
	r.SetIsEncrypted(true)
	return nil

}

func (r *Cluster) DecryptedKubeConf() error {
	if r.isEncrypted {
		cipherdText, err := base64.StdEncoding.DecodeString(r.KubeConfig)
		if err != nil {
			return err
		}

		key, err := base64.StdEncoding.DecodeString(SECRET_KEY)
		if err != nil {
			return err
		}

		plainText, err := cbc.MustNewAESCBCCihper(key).Decrypt([]byte(cipherdText))
		if err != nil {
			return err
		}
		r.KubeConfig = string(plainText)
		r.SetIsEncrypted(false)
	}
	return nil
}

func (c *Cluster) GetK8sClient() (*k8s_client.Client, error) {
	return k8s_client.NewClient(c.KubeConfig)
}

func (i *Cluster) IsAlive() error {
	if !i.Status.IsAlive {
		return fmt.Errorf("%s", i.Status.Message)
	}

	return nil
}

type ServerInfo struct {
	// k8s的地址
	Server string `json:"server" bson:"server"`
	// k8s版本
	Version string `json:"version" bson:"version"`
	// 连接用户
	AuthUser string `json:"auth_user" bson:"auth_user"`
}

func NewCluster(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Cluster{
		CreateAt:             time.Now().Unix(),
		UpdateAt:             time.Now().Unix(),
		CreateClusterRequest: *req,
	}, nil
}

type CreateClusterRequest struct {
	// 集群所属域
	Domain string `json:"domain" form:"domain" bson:"domain"`
	// 集群所属空间
	Namespace string `json:"namespace" form:"namespace" bson:"namespace"`
	// 创建人
	CreateBy string `json:"create_by" form:"create_by" bson:"create_by"`
	// 集群提供商
	Provider string `json:"provider" bson:"provider" form:"provider" validate:"required"`
	// 集群所处地域
	Region string `json:"region" bson:"region" form:"region" validate:"required"`
	// 名称
	Name string `json:"name" bson:"name" form:"name" validate:"required"`
	// 集群客户端访问凭证
	KubeConfig string `json:"kube_config" bson:"kube_config" form:"kube_config" validate:"required" mask:",10,10"`
	// 集群描述
	Description string `json:"description" form:"description" bson:"description"`
	// 集群标签, env=prod
	Lables map[string]string `json:"lables" form:"lables" bson:"lables"`
}

func (req CreateClusterRequest) Validate() error {
	return validator.Validate(req)
}

type Status struct {
	// 检查时间
	CheckAt int64 `json:"check_at" bson:"check_at"`
	// API Server是否正常
	IsAlive bool `json:"is_alive" bson:"is_alive"`
	// 异常消息
	Message string `json:"message" bson:"message"`
}
