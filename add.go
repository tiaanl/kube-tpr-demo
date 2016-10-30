package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"k8s.io/client-go/1.4/pkg/api/errors"
	"k8s.io/client-go/1.4/rest"
)

var (
	cmdAdd = cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(optsAdd.demoName) == 0 {
				return fmt.Errorf("--name not specified")
			}

			err := addDemo(optsAdd.demoName)
			if err != nil {
				return err
			}

			return nil
		},
	}

	optsAdd = struct {
		demoName string
	}{}
)

func init() {
	cmdAdd.Flags().StringVar(&optsAdd.demoName, "name", "", "Name of the demo to add")

	cmdRoot.AddCommand(&cmdAdd)
}

func addDemo(name string) error {
	// Create the config.
	config, err := CreateConfig(optsRoot.kubeConfigFile)
	if err != nil {
		return err
	}

	// Create a clientset.
	clientset, err := CreateClientset(config)
	if err != nil {
		return err
	}

	// Make sure the namespace and third party resources are registered.
	if err := EnsureNamespace(clientset); err != nil {
		return err
	}

	if err := EnsureThirdPartyResource(clientset); err != nil {
		return err
	}

	// Create a connection to the cluster.
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		return err
	}

	// Create the demo.
	demo := Demo{
		Spec: DemoSpec{
			Name:        name,
			Description: "Description for " + name + ".",
		},
	}
	demo.APIVersion = "third.com/v1"
	demo.Kind = "Demo"
	demo.SetName(name)
	demo.SetNamespace("third")

	log.Println("Adding demo:", demo)

	newDemo := Demo{}
	err = restClient.Post().
		Namespace("third").
		Resource("demos").
		Body(&demo).
		Do().
		Into(&newDemo)
	if err != nil {
		if errors.IsAlreadyExists(err) {
			log.Printf("Demo %q already exists.", name)
			return nil
		}
		return err
	}

	return nil
}
