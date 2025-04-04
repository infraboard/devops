package event_test

import (
	"context"

	"github.com/infraboard/devops/pkg/k8s"
	"github.com/infraboard/devops/pkg/k8s/event"
)

var (
	impl *event.Client
	ctx  = context.Background()
)

func init() {
	client, err := k8s.NewClientFromFile("../kube_config.yml")
	if err != nil {
		panic(err)
	}
	impl = client.Event()
}
