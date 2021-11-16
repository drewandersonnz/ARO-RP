package namespace

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"strings"
)

// IsOpenShift returns true if ns is an openshift managed namespace.
func IsOpenShift(ns string) bool {
	return ns == "" ||
		ns == "default" ||
		ns == "openshift" ||
		strings.HasPrefix(ns, "kube-") ||
		strings.HasPrefix(ns, "openshift-")
}

// IsKnownCustomerNamespace returns true if ns is known to be used by customers
func IsKnownCustomerNamespace(ns string) bool {
	// The above IsOpenShift test accidentally captures these known customer namespaces
	return ns == "openshift-authentication" || // while important, doesn't affect cluster availability
		ns == "openshift-gitops" ||
		ns == "openshift-logging" || // logging stack is customer workload
		ns == "openshift-marketplace" ||
		ns == "openshift-operators" ||
		ns == "openshift-redhat-marketplace" ||
		ns == "Default" // customers choose to use this
}
