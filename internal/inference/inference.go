// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package inference

import (
	"fmt"
	"github.com/nomadATdev/usedcontract/internal/ir"
	"sort"
	"sync"
)

type Requirement struct {
	Consumer         ir.Consumer      `json:"consumer"`
	Endpoint         ir.Endpoint      `json:"endpoint"`
	Path             string           `json:"path"`
	Reads            int              `json:"reads"`
	Required         bool             `json:"required"`
	NonNull          bool             `json:"non_null"`
	Comparisons      map[string][]any `json:"comparisons,omitempty"`
	ControlInfluence bool             `json:"control_influence"`
}

type key struct{ consumer, version, method, route, path string }
type Store struct {
	mu  sync.RWMutex
	req map[key]*Requirement
}

func NewStore() *Store { return &Store{req: map[key]*Requirement{}} }
func (s *Store) Add(e ir.Event) {
	k := key{e.Consumer.Name, e.Consumer.Version, e.Endpoint.Method, e.Endpoint.Route, e.Path}
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.req[k]
	if r == nil {
		r = &Requirement{Consumer: e.Consumer, Endpoint: e.Endpoint, Path: e.Path, Comparisons: map[string][]any{}}
		s.req[k] = r
	}
	switch e.Kind {
	case "read":
		r.Reads++
		r.Required = true
	case "null_check":
		if e.Operator == "not_null" {
			r.NonNull = true
		}
	case "compare":
		r.Comparisons[e.Operator] = appendUnique(r.Comparisons[e.Operator], e.Operand)
	case "control":
		r.ControlInfluence = true
	}
}
func appendUnique(xs []any, v any) []any {
	sv := fmt.Sprint(v)
	for _, x := range xs {
		if fmt.Sprint(x) == sv {
			return xs
		}
	}
	return append(xs, v)
}
func (s *Store) List() []Requirement {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Requirement, 0, len(s.req))
	for _, r := range s.req {
		cp := *r
		out = append(out, cp)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Consumer.Name+out[i].Path < out[j].Consumer.Name+out[j].Path })
	return out
}
