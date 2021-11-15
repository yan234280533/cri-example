package detect

import (
	"k8s.io/apimachinery/pkg/types"
)

type DetectionCondition struct {
	// the name of detection policy
	PolicyName string
	// the namespaces of pod policy
	Namespace string
	// if the policy triggered action
	Triggered bool
	// the influenced pod list, which is used to determine which pods should to do action
	// node detection the pod list is empty
	BeInfluencedPods []types.NamespacedName
}

type DetectionInterface interface {
	GetNodeDetectionCondition() []DetectionCondition
	GetPodDetectionCondition() []DetectionCondition
}
