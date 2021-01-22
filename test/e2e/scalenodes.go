package e2e

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"

	"github.com/Azure/ARO-RP/pkg/util/ready"
)

var _ = Describe("Scale nodes", func() {
	// hack: do this before we scale down, because it takes a while for the
	// nodes to settle after scale down
	Specify("node count should match the cluster resource and nodes should be ready", func() {
		ctx := context.Background()

		oc, err := clients.OpenshiftClustersv20200430.Get(ctx, clusterEnv.ResourceGroup(), clusterName)
		Expect(err).NotTo(HaveOccurred())

		expectedNodeCount := 3 // for masters
		for _, wp := range *oc.WorkerProfiles {
			expectedNodeCount += int(*wp.Count)
		}

		// another hack: we don't currently instantaneously expect all nodes to
		// be ready, it could be that the workaround operator is busy rotating
		// them, which we don't currently wait for on create
		err = wait.PollImmediate(10*time.Second, 30*time.Minute, func() (bool, error) {
			nodes, err := clients.Kubernetes.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
			if err != nil {
				log.Warn(err)
				return false, nil // swallow error
			}

			var nodeCount int
			for _, node := range nodes.Items {
				if ready.NodeIsReady(&node) {
					nodeCount++
				} else {
					for _, c := range node.Status.Conditions {
						log.Warnf("node %s status %s", node.Name, c.String())
					}
				}
			}

			return nodeCount == expectedNodeCount, nil
		})
		Expect(err).NotTo(HaveOccurred())
	})

	Specify("nodes should scale up and down", func() {
		mss, err := clients.MachineAPI.MachineV1beta1().MachineSets("openshift-machine-api").List(context.Background(), metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(mss.Items).NotTo(BeEmpty())

		err = scale(mss.Items[0].Name, 1)
		Expect(err).NotTo(HaveOccurred())

		err = waitForScale(mss.Items[0].Name)
		Expect(err).NotTo(HaveOccurred())

		err = scale(mss.Items[0].Name, -1)
		Expect(err).NotTo(HaveOccurred())

		err = waitForScale(mss.Items[0].Name)
		Expect(err).NotTo(HaveOccurred())
	})
})

func scale(name string, delta int32) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		ms, err := clients.MachineAPI.MachineV1beta1().MachineSets("openshift-machine-api").Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		if ms.Spec.Replicas == nil {
			ms.Spec.Replicas = to.Int32Ptr(1)
		}
		*ms.Spec.Replicas += delta

		_, err = clients.MachineAPI.MachineV1beta1().MachineSets(ms.Namespace).Update(context.Background(), ms, metav1.UpdateOptions{})
		return err
	})
}

func waitForScale(name string) error {
	return wait.PollImmediate(10*time.Second, 30*time.Minute, func() (bool, error) {
		ms, err := clients.MachineAPI.MachineV1beta1().MachineSets("openshift-machine-api").Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			log.Warn(err)
			return false, nil // swallow error
		}

		if ms.Spec.Replicas == nil {
			ms.Spec.Replicas = to.Int32Ptr(1)
		}

		return ms.Status.ObservedGeneration == ms.Generation &&
			ms.Status.AvailableReplicas == *ms.Spec.Replicas, nil
	})
}
