package detect

import (
	v1 "k8s.io/api/core/v1"
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

type AvoidanceInterface interface {
	AvoidanceActionMerge(detectConditions []DetectionCondition) (*AvoidanceActionStruct, error)
}

type AvoidanceActionStruct struct {
	BlockScheduledAction *BlockScheduledActionStruct
	RetainActions        []RetainActionStruct
	EvictActions         []EvictActionStruct
}

type BlockScheduledActionStruct struct {
	PodQOSClass        v1.PodQOSClass
	PriorityClassValue uint64
}

type CPURestrainActionStruct struct {
	//the step of cpu share and limit for once down-size (1-100)
	// +optional
	StepCPURatio uint64 `json:"stepCPURatio,omitempty"`
}

type MemoryRestrainActionStruct struct {
	// to force gc the page cache of low level pods
	// +optional
	ForceGC bool `json:"forceGC,omitempty"`
}

type RetainActionStruct struct {
	CPURestrain    *CPURestrainActionStruct
	MemoryRestrain *MemoryRestrainActionStruct
	RetainPods     []types.NamespacedName
}

type EvictActionStruct struct {
	DeletionGracePeriodSeconds *int32 `json:"deletionGracePeriodSeconds,omitempty"`
	EvictPods                  []types.NamespacedName
}

func AvoidanceActionMerge(detectConditions []DetectionCondition) (*AvoidanceActionStruct, error) {
	//step1 do BlockScheduled merge
	//step2 do Retain merge
	//step3 do Evict merge  FilterAndSortEvictPods
}
