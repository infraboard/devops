package impl_test

import (
	"testing"

	"github.com/infraboard/devops/cmdb/resource"
)

func TestSave(t *testing.T) {
	res := resource.NewResource()
	res.Id = "test01"
	res.ResourceType = resource.TYPE_HOST
	res.Name = "test01"
	res.Cpu = 2
	res.Memory = 4096
	res.Storage = 50
	res.Phase = "Stopped"
	res.PrivateAddress = append(res.PrivateAddress, "10.10.10.2")
	res.Tags["app"] = "app01"
	res, err := svc.Save(ctx, res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestSearch(t *testing.T) {
	req := resource.NewSearchRequest()
	v, err := svc.Search(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestDeleteResource(t *testing.T) {
	req := resource.NewDeleteResourceRequest()
	req.ResourceIds = append(req.ResourceIds, "test01")
	err := svc.DeleteResource(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
