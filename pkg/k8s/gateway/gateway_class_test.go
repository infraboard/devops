package gateway_test

import (
	"testing"

	"github.com/infraboard/devops/pkg/k8s/meta"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func TestListGatewayClass(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListGatewayClass(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}
