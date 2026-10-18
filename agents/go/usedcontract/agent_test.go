// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package usedcontract

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmit(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(202) }))
	defer s.Close()
	a := Agent{Collector: s.URL, Consumer: "go-svc"}
	if err := a.Emit("GET", "/a", "$.id", "read", "", nil, nil); err != nil {
		t.Fatal(err)
	}
}
