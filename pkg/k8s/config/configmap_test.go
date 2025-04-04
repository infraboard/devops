package config_test

import (
	"testing"

	"github.com/infraboard/devops/pkg/k8s/meta"
)

func TestListConfigMap(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListConfigMap(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}
