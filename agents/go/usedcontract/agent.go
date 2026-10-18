// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package usedcontract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Agent struct {
	Collector, Consumer, Version string
	Client                       *http.Client
}

func (a Agent) Emit(method, route, path, kind, operator string, operand any, result *bool) error {
	e := map[string]any{"consumer": map[string]string{"name": a.Consumer, "version": a.Version, "language": "go"}, "endpoint": map[string]string{"method": method, "route": route}, "path": path, "kind": kind}
	if operator != "" {
		e["operator"] = operator
	}
	if operand != nil {
		e["operand"] = operand
	}
	if result != nil {
		e["result"] = *result
	}
	b, _ := json.Marshal(e)
	c := a.Client
	if c == nil {
		c = http.DefaultClient
	}
	r, err := c.Post(strings.TrimRight(a.Collector, "/")+"/v1/events", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode/100 != 2 {
		return fmt.Errorf("collector status %d", r.StatusCode)
	}
	return nil
}
