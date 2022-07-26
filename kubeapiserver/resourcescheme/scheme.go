package resourcescheme

import (
	unstructuredresourcescheme "github.com/clusterpedia-io/fake-apiserver/kubeapiserver/resourcescheme/unstructured"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

var (
	LegacyResourceScheme         = legacyscheme.Scheme
	LegacyResourceCodecs         = legacyscheme.Codecs
	LegacyResourceParameterCodec = legacyscheme.ParameterCodec

	CustomResourceScheme = unstructuredresourcescheme.NewScheme()
	CustomResourceCodecs = unstructured.UnstructuredJSONScheme
)
