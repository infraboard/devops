package impl_test

import (
	"context"

	"github.com/infraboard/devops/cmdb/secret"
	"github.com/infraboard/devops/test"
)

var (
	ctx = context.Background()
	svc = secret.GetService()
)

func init() {
	test.SetUp()
}
