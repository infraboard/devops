package secret

import (
	"github.com/infraboard/devops/cmdb/resource"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

func (s *Secret) Sync(cb SyncResourceHandleFunc) error {
	switch s.Vendor {
	case resource.VENDOR_TENCENT:
		// 腾讯云的API来进行同步, 云资源(API： https://console.cloud.tencent.com/api/explorer?Product=cvm&Version=2017-03-12&Action=DescribeRegions)

		// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
		// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
		// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
		credential := common.NewCredential(
			s.ApiKey,
			s.ApiSecret,
		)
		// 实例化一个client选项，可选的，没有特殊需求可以跳过
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"

		for i := range s.Regions {
			region := s.Regions[i]
			// 实例化要请求产品的client对象,clientProfile是可选的
			client, _ := lighthouse.NewClient(credential, region, cpf)

			// 实例化一个请求对象,每个接口都会对应一个request对象
			request := lighthouse.NewDescribeInstancesRequest()
			SetLimit(request, s.SyncLimit)
			SetOffset(request, 0)

			hasNext := true
			for hasNext {
				// 返回的resp是一个DescribeInstancesResponse的实例，与请求对象对应
				response, err := client.DescribeInstances(request)
				if err != nil {
					return err
				}

				// 处理拉取到的一页的数据
				for _, ins := range response.Response.InstanceSet {
					cb(ResourceResponse{
						Resource: TransferLighthouseToResource(ins),
					})
				}

				// 当前数据，都没有填满一页，说明后面没有数据
				if *response.Response.TotalCount < *request.Limit {
					hasNext = false
				} else {
					SetOffset(request, GetValue(request.Offset)+GetValue(request.Limit))
				}
			}
		}
	case resource.VENDOR_ALIYUN:
		// 阿里云API
	}

	return nil
}

// 云商数据结构lighthouse.Instance --> Resource
func TransferLighthouseToResource(ins *lighthouse.Instance) *resource.Resource {
	res := resource.NewResource()
	// 具体的转化逻辑
	res.Id = GetValue(ins.InstanceId)
	res.Name = GetValue(ins.InstanceName)
	res.Cpu = GetValue(ins.CPU)
	res.Memory = GetValue(ins.Memory)
	res.Storage = GetValue(ins.SystemDisk.DiskSize)
	res.PrivateAddress = common.StringValues(ins.PrivateAddresses)
	return res
}

func SetOffset(req *lighthouse.DescribeInstancesRequest, v int64) {
	req.Offset = &v
}

func SetLimit(req *lighthouse.DescribeInstancesRequest, v int64) {
	req.Limit = &v
}

func GetValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}

	return *ptr
}
