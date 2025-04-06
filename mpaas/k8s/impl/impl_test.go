package impl_test

import (
	"context"

	"github.com/infraboard/devops/mpaas/k8s"
	"github.com/infraboard/devops/test"
)

var (
	ctx = context.Background()
	svc k8s.Service
)

func init() {
	test.SetUp()
	svc = k8s.GetService()
}
