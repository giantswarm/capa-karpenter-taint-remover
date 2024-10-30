[![CircleCI](https://circleci.com/gh/giantswarm/capa-karpenter-taint-remover.svg?&style=shield)](https://circleci.com/gh/giantswarm/capa-karpenter-taint-remover)

# capa-karpenter-taint-remover

CAPI adds the taint `node.cluster.x-k8s.io/uninitialized` to all the nodes, and only removes it if the node is in the `MachinePool` instance list, or it's a `Machine`.
Karpenter creates machines that CAPI does not know about so the taint is not removed automatically.

This is a small utility meant to run as a `DaemonSet` on all Karpenter nodes in a CAPA cluster to remove the CAPI taints from the nodes.
