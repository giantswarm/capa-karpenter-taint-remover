package taintsfilter

import (
	"reflect"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestTaints_isUndesiredTaint(t1 *testing.T) {
	tests := []struct {
		name           string
		unwantedTaints []string
		taint          v1.Taint
		want           bool
	}{
		{
			name:           "No unwanted taints",
			taint:          v1.Taint{Key: "ANY"},
			unwantedTaints: make([]string, 0),
			want:           false,
		},
		{
			name:           "Unwanted taint",
			taint:          v1.Taint{Key: "unwanted"},
			unwantedTaints: []string{"unwanted"},
			want:           true,
		},
		{
			name:           "substring",
			taint:          v1.Taint{Key: "unwanted.io/unwanted"},
			unwantedTaints: []string{"unwanted"},
			want:           false,
		},
		{
			name:           "superstring",
			taint:          v1.Taint{Key: "unwanted"},
			unwantedTaints: []string{"unwanted.io/unwanted"},
			want:           false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Taints{
				unwantedTaints: tt.unwantedTaints,
			}
			if got := t.isUndesiredTaint(tt.taint); got != tt.want {
				t1.Errorf("isUndesiredTaint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaints_FilterUndesiredTaints(t1 *testing.T) {
	now := metav1.Time{Time: time.Now()}

	tests := []struct {
		name           string
		unwantedTaints []string
		taints         []v1.Taint
		filteredTaints []v1.Taint
		needsUpdate    bool
	}{
		{
			name:           "Nil unwanted taints, Nil taints",
			unwantedTaints: nil,
			taints:         nil,
			filteredTaints: nil,
			needsUpdate:    false,
		},
		{
			name:           "Empty unwanted taints, Empty taints",
			unwantedTaints: make([]string, 0),
			taints:         make([]v1.Taint, 0),
			filteredTaints: make([]v1.Taint, 0),
			needsUpdate:    false,
		},
		{
			name:           "Unwanted taints present, Empty taints",
			unwantedTaints: []string{"unwanted"},
			taints:         make([]v1.Taint, 0),
			filteredTaints: make([]v1.Taint, 0),
			needsUpdate:    false,
		},
		{
			name:           "Empty unwanted taints, one taint present",
			unwantedTaints: make([]string, 0),
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: false,
		},
		{
			name:           "Empty unwanted taints, many taints present",
			unwantedTaints: make([]string, 0),
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: false,
		},
		{
			name:           "Single unwanted taint, taints present but not unwanted",
			unwantedTaints: []string{"unwanted"},
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: false,
		},
		{
			name:           "Multiple unwanted taints, taints present but not unwanted",
			unwantedTaints: []string{"unwanted", "unwanted2"},
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "test2",
					Value:     "test2",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: false,
		},
		{
			name:           "Multiple unwanted taints, one unwanted taint",
			unwantedTaints: []string{"unwanted", "unwanted2"},
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "unwanted",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: true,
		},
		{
			name:           "Multiple unwanted taints, all taints unwanted",
			unwantedTaints: []string{"unwanted", "unwanted2"},
			taints: []v1.Taint{
				{
					Key:       "unwanted",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "unwanted2",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: make([]v1.Taint, 0),
			needsUpdate:    true,
		},
		{
			name:           "Multiple unwanted taints, some taints unwanted",
			unwantedTaints: []string{"unwanted", "unwanted2"},
			taints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "unwanted",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
				{
					Key:       "unwanted2",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			filteredTaints: []v1.Taint{
				{
					Key:       "test",
					Value:     "test",
					Effect:    "NoSchedule",
					TimeAdded: &now,
				},
			},
			needsUpdate: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Taints{
				unwantedTaints: tt.unwantedTaints,
			}
			filteredTaints, needsUpdate := t.FilterUndesiredTaints(tt.taints)
			if !isSameSlice(filteredTaints, tt.filteredTaints) {
				t1.Errorf("FilteredTaints got = %v, want %v", filteredTaints, tt.filteredTaints)
			}
			if needsUpdate != tt.needsUpdate {
				t1.Errorf("needsUpdate got = %v, want %v", needsUpdate, tt.needsUpdate)
			}
		})
	}
}

func isSameSlice(one []v1.Taint, two []v1.Taint) bool {
	if len(one) != len(two) {
		return false
	}

	// To avoid issues with nil slices.
	if len(one) == 0 {
		return true
	}

	return reflect.DeepEqual(one, two)
}
