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
	// if the policy triggered restored action
	Restored bool
	// the influenced pod list
	// node detection the pod list is empty
	BeInfluencedPods []types.NamespacedName
}

type DetectionInterface interface {
	GetNodeDetectionCondition() []DetectionCondition
	GetPodDetectionCondition() []DetectionCondition
}

type AvoidanceInterface interface {
	AvoidanceActionMerge(podDetectConditions []DetectionCondition, nodeDetectConditions []DetectionCondition) (*AvoidanceActionStruct, error)
}

type AvoidanceActionStruct struct {
	BlockScheduledAction *BlockScheduledActionStruct
	ThrottleActions      []CPUThrottleActionStruct
	EvictActions         []EvictActionStruct
}

type BlockScheduledActionStruct struct {
	BlockScheduledQOSPriority   *ScheduledQOSPriority
	RestoreScheduledQOSPriority *ScheduledQOSPriority
}

type ScheduledQOSPriority struct {
	PodQOSClass        v1.PodQOSClass
	PriorityClassValue uint64
}

type CPUThrottleActionStruct struct {
	CPUDownAction *CPURatioStruct
	CPUUpAction   *CPURatioStruct
}

type CPURatioStruct struct {
	//the min of cpu ratio for pods
	// +optional
	MinCPURatio uint64 `json:"minCPURatio,omitempty"`

	//the step of cpu share and limit for once down-size (1-100)
	// +optional
	StepCPURatio uint64 `json:"stepCPURatio,omitempty"`
}

type MemoryThrottleActionStruct struct {
	// to force gc the page cache of low level pods
	// +optional
	ForceGC bool `json:"forceGC,omitempty"`
}

type ThrottleActionStruct struct {
	CPUThrottle    *CPUThrottleActionStruct
	MemoryThrottle *MemoryThrottleActionStruct
	ThrottlePods   []types.NamespacedName
}

type EvictActionStruct struct {
	DeletionGracePeriodSeconds *int32 `json:"deletionGracePeriodSeconds,omitempty"`
	EvictPods                  []types.NamespacedName
}

func AvoidanceActionsMerge(podDetectConditions []DetectionCondition, nodeDetectConditions []DetectionCondition) (*AvoidanceActionStruct, error) {
	//step1 do BlockScheduled merge
	//step2 do Retain merge FilterAndSortRestrainPods
	//step3 do Evict merge  FilterAndSortEvictPods

	return &AvoidanceActionStruct{}, nil
}
