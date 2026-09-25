// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package inference

import (
	"github.com/usedcontract/usedcontract/internal/ir"
	"testing"
)

func TestInfer(t *testing.T) {
	s := NewStore()
	base := ir.Event{Consumer: ir.Consumer{Name: "web", Version: "1"}, Endpoint: ir.Endpoint{Method: "GET", Route: "/a"}, Path: "$.status"}
	e := base
	e.Kind = "read"
	s.Add(e)
	e = base
	e.Kind = "compare"
	e.Operator = "eq"
	e.Operand = "ACTIVE"
	s.Add(e)
	e = base
	e.Kind = "control"
	s.Add(e)
	r := s.List()[0]
	if !r.Required || !r.ControlInfluence || len(r.Comparisons["eq"]) != 1 {
		t.Fatalf("%+v", r)
	}
}
