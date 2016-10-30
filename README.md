# kube-tpr-demo

Demo for using Kubernetes' ThirdPartyResource resources through the [client-go](https://github.com/kubernetes/client-go)
library.

## Installation

Check out the repository to your go workspace, e.g. `~/go/src/github.com/tiaanl/kube-tpr-demo`

Build the binary:

```bash
# Install glide to manage dependencies
go get github.com/Masterminds/glide

# Install dependencies
glide install

# Build the binary
go install -v
```

# Getting Started

Register the Namespace and ThirdPartyResource in kubernetes:

```bash
third init
```

Add a new demo:

```bash
third add --name=demo1
third add --name=demo2
```

Get a list of all the demos:

```bash
third list
```

Watch demos as they get created:

```bash
third watch

# In another terminal start adding new demos
third add --name=demo3
```
