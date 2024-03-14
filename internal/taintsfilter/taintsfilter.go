package taintsfilter

import (
	v1 "k8s.io/api/core/v1"
)

type Taints struct {
	unwantedTaints []string
}

func New(unwantedTaints []string) *Taints {
	return &Taints{
		unwantedTaints: unwantedTaints,
	}
}

func (t *Taints) FilterUndesiredTaints(taints []v1.Taint) ([]v1.Taint, bool) {
	var shouldUpdate bool
	var filteredTaints []v1.Taint
	for _, taint := range taints {
		if t.isUndesiredTaint(taint) {
			shouldUpdate = true
		} else {
			filteredTaints = append(filteredTaints, taint)
		}
	}

	return filteredTaints, shouldUpdate
}

func (t *Taints) isUndesiredTaint(taint v1.Taint) bool {
	for _, unwantedTaint := range t.unwantedTaints {
		if taint.Key == unwantedTaint {
			return true
		}
	}

	return false
}
