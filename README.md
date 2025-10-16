[![CircleCI](https://circleci.com/gh/giantswarm/capa-karpenter-taint-remover.svg?&style=shield)](https://circleci.com/gh/giantswarm/capa-karpenter-taint-remover)

# capa-karpenter-taint-remover

CAPA adds the taint `karpenter.sh/unregistered` sometimes due to [issue](https://github.com/kubernetes-sigs/cluster-api/issues/11401) after Karpenter has already removed it.

This is a small utility meant to run as a `DaemonSet` on all Karpenter nodes in a CAPA cluster to remove the Karpenter taints from the nodes.
