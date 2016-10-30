package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	cmdInit = cobra.Command{
		Use: "init",
		RunE: func(*cobra.Command, []string) error {
			return ensureNamespaceAndThirdPartyResource()
		},
	}
)

func init() {
	cmdRoot.AddCommand(&cmdInit)
}

func ensureNamespaceAndThirdPartyResource() error {
	// Create the config.
	config, err := CreateConfig(optsRoot.kubeConfigFile)
	if err != nil {
		return err
	}

	clientset, err := CreateClientset(config)
	if err != nil {
		return err
	}

	if err := EnsureNamespace(clientset); err != nil {
		return err
	}

	log.Println("Namespace created.")

	if err := EnsureThirdPartyResource(clientset); err != nil {
		return err
	}

	log.Println("ThirdPartyResource created.")

	return nil
}
