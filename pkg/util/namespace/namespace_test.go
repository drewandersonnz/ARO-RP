package namespace

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"testing"
)

func TestIsOpenShift(t *testing.T) {
	for _, tt := range []struct {
		namespace string
		want      bool
	}{
		{
			want: true,
		},
		{
			namespace: "openshift-ns",
			want:      true,
		},
		{
			namespace: "openshift",
			want:      true,
		},
		{
			namespace: "kube-ns",
			want:      true,
		},
		{
			namespace: "default",
			want:      true,
		},
		{
			namespace: "customer",
		},
	} {
		t.Run(tt.namespace, func(t *testing.T) {
			got := IsOpenShift(tt.namespace)
			if tt.want != got {
				t.Error(got)
			}
		})
	}
}

func TestIsKnownCustomerNamespace(t *testing.T) {
	for _, tt := range []struct {
		namespace string
		want      bool
	}{
		{
			// customers install their own authentication here, while important this is outside our SLA
			namespace: "openshift-authentication",
			want:      true,
		},
		{
			// customers choose to use the Default namespace, we don't use it, therefore this is considered customer workload
			namespace: "Default",
			want:      true,
		},
		{
			// openshift-kube-apiserver can affect SRE SLA
			namespace: "openshift-kube-apiserver",
			want:      false,
		},
	} {
		t.Run(tt.namespace, func(t *testing.T) {
			got := IsKnownCustomerNamespace(tt.namespace)
			if tt.want != got {
				t.Error(got)
			}
		})
	}
}
