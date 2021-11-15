package informer

import (
	"fmt"

	"golang.org/x/net/context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// updateNodeConditions be used to update node condition
func updateNodeConditions(node *v1.Node, condition v1.NodeCondition) (*v1.Node, error) {
	if node == nil {
		return nil, fmt.Errorf("updateNodeCondition node is empty")
	}

	updatedNode := node.DeepCopy()

	conditions, err := GetNodeCondition(updatedNode)
	if err != nil {
		return nil, err
	}

	var bFound = false
	for i, cond := range conditions {
		if cond.Type == condition.Type {
			conditions[i] = condition
			bFound = true
		}
	}

	if !bFound {
		conditions = append(conditions, condition)
	}

	updatedNode.Status.Conditions = conditions

	return updatedNode, nil
}

// updateNodeTaints be used to update node taint
func updateNodeTaints(node *v1.Node, condition v1.Taint) (*v1.Node, error) {
	return nil, nil
}

func GetNodeCondition(node *v1.Node) ([]v1.NodeCondition, error) {

	if node == nil {
		return []v1.NodeCondition{}, fmt.Errorf("node resource is empty")
	}

	return node.Status.Conditions, nil
}

func FilterNodeConditionByType(conditions []v1.NodeCondition, conditionType string) (v1.NodeCondition, error) {
	for _, cond := range conditions {
		if string(cond.Type) == conditionType {
			return cond, nil
		}
	}

	return v1.NodeCondition{}, fmt.Errorf("condition %s is not found", conditionType)
}

// updateNodeStatus be used to update node status by communication with api-server
func updateNodeStatus(client clientset.Interface, updateNode *v1.Node) error {
	for i := 0; i < 3; i++ {
		_, err := client.CoreV1().Nodes().UpdateStatus(context.Background(), updateNode, metav1.UpdateOptions{})
		if err != nil {
			if errors.IsConflict(err) {
				continue
			} else {
				return err
			}
		}

		return nil

	}

	return fmt.Errorf("update node status failed, conflict too more times")
}
