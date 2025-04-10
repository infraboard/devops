package k8s

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/infraboard/devops/pkg/k8s/admin"
	"github.com/infraboard/devops/pkg/k8s/config"
	"github.com/infraboard/devops/pkg/k8s/event"
	"github.com/infraboard/devops/pkg/k8s/gateway"
	"github.com/infraboard/devops/pkg/k8s/meta"
	"github.com/infraboard/devops/pkg/k8s/network"
	"github.com/infraboard/devops/pkg/k8s/storage"
	"github.com/infraboard/devops/pkg/k8s/workload"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func NewClientFromFile(kubeConfPath string) (*Client, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	kc, err := os.ReadFile(filepath.Join(wd, kubeConfPath))
	if err != nil {
		return nil, err
	}

	return NewClient(string(kc))
}

func NewClient(kubeConfigYaml string) (*Client, error) {
	// 加载kubeconfig配置
	kubeConf, err := clientcmd.Load([]byte(kubeConfigYaml))
	if err != nil {
		return nil, err
	}

	// 构造Restclient Config
	restConf, err := clientcmd.BuildConfigFromKubeconfigGetter("",
		func() (*clientcmdapi.Config, error) {
			return kubeConf, nil
		},
	)
	if err != nil {
		return nil, err
	}

	// 初始化客户端
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}

	items, err := client.ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	return &Client{
		kubeconf:  kubeConf,
		restconf:  restConf,
		client:    client,
		resources: meta.NewApiResourceList(items),
		log:       log.Sub("provider.k8s"),
	}, nil
}

type Client struct {
	kubeconf *clientcmdapi.Config
	restconf *rest.Config
	client   *kubernetes.Clientset
	log      *zerolog.Logger

	resources *meta.ApiResourceList
}

func (c *Client) ServerVersion() (string, error) {
	si, err := c.client.ServerVersion()
	if err != nil {
		return "", err
	}

	return si.String(), nil
}

func (c *Client) ServerResources() *meta.ApiResourceList {
	return c.resources
}

func (c *Client) GetContexts() map[string]*clientcmdapi.Context {
	return c.kubeconf.Contexts
}

func (c *Client) CurrentContext() *clientcmdapi.Context {
	return c.kubeconf.Contexts[c.kubeconf.CurrentContext]
}

func (c *Client) CurrentCluster() *clientcmdapi.Cluster {
	ctx := c.CurrentContext()
	if ctx == nil {
		return nil
	}

	return c.kubeconf.Clusters[ctx.Cluster]
}

// 应用负载
func (c *Client) WorkLoad() *workload.Client {
	return workload.NewWorkload(c.client, c.restconf)
}

// 应用配置
func (c *Client) Config() *config.Client {
	return config.NewConfig(c.client)
}

// 应用存储
func (c *Client) Storage() *storage.Client {
	return storage.NewStorage(c.client)
}

// 应用网络
func (c *Client) Network() *network.Client {
	return network.NewNetwork(c.client)
}

// 应用事件
func (c *Client) Event() *event.Client {
	return event.NewEvent(c.client)
}

// 集群管理
func (c *Client) Admin() *admin.Client {
	return admin.NewAdmin(c.client)
}

func (c *Client) Gateway() *gateway.Client {
	return gateway.NewGateway(c.restconf, c.resources)
}
