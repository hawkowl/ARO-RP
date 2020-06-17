package e2e

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"

	aro "github.com/Azure/ARO-RP/operator/apis/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/util/ready"
)

const (
	psMsgRecreate = "Re-creating the Pull Secret"
	psMsgUpdate   = "Updating the Pull Secret"
)

var pullSecretName = types.NamespacedName{Name: "pull-secret", Namespace: "openshift-config"}

func pullSecretExistsAndClusterReady(namespace string, name string) (done bool, err error) {
	_, err = clients.Kubernetes.CoreV1().Secrets(pullSecretName.Namespace).Get(pullSecretName.Name, metav1.GetOptions{})
	if err != nil && apierrors.IsNotFound(err) {
		return false, nil
	}

	// all nodes should be ready too (changing the pull secret cause nodes to re-start)
	nodes, err := clients.Kubernetes.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	for _, node := range nodes.Items {
		if !ready.NodeIsReady(&node) {
			log.Infof("node %s not ready", node.Name)
			return false, nil
		}
	}

	apiserver, err := clients.Operators.OperatorV1().KubeAPIServers().Get("cluster", metav1.GetOptions{})
	if err == nil {
		m := make(map[string]operatorv1.ConditionStatus, len(apiserver.Status.Conditions))
		for _, cond := range apiserver.Status.Conditions {
			m[cond.Type] = cond.Status
		}
		if m["Available"] == operatorv1.ConditionTrue && m["Progressing"] == operatorv1.ConditionFalse {
			return true, nil
		}
	}
	log.Info("apiserver not ready ", err)

	return false, nil
}

func operatorLogs() ([]byte, error) {
	var logs []byte
	err := retry.OnError(wait.Backoff{Steps: 5, Duration: 30 * time.Second, Factor: 1.5}, func(err error) bool {
		// just after changing the pull secret there are some transient errors as
		// the nodes are rotated
		log.Info("operatorLogs: ", err)
		return apierrors.IsTimeout(err) || apierrors.IsInvalid(err)
	}, func() error {
		pods, err := clients.Kubernetes.CoreV1().Pods("openshift-azure-operator").List(metav1.ListOptions{
			LabelSelector: "app=aro-operator-master",
		})
		if err != nil {
			return err
		}
		if len(pods.Items) != 1 {
			return apierrors.NewInvalid(
				pods.Items[0].GroupVersionKind().GroupKind(),
				"aro-operator-master",
				[]*field.Error{
					field.Duplicate(field.NewPath("aroPodCount"), len(pods.Items)),
				},
			)
		}
		logs, err = clients.Kubernetes.CoreV1().Pods("openshift-azure-operator").GetLogs(pods.Items[0].Name, &corev1.PodLogOptions{}).DoRaw()
		return err
	})
	return logs, err
}

func countMessages(msg []string) (map[string]int, error) {
	b, err := operatorLogs()
	if err != nil {
		return nil, err
	}

	result := map[string]int{}
	rx := regexp.MustCompile(fmt.Sprintf(`.*msg="(%s).*`, strings.Join(msg, "|")))
	changes := rx.FindAllStringSubmatch(string(b), -1)
	if len(changes) > 0 {
		for _, change := range changes {
			if len(change) == 2 {
				result[change[1]]++
			}
		}
	} else {
		log.Warnf("FindAllStringSubmatch: returned %v", changes)
	}
	return result, nil
}

func updatedObjects() ([]string, error) {
	b, err := operatorLogs()
	if err != nil {
		return nil, err
	}

	result := []string{}
	rx := regexp.MustCompile(`.*msg="(Update|Create) ([a-zA-Z\/.]+).*`)
	changes := rx.FindAllStringSubmatch(string(b), -1)
	if len(changes) > 0 {
		for _, change := range changes {
			if len(change) == 3 {
				result = append(result, change[1]+" "+change[2])
			}
		}
	} else {
		log.Warnf("FindAllStringSubmatch: returned %v", changes)
	}
	return result, nil
}

