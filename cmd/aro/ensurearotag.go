// +build aro

package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

func init() {
	// TODO: Replace ensurearotag.go with github.com/openshift/installer/pkg/types/azure when release-4.9-azure is available.
	/*
		This file is used to fake the functionality from https://github.com/openshift/installer/pull/4843/files
		while waiting for release-4.9-azure branch.

		This file will only be included when `-tags aro`.

		To revert, remove the following files:
		* ensurearotag.go
		* ensurearotagconst.go

		Keep:
		* ensurearotag_test.go

		The import: azuretypes "github.com/openshift/installer/pkg/types/azure"

		Add import to:
		* rp.go
		* ensurearotag_test.go
	*/
	platformIsAro = true
}
