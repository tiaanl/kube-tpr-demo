package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	cmdRoot = cobra.Command{
		Use: "third",
	}
	optsRoot struct{
		kubeConfigFile string
	}
)

func init() {
	cmdRoot.PersistentFlags().StringVar(&optsRoot.kubeConfigFile, "kubeconfig", "./config", "Kubernetes config file.")
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}
