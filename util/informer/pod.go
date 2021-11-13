package informer

import (
	"context"

	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

func EvictPodWithGracePerio(client clientset.Interface, pod v1.Pod, gracePeriodSeconds int64) error {
	return client.CoreV1().Pods(pod.Namespace).Evict(context.Background(), &policy.Eviction{
		ObjectMeta:    pod.ObjectMeta,
		DeleteOptions: metav1.NewDeleteOptions(gracePeriodSeconds),
	})
}
