package main

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"k8s.io/client-go/1.4/dynamic"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/watch"
)

var (
	cmdWatch = cobra.Command{
		Use: "watch",
		RunE: func(*cobra.Command, []string) error {
			return watchForChanges()
		},
	}
)

func init() {
	cmdRoot.AddCommand(&cmdWatch)
}

func watchForChanges() error {
	log.Println("Watching for changes to demos...")

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
	watcher, err := client.Resource(&resource, "third").Watch(&v1.ListOptions{})
	if err != nil {
		return err
	}

	for {
		result := <-watcher.ResultChan()
		if result.Type == watch.Error || result.Type == "" {
			break
		}

		data, err := json.MarshalIndent(result.Object, "", "  ")
		if err != nil {
			return err
		}

		log.Println(string(data))
	}

	return nil
}
