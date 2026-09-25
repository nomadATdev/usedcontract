// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package ir

import "testing"

func TestNormalizeValidate(t *testing.T) {
	e := Event{Consumer: Consumer{Name: "billing"}, Endpoint: Endpoint{Method: "get", Route: "/v1/account"}, Path: "$.status", Kind: "read"}
	e.Normalize()
	if e.Endpoint.Method != "GET" {
		t.Fatal(e.Endpoint.Method)
	}
	if err := e.Validate(); err != nil {
		t.Fatal(err)
	}
}
func TestRejectBadPath(t *testing.T) {
	e := Event{Schema: SchemaVersion, Consumer: Consumer{Name: "x"}, Endpoint: Endpoint{Method: "GET", Route: "/"}, Path: "status", Kind: "read"}
	if e.Validate() == nil {
		t.Fatal("expected error")
	}
}
