// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/clusterpedia-io/api/cluster/v1alpha2"
	"github.com/clusterpedia-io/client-go/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ClusterV1alpha2Interface interface {
	RESTClient() rest.Interface
	PediaClustersGetter
}

// ClusterV1alpha2Client is used to interact with features provided by the cluster.clusterpedia.io group.
type ClusterV1alpha2Client struct {
	restClient rest.Interface
}

func (c *ClusterV1alpha2Client) PediaClusters() PediaClusterInterface {
	return newPediaClusters(c)
}

// NewForConfig creates a new ClusterV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*ClusterV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *ClusterV1alpha2Client {
	return &ClusterV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ClusterV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
