package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/utils/env"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/capa-karpenter-taint-remover/internal/taintsfilter"
)

const (
	retryAttempts = 5

	defaultUnwantedTaints = "node.cluster.x-k8s.io/uninitialized"
)

var (
	scheme = runtime.NewScheme()

	unwantedTaints = make([]string, 0)
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

func main() {
	var unwantedTaintsRaw string
	flag.StringVar(&unwantedTaintsRaw, "unwanted-taints", defaultUnwantedTaints, "")
	flag.Parse()

	unwantedTaints = strings.Split(unwantedTaintsRaw, ",")
	fmt.Printf("The list of onwanted taints is:\n%s\n", strings.Join(unwantedTaints, "\n"))

	taintsFilter := taintsfilter.New(unwantedTaints)

	nodeName := env.GetString("NODE_NAME", "")
	if nodeName == "" {
		fmt.Printf("ERROR: NODE_NAME env cannot be empty\n")
		os.Exit(1)
	}

	config, err := ctrl.GetConfig()
	if err != nil {
		fmt.Printf("ERROR: failed to get config for controlelr runtime client\n")
		panic(err)
	}

	ctrlClient, err := client.New(config, client.Options{})
	if err != nil {
		fmt.Printf("ERROR: failed to create controller runtime client\n")
		panic(err)
	}

	ctx := context.Background()

	patch := func() error {
		var node v1.Node

		err = ctrlClient.Get(ctx, client.ObjectKey{Name: nodeName}, &node)
		if err != nil {
			fmt.Printf("ERROR: failed to get node %s\n", nodeName)
			return err
		}

		// Check if node is managed by karpenter.
		if node.Labels["managed-by"] != "karpenter" {
			fmt.Printf("ERROR: this node is missing the `managed-by: karpenter` label. Aborting.\n")
			return nil
		}

		newTaints, shouldUpdate := taintsFilter.FilterUndesiredTaints(node.Spec.Taints)

		if shouldUpdate {
			fmt.Printf("removing capa taints from node %s\n", nodeName)
			fmt.Printf("old taints: %v\n", node.Spec.Taints)
			fmt.Printf("new taints: %v\n", newTaints)

			node.Spec.Taints = newTaints
			err = ctrlClient.Update(ctx, &node)
			if err != nil {
				fmt.Printf("ERROR: failed to save changes to taints: %v\n", err)
				return err
			}
			fmt.Printf("taints removed correctly\n")
		} else {
			fmt.Printf("no undesired taint was found\n")
		}

		return nil
	}

	attempts := 0
	for {
		err = patch()
		if err == nil {
			break
		}

		attempts++

		if attempts < retryAttempts {
			fmt.Printf("failed to patch node, retrying\n")
			continue
		}

		panic("failed to patch node, aborting\n")
	}

	fmt.Printf("capa-karpenter-taint-remover finished successfully\n")
	fmt.Printf("sleeping forever\n")

	select {
	// sleeping forever
	}
}
