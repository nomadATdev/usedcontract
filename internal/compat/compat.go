// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package compat

import (
	"encoding/json"
	"fmt"
	"github.com/usedcontract/usedcontract/internal/inference"
	"strings"
)

type Finding struct {
	Consumer string `json:"consumer"`
	Path     string `json:"path"`
	Severity string `json:"severity"`
	Reason   string `json:"reason"`
}

func Evaluate(reqs []inference.Requirement, candidate []byte) ([]Finding, error) {
	var root any
	if err := json.Unmarshal(candidate, &root); err != nil {
		return nil, err
	}
	out := make([]Finding, 0)
	for _, r := range reqs {
		v, ok := lookup(root, r.Path)
		if !ok {
			if r.Required {
				out = append(out, Finding{r.Consumer.Name, r.Path, "breaking", "required observed field is missing"})
			}
			continue
		}
		if r.NonNull && v == nil {
			out = append(out, Finding{r.Consumer.Name, r.Path, "breaking", "observed non-null requirement now receives null"})
		}
		for op, vals := range r.Comparisons {
			if op == "eq" && r.ControlInfluence && !contains(vals, v) {
				out = append(out, Finding{r.Consumer.Name, r.Path, "breaking", fmt.Sprintf("candidate value %v was not among observed equality constants %v used in control flow", v, vals)})
			}
		}
	}
	return out, nil
}
func contains(xs []any, v any) bool {
	a := fmt.Sprint(v)
	for _, x := range xs {
		if fmt.Sprint(x) == a {
			return true
		}
	}
	return false
}
func lookup(root any, path string) (any, bool) {
	if path == "$" {
		return root, true
	}
	if !strings.HasPrefix(path, "$.") {
		return nil, false
	}
	cur := root
	for _, p := range strings.Split(strings.TrimPrefix(path, "$."), ".") {
		m, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = m[p]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}
