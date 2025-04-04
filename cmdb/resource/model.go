package resource

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
)

func NewResourceSet() *types.Set[*Resource] {
	return types.New[*Resource]()
}

func NewResource() *Resource {
	return &Resource{
		Meta: Meta{},
		Spec: Spec{
			Tags:  map[string]string{},
			Extra: map[string]string{},
		},
		Status: Status{},
	}
}

// 资源
// https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/#struct-tags
type Resource struct {
	Meta   `bson:"inline"`
	Spec   `bson:"inline"`
	Status `bson:"inline"`
}

func (r *Resource) String() string {
	return pretty.ToJSON(r)
}

// 元数据，不会变的
type Meta struct {
	// 全局唯一Id, 直接使用个云商自己的Id
	Id string `bson:"_id" json:"id" validate:"required"`
	// 资源所属域
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" validate:"required"`
	// 资源所属空间
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" validate:"required"`
	// 资源所属环境
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env"`
	// 创建时间
	CreateAt int64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at"`
	// 删除时间
	DeteteAt int64 `protobuf:"varint,6,opt,name=detete_at,json=deteteAt,proto3" json:"detete_at"`
	// 删除人
	DeteteBy string `protobuf:"bytes,7,opt,name=detete_by,json=deteteBy,proto3" json:"detete_by"`
	// 同步时间
	SyncAt int64 `protobuf:"varint,8,opt,name=sync_at,json=syncAt,proto3" json:"sync_at" validate:"required"`
	// 同步人
	SyncBy string `protobuf:"bytes,9,opt,name=sync_by,json=syncBy,proto3" json:"sync_by"`
	// 用于同步的凭证ID
	CredentialId string `protobuf:"bytes,10,opt,name=credential_id,json=credentialId,proto3" json:"credential_id"`
	// 序列号
	SerialNumber string `protobuf:"bytes,11,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number"`
}

// 表单数据
type Spec struct {
	// 厂商
	Vendor VENDOR `protobuf:"varint,1,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.VENDOR" json:"vendor"`
	// 资源类型
	ResourceType TYPE `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=infraboard.cmdb.resource.TYPE" json:"resource_type"`
	// 地域
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region"`
	// 区域
	Zone string `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone"`
	// 资源所属账号
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner"`
	// 名称
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	// 种类
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category"`
	// 规格
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	// 描述
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description"`
	// 过期时间
	ExpireAt int64 `protobuf:"varint,10,opt,name=expire_at,json=expireAt,proto3" json:"expire_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,11,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 资源占用Cpu数量
	Cpu int64 `protobuf:"varint,15,opt,name=cpu,proto3" json:"cpu"`
	// GPU数量
	Gpu int64 `protobuf:"varint,16,opt,name=gpu,proto3" json:"gpu"`
	// 资源使用的内存
	Memory int64 `protobuf:"varint,17,opt,name=memory,proto3" json:"memory"`
	// 资源使用的存储
	Storage int64 `protobuf:"varint,18,opt,name=storage,proto3" json:"storage"`
	// 公网IP带宽, 单位M
	BandWidth int32 `protobuf:"varint,19,opt,name=band_width,json=bandWidth,proto3" json:"band_width"`
	// 资源标签
	Tags map[string]string `protobuf:"bytes,25,rep,name=tags,proto3" json:"tags"`
	// 额外的通用属性
	Extra map[string]string `protobuf:"bytes,26,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" gorm:"serializer:json"`
}

// 资源当前状态
type Status struct {
	// 资源当前状态
	Phase string `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase"`
	// 资源当前状态描述
	Describe string `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe"`
	// 实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定
	LockMode string `protobuf:"bytes,3,opt,name=lock_mode,json=lockMode,proto3" json:"lock_mode"`
	// 锁定原因
	LockReason string `protobuf:"bytes,4,opt,name=lock_reason,json=lockReason,proto3" json:"lock_reason"`
	// 资源访问地址
	// 公网地址, 或者域名
	PublicAddress []string `protobuf:"bytes,5,rep,name=public_address,json=publicAddress,proto3" json:"public_address" gorm:"serializer:json"`
	// 内网地址, 或者域名
	PrivateAddress []string `protobuf:"bytes,6,rep,name=private_address,json=privateAddress,proto3" json:"private_address" gorm:"serializer:json"`
}

func (s *Status) GetFirstPrivateAddress() string {
	if len(s.PrivateAddress) > 0 {
		return s.PrivateAddress[0]
	}

	return ""
}
