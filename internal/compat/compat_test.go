// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package compat

import (
	"github.com/usedcontract/usedcontract/internal/inference"
	"github.com/usedcontract/usedcontract/internal/ir"
	"testing"
)

func TestSemanticBreak(t *testing.T) {
	r := inference.Requirement{Consumer: ir.Consumer{Name: "checkout"}, Path: "$.status", Required: true, ControlInfluence: true, Comparisons: map[string][]any{"eq": {"ACTIVE"}}}
	f, e := Evaluate([]inference.Requirement{r}, []byte(`{"status":"active"}`))
	if e != nil || len(f) != 1 {
		t.Fatalf("%v %v", f, e)
	}
}
func TestCompatible(t *testing.T) {
	r := inference.Requirement{Consumer: ir.Consumer{Name: "checkout"}, Path: "$.status", Required: true, ControlInfluence: true, Comparisons: map[string][]any{"eq": {"ACTIVE"}}}
	f, e := Evaluate([]inference.Requirement{r}, []byte(`{"status":"ACTIVE"}`))
	if e != nil || len(f) != 0 {
		t.Fatalf("%v %v", f, e)
	}
}
