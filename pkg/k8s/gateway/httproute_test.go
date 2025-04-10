package gateway_test

import (
	"testing"

	"github.com/infraboard/devops/pkg/k8s/meta"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func TestListHttpRoute(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListHttpRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}

func TestGetHttpRoute(t *testing.T) {
	req := meta.NewGetRequest("coffee")
	v, err := impl.GetHttpRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}
