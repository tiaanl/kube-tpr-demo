package main

import (
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/errors"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/1.4/pkg/runtime/serializer"
	"k8s.io/client-go/1.4/rest"
	"k8s.io/client-go/1.4/tools/clientcmd"
)

func CreateConfig(kubeconfig string) (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	config.GroupVersion = &unversioned.GroupVersion{
		Group:   "third.com",
		Version: "v1",
	}

	config.NegotiatedSerializer = serializer.DirectCodecFactory{
		CodecFactory: api.Codecs,
	}

	return config, nil
}

func CreateClientset(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}

func EnsureNamespace(client *kubernetes.Clientset) error {
	namespace := v1.Namespace{}
	namespace.SetName("third")

	if _, err := client.Namespaces().Create(&namespace); err != nil && !errors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func EnsureThirdPartyResource(client *kubernetes.Clientset) error {
	// Try to get the resource.
	_, err := client.Extensions().ThirdPartyResources().Get("demo.third.com")
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return err
		}

		// The resource doesn't exist, so we create it.
		newResource := v1beta1.ThirdPartyResource{
			Description: "This represents a demo object.",
			Versions: []v1beta1.APIVersion{
				{Name: "v1"},
			},
		}
		newResource.SetNamespace("third")
		newResource.SetName("demo.third.com")

		_, err = client.Extensions().ThirdPartyResources().Create(&newResource)
		if err != nil {
			log.Panic(err)
		}
	}

	return nil
}
