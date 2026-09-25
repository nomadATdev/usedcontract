// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package collector

import (
	"encoding/json"
	"github.com/nomadATdev/usedcontract/internal/inference"
	"github.com/nomadATdev/usedcontract/internal/ir"
	"net/http"
)

type Server struct {
	Store *inference.Store
	mux   *http.ServeMux
}

func New(s *inference.Store) *Server {
	x := &Server{Store: s, mux: http.NewServeMux()}
	x.mux.HandleFunc("POST /v1/events", x.events)
	x.mux.HandleFunc("GET /v1/contracts", x.contracts)
	x.mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
	return x
}
func (s *Server) Handler() http.Handler { return s.mux }
func (s *Server) events(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var e ir.Event
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&e); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	e.Normalize()
	if err := e.Validate(); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	s.Store.Add(e)
	w.WriteHeader(202)
}
func (s *Server) contracts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Store.List())
}
