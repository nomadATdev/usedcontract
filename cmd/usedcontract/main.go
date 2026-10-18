// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nomadATdev/usedcontract/internal/collector"
	"github.com/nomadATdev/usedcontract/internal/compat"
	"github.com/nomadATdev/usedcontract/internal/inference"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "serve":
		serve()
	case "verify":
		verify()
	default:
		usage()
		os.Exit(2)
	}
}
func usage() {
	fmt.Fprintln(os.Stderr, "usedcontract serve [-addr :8080]\nusedcontract verify -contract contract.json -candidate response.json")
}
func serve() {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	addr := fs.String("addr", ":8080", "listen address")
	fs.Parse(os.Args[2:])
	s := collector.New(inference.NewStore())
	log.Printf("UsedContract collector listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, s.Handler()))
}
func verify() {
	fs := flag.NewFlagSet("verify", flag.ExitOnError)
	cf := fs.String("contract", "", "contract JSON")
	cand := fs.String("candidate", "", "candidate JSON")
	fs.Parse(os.Args[2:])
	if *cf == "" || *cand == "" {
		fs.Usage()
		os.Exit(2)
	}
	a, err := os.ReadFile(*cf)
	must(err)
	b, err := os.ReadFile(*cand)
	must(err)
	var req []inference.Requirement
	must(json.Unmarshal(a, &req))
	f, err := compat.Evaluate(req, b)
	must(err)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	must(enc.Encode(f))
	if len(f) > 0 {
		os.Exit(3)
	}
}
func must(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