var _ = Describe("ARO Operator", func() {
	Specify("the InternetReachable default list should all be reachable", func() {
		co, err := clients.AROClusters.Clusters().Get("cluster", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(co.Status.Conditions.IsTrueFor(aro.InternetReachableFromMaster)).To(BeTrue())
		Expect(co.Status.Conditions.IsTrueFor(aro.InternetReachableFromWorker)).To(BeTrue())
	})
	Specify("custom invalid site shows not InternetReachable", func() {
		var originalSites []string
		// set an unreachable site
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			co, err := clients.AROClusters.Clusters().Get("cluster", metav1.GetOptions{})
			if err != nil {
				return err
			}
			originalSites = co.Spec.InternetChecker.Sites
			co.Spec.InternetChecker.Sites = []string{"https://localhost:1234/shouldnotexist"}
			_, err = clients.AROClusters.Clusters().Update(co)
			return err
		})
		Expect(err).NotTo(HaveOccurred())

		// confirm the conditions are correct
		timeoutCtx, cancel := context.WithTimeout(context.TODO(), 10*time.Minute)
		defer cancel()
		err = wait.PollImmediateUntil(time.Minute, func() (bool, error) {
			co, err := clients.AROClusters.Clusters().Get("cluster", metav1.GetOptions{})
			if err != nil {
				return false, err
			}
			log.Info(co.Status.Conditions)
			return co.Status.Conditions.IsFalseFor(aro.InternetReachableFromMaster) &&
				co.Status.Conditions.IsFalseFor(aro.InternetReachableFromWorker), nil
		}, timeoutCtx.Done())
		Expect(err).NotTo(HaveOccurred())

		// set the sites back again
		err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
			co, err := clients.AROClusters.Clusters().Get("cluster", metav1.GetOptions{})
			if err != nil {
				return err
			}
			co.Spec.InternetChecker.Sites = originalSites
			_, err = clients.AROClusters.Clusters().Update(co)
			return err
		})
		Expect(err).NotTo(HaveOccurred())
	})
	Specify("genevalogging must be repaired if deployment deleted", func() {
		mdsdReady := ready.CheckDaemonSetIsReadyRetryOnAllErrors(log, clients.Kubernetes.AppsV1().DaemonSets("openshift-azure-logging"), "mdsd")

		err := wait.PollImmediate(30*time.Second, 15*time.Minute, mdsdReady)
		Expect(err).NotTo(HaveOccurred())
		initial, err := updatedObjects()
		Expect(err).NotTo(HaveOccurred())

		// delete the mdsd daemonset
		err = clients.Kubernetes.AppsV1().DaemonSets("openshift-azure-logging").Delete("mdsd", nil)
		Expect(err).NotTo(HaveOccurred())

		// Wait for it to be fixed
		err = wait.PollImmediate(30*time.Second, 15*time.Minute, mdsdReady)
		Expect(err).NotTo(HaveOccurred())

		// confirm that only one object was updated
		final, err := updatedObjects()
		Expect(err).NotTo(HaveOccurred())
		if len(final)-len(initial) != 1 {
			log.Error("initial changes ", initial)
			log.Error("final changes ", final)
		}
		Expect(len(final) - len(initial)).To(Equal(1))
	})
	Specify("the pull secret should be re-added when deleted", func() {
		initial, err := countMessages([]string{psMsgRecreate, psMsgUpdate})
		Expect(err).NotTo(HaveOccurred())
		log.Info("initialCount ", initial)

		// Verify pull secret exists
		exists, err := pullSecretExistsAndClusterReady(pullSecretName.Namespace, pullSecretName.Name)
		Expect(err).NotTo(HaveOccurred())
		Expect(exists).To(BeTrue())

		// Delete pull secret
		err = clients.Kubernetes.CoreV1().Secrets(pullSecretName.Namespace).Delete(pullSecretName.Name, &metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())

		// give the cluster some time to start rotating the nodes
		time.Sleep(5 * time.Minute)

		// Wait for it to be fixed
		err = wait.PollImmediate(30*time.Second, 15*time.Minute, func() (bool, error) {
			newCount, err := countMessages([]string{psMsgRecreate, psMsgUpdate})
			log.Info("newCount ", newCount)
			return (newCount[psMsgRecreate]-initial[psMsgRecreate] == 1 && newCount[psMsgUpdate]-initial[psMsgUpdate] == 0), err
		})
		Expect(err).NotTo(HaveOccurred())

		// Wait for secret and nodes to settle
		err = wait.PollImmediate(30*time.Second, 15*time.Minute, func() (bool, error) {
			return pullSecretExistsAndClusterReady(pullSecretName.Namespace, pullSecretName.Name)
		})
		Expect(err).NotTo(HaveOccurred())
	})
})
