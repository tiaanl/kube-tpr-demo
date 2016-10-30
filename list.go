package main

import (
	"log"

	"github.com/spf13/cobra"
	"k8s.io/client-go/1.4/dynamic"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/api/v1"
)

var (
	cmdList = cobra.Command{
		Use: "list",
		RunE: func(*cobra.Command, []string) error {
			return listDemos()
		},
	}
)

func init() {
	cmdRoot.AddCommand(&cmdList)
}

func listDemos() error {
	// Create the config.
	config, err := CreateConfig(optsRoot.kubeConfigFile)
	if err != nil {
		return err
	}

	resource := unversioned.APIResource{
		Name:       "demos",
		Namespaced: true,
	}

	client, err := dynamic.NewClient(config)
	if err != nil {
		return err
	}

	items, err := client.Resource(&resource, "third").List(&v1.ListOptions{})
	if err != nil {
		return err
	}

	log.Println(items)

	return nil
}
