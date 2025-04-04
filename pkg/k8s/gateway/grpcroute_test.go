package gateway_test

import (
	"testing"

	"github.com/infraboard/devops/pkg/k8s/meta"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func TestListGRPCRouteList(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListGRPCRouteList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}
