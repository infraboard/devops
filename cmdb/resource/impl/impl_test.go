package impl_test

import (
	"context"

	"github.com/infraboard/devops/cmdb/resource"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	ctx = context.Background()
	svc = resource.GetService()
)

func init() {
	ioc.DevelopmentSetup()
}
