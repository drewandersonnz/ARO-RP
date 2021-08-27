// +build !aro

package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"testing"
)

func TestEnsureAroTag(t *testing.T) {
	/*
		This file should always fail this test when "+build !aro".
		The 'aro' tag is required for the openshift/installer to disable certain
		functionality which are valid for OpenShift on Azure, but not valid for ARO deployments.
	*/
	t.Fatalf("ARO-RP must be built, run, and tested with '-tags aro'")
}
