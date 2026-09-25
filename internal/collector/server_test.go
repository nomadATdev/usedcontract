// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package collector

import (
	"github.com/nomadATdev/usedcontract/internal/inference"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEventEndpoint(t *testing.T) {
	s := New(inference.NewStore())
	b := `{"consumer":{"name":"x"},"endpoint":{"method":"GET","route":"/a"},"path":"$.id","kind":"read"}`
	r := httptest.NewRequest("POST", "/v1/events", strings.NewReader(b))
	w := httptest.NewRecorder()
	s.Handler().ServeHTTP(w, r)
	if w.Code != 202 {
		t.Fatalf("%d %s", w.Code, w.Body.String())
	}
	if len(s.Store.List()) != 1 {
		t.Fatal("missing")
	}
}
