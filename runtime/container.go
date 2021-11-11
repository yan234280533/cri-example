package runtime

import (
	"fmt"
	goruntime "runtime"

	"golang.org/x/net/context"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"k8s.io/klog"
)

type UpdateOptions struct {
	// (Windows only) Number of CPUs available to the container.
	CPUCount int64
	// (Windows only) Portion of CPU cycles specified as a percentage * 100.
	CPUMaximum int64
	// CPU CFS (Completely Fair Scheduler) period. Default: 0 (not specified).
	CPUPeriod int64
	// CPU CFS (Completely Fair Scheduler) quota. Default: 0 (not specified).
	CPUQuota int64
	// CPU shares (relative weight vs. other containers). Default: 0 (not specified).
	CPUShares int64
	// Memory limit in bytes. Default: 0 (not specified).
	MemoryLimitInBytes int64
	// OOMScoreAdj adjusts the oom-killer score. Default: 0 (not specified).
	OomScoreAdj int64
	// CpusetCpus constrains the allowed set of logical CPUs. Default: "" (not specified).
	CpusetCpus string
	// CpusetMems constrains the allowed set of memory nodes. Default: "" (not specified).
	CpusetMems string
}

// UpdateContainerResources sends an UpdateContainerResourcesRequest to the server, and parses the returned UpdateContainerResourcesResponse.
func UpdateContainerResources(client pb.RuntimeServiceClient, containerId string, opts UpdateOptions) error {
	if containerId == "" {
		return fmt.Errorf("containerId cannot be empty")
	}
	request := &pb.UpdateContainerResourcesRequest{
		ContainerId: containerId,
	}
	if goruntime.GOOS != "windows" {
		request.Linux = &pb.LinuxContainerResources{
			CpuPeriod:          opts.CPUPeriod,
			CpuQuota:           opts.CPUQuota,
			CpuShares:          opts.CPUShares,
			CpusetCpus:         opts.CpusetCpus,
			CpusetMems:         opts.CpusetMems,
			MemoryLimitInBytes: opts.MemoryLimitInBytes,
			OomScoreAdj:        opts.OomScoreAdj,
		}
	} else {
		request.Windows = &pb.WindowsContainerResources{
			CpuCount:           opts.CPUCount,
			CpuMaximum:         opts.CPUMaximum,
			CpuShares:          opts.CPUShares,
			MemoryLimitInBytes: opts.MemoryLimitInBytes,
		}
	}

	klog.V(5).Infof("UpdateContainerResourcesRequest: %v", request)
	r, err := client.UpdateContainerResources(context.Background(), request)
	if err != nil {
		return err
	}

	klog.V(5).Infof("UpdateContainerResourcesResponse: %v", r)

	return nil
}

// RemoveContainer sends a RemoveContainerRequest to the server, and parses
// the returned RemoveContainerResponse.
func RemoveContainer(client pb.RuntimeServiceClient, ContainerId string) error {
	if ContainerId == "" {
		return fmt.Errorf("ID cannot be empty")
	}

	request := &pb.RemoveContainerRequest{
		ContainerId: ContainerId,
	}

	klog.V(5).Infof("RemoveContainerRequest: %v", request)

	r, err := client.RemoveContainer(context.Background(), request)
	if err != nil {
		return err
	}

	klog.V(5).Infof("RemoveContainerResponse: %v", r)
	return nil
}
