// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v0alpha1 "github.com/grafana/grafana/pkg/generated/clientset/versioned/typed/alerting_notifications/v0alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNotificationsV0alpha1 struct {
	*testing.Fake
}

func (c *FakeNotificationsV0alpha1) TimeIntervals(namespace string) v0alpha1.TimeIntervalInterface {
	return &FakeTimeIntervals{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNotificationsV0alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
