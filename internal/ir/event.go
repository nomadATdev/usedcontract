// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package ir

import (
	"errors"
	"strings"
	"time"
)

const SchemaVersion = "usedcontract.event.v1"

type Consumer struct {
	Name     string `json:"name"`
	Version  string `json:"version,omitempty"`
	Language string `json:"language,omitempty"`
}

type Endpoint struct {
	Method string `json:"method"`
	Route  string `json:"route"`
}

type Source struct {
	File string `json:"file,omitempty"`
	Line int    `json:"line,omitempty"`
}

type Event struct {
	Schema    string         `json:"schema"`
	EventID   string         `json:"event_id,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	TraceID   string         `json:"trace_id,omitempty"`
	Consumer  Consumer       `json:"consumer"`
	Endpoint  Endpoint       `json:"endpoint"`
	Path      string         `json:"path"`
	Kind      string         `json:"kind"`
	Operator  string         `json:"operator,omitempty"`
	Operand   any            `json:"operand,omitempty"`
	Result    *bool          `json:"result,omitempty"`
	Source    Source         `json:"source,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
}

var allowedKinds = map[string]bool{
	"read": true, "compare": true, "null_check": true,
	"arithmetic": true, "control": true, "propagate": true,
}

func (e *Event) Normalize() {
	if e.Schema == "" {
		e.Schema = SchemaVersion
	}
	e.Endpoint.Method = strings.ToUpper(strings.TrimSpace(e.Endpoint.Method))
	e.Endpoint.Route = strings.TrimSpace(e.Endpoint.Route)
	e.Path = strings.TrimSpace(e.Path)
	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now().UTC()
	}
}

func (e Event) Validate() error {
	if e.Schema != SchemaVersion {
		return errors.New("unsupported event schema")
	}
	if e.Consumer.Name == "" {
		return errors.New("consumer.name is required")
	}
	if e.Endpoint.Method == "" || e.Endpoint.Route == "" {
		return errors.New("endpoint method and route are required")
	}
	if e.Path == "" || !strings.HasPrefix(e.Path, "$") {
		return errors.New("path must be a JSONPath-like value beginning with $")
	}
	if !allowedKinds[e.Kind] {
		return errors.New("unsupported event kind")
	}
	if e.Kind == "compare" && e.Operator == "" {
		return errors.New("compare event requires operator")
	}
	return nil
}
