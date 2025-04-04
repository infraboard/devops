package workload_test

import (
	"testing"

	"github.com/infraboard/devops/pkg/k8s/meta"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
	v1 "k8s.io/api/batch/v1"
)

func TestListJob(t *testing.T) {
	req := meta.NewListRequest()
	req.Namespace = "default"
	list, err := impl.ListJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	for _, v := range list.Items {
		t.Log(tools.MustToYaml(v))
	}
}

func TestGetJob(t *testing.T) {
	req := meta.NewGetRequest(conf.C.MCENTER_BUILD_TASK_ID)
	req.Namespace = "default"
	ins, err := impl.GetJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	t.Log(tools.MustToYaml(ins))
}

func TestCreateJob(t *testing.T) {
	job := &v1.Job{}
	tools.MustReadYamlFile("test/job.yml", job)
	job, err := impl.CreateJob(ctx, job)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(job)
}

func TestDeleteJob(t *testing.T) {
	req := meta.NewDeleteRequest("cfkeo05s99bvio4olvvg")
	err := impl.DeleteJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
