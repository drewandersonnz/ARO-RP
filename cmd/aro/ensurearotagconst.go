package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// TODO: See ensurearotag.go

var platformIsAro bool

type Platform struct{}

func (o *Platform) IsARO() bool {
	return platformIsAro
}

type AzureTypes struct {
	Platform Platform
}

var azuretypes AzureTypes
