// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package evidence

import (
	"github.com/nomadATdev/usedcontract/internal/inference"
	"testing"
)

func TestStableHash(t *testing.T) {
	r := []inference.Requirement{{Path: "$.x"}}
	a := New(r)
	b := New(r)
	if a.ContractSHA256 != b.ContractSHA256 || a.Requirements != 1 {
		t.Fatal("bad evidence")
	}
}
