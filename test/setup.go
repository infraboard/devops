package test

import (
	"fmt"
	"os"

	"github.com/infraboard/mcube/v2/ioc"

	_ "github.com/infraboard/devops/cmdb"
)

func SetUp() {
	fmt.Println(os.Getwd())
	ioc.DevelopmentSetupWithPath(os.Getenv("workspaceFolder"))
}
