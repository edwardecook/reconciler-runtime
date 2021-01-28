/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package reconcilers

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/vmware-labs/reconciler-runtime/tracker"
)

func EnqueueTracked(gvk schema.GroupVersionKind, t tracker.Tracker) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(
		func(a client.Object) []reconcile.Request {
			var requests []reconcile.Request

			key := tracker.NewKey(
				gvk,
				types.NamespacedName{Namespace: a.GetNamespace(), Name: a.GetName()},
			)
			for _, item := range t.Lookup(key) {
				requests = append(requests, reconcile.Request{NamespacedName: item})
			}

			return requests
		},
	)
}
