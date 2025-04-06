package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/devops/mpaas/k8s"
)

func TestCreateCluster(t *testing.T) {
	// 读取kubeconf
	filePath := os.Getenv("workspaceFolder") + "/etc/kubeconf.yaml"
	kubeconf, err := os.ReadFile(filePath)
	if err != nil {
		t.Log(err)
	}

	ins, err := svc.CreateCluster(ctx, &k8s.CreateClusterRequest{
		Provider:    "docker-desktop",
		Region:      "local",
		Name:        "decker desktop",
		Description: "本地调试使用",
		KubeConfig:  string(kubeconf),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeCluster(t *testing.T) {
	ins, err := svc.DescribeCluster(ctx, &k8s.DescribeClusterRequest{
		Id: "docker-desktop",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}
