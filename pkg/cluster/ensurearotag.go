// +build !aro

package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

func init() {
	/*
		This file should always panic when "+build !aro".
		The 'aro' tag is required for the openshift/installer to disable certain
		functionality which are valid for OpenShift on Azure, but not valid for ARO deployments.
	*/
	panic("ARO-RP must be built, run, and tested with '-tags aro'")
}
