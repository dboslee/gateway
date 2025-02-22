// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package status

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
)

func SetEnvoyPatchPolicyCondition(e *egv1a1.EnvoyPatchPolicy, conditionType gwv1a2.PolicyConditionType, status metav1.ConditionStatus, reason gwv1a2.PolicyConditionReason, message string) {
	cond := newCondition(string(conditionType), status, string(reason), message, time.Now(), e.Generation)
	e.Status.Conditions = MergeConditions(e.Status.Conditions, cond)
}

func SetEnvoyPatchPolicyProgrammedIfUnset(s *egv1a1.EnvoyPatchPolicyStatus, message string) {
	// Return early if Programmed condition is already set
	for _, c := range s.Conditions {
		if c.Type == string(egv1a1.PolicyConditionProgrammed) {
			return
		}
		if c.Type == string(gwv1a2.PolicyConditionAccepted) && c.Status == metav1.ConditionFalse {
			return
		}
	}

	cond := newCondition(string(egv1a1.PolicyConditionProgrammed), metav1.ConditionTrue, string(egv1a1.PolicyReasonProgrammed), message, time.Now(), 0)
	s.Conditions = MergeConditions(s.Conditions, cond)
}

func SetEnvoyPatchPolicyInvalid(s *egv1a1.EnvoyPatchPolicyStatus, message string) {
	cond := newCondition(string(egv1a1.PolicyConditionProgrammed), metav1.ConditionFalse, string(egv1a1.PolicyReasonInvalid), message, time.Now(), 0)
	s.Conditions = MergeConditions(s.Conditions, cond)
}

func SetEnvoyPatchPolicyResourceNotFound(s *egv1a1.EnvoyPatchPolicyStatus, message string) {
	cond := newCondition(string(egv1a1.PolicyConditionProgrammed), metav1.ConditionFalse, string(egv1a1.PolicyReasonResourceNotFound), message, time.Now(), 0)
	s.Conditions = MergeConditions(s.Conditions, cond)
}
