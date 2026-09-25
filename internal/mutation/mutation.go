// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package mutation

import (
	"encoding/json"
	"fmt"
	"github.com/usedcontract/usedcontract/internal/inference"
	"sort"
	"strings"
)

type Case struct {
	ID       string          `json:"id"`
	Consumer string          `json:"consumer"`
	Path     string          `json:"path"`
	Kind     string          `json:"kind"`
	Body     json.RawMessage `json:"body"`
}

func Generate(reqs []inference.Requirement, baseline []byte) ([]Case, error) {
	var root any
	if err := json.Unmarshal(baseline, &root); err != nil { return nil, err }
	var out []Case
	for _, r := range reqs {
		if r.Required {
			if m, ok := clone(root); ok && deletePath(m, r.Path) {
				b, _ := json.Marshal(m)
				out = append(out, Case{fmt.Sprintf("%s:%s:missing", r.Consumer.Name,r.Path), r.Consumer.Name,r.Path,"missing",b})
			}
		}
		if r.NonNull {
			if m, ok := clone(root); ok && setPath(m,r.Path,nil) {
				b,_:=json.Marshal(m); out=append(out,Case{fmt.Sprintf("%s:%s:null",r.Consumer.Name,r.Path),r.Consumer.Name,r.Path,"null",b})
			}
		}
		for op, vals := range r.Comparisons {
			if op=="eq" && len(vals)>0 {
				if m,ok:=clone(root); ok && setPath(m,r.Path,"__USEDCONTRACT_UNSEEN__") {
					b,_:=json.Marshal(m); out=append(out,Case{fmt.Sprintf("%s:%s:unseen-equality",r.Consumer.Name,r.Path),r.Consumer.Name,r.Path,"unseen_equality",b})
				}
			}
		}
	}
	sort.Slice(out,func(i,j int)bool{return out[i].ID<out[j].ID})
	return out,nil
}
func clone(v any)(any,bool){ b,e:=json.Marshal(v);if e!=nil{return nil,false};var x any;if json.Unmarshal(b,&x)!=nil{return nil,false};return x,true}
func parts(path string)[]string{if !strings.HasPrefix(path,"$."){return nil};return strings.Split(strings.TrimPrefix(path,"$."),".")}
func deletePath(root any,path string)bool{ps:=parts(path);if len(ps)==0{return false};cur,ok:=root.(map[string]any);if !ok{return false};for _,x:=range ps[:len(ps)-1]{n,ok:=cur[x].(map[string]any);if !ok{return false};cur=n};if _,ok:=cur[ps[len(ps)-1]];!ok{return false};delete(cur,ps[len(ps)-1]);return true}
func setPath(root any,path string,v any)bool{ps:=parts(path);if len(ps)==0{return false};cur,ok:=root.(map[string]any);if !ok{return false};for _,x:=range ps[:len(ps)-1]{n,ok:=cur[x].(map[string]any);if !ok{return false};cur=n};if _,ok:=cur[ps[len(ps)-1]];!ok{return false};cur[ps[len(ps)-1]]=v;return true}
