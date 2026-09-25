// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package mutation
import("testing";"github.com/usedcontract/usedcontract/internal/inference";"github.com/usedcontract/usedcontract/internal/ir")
func TestGenerate(t *testing.T){s:=inference.NewStore(); b:=true
for _,e:=range []ir.Event{
{Consumer:ir.Consumer{Name:"c"},Endpoint:ir.Endpoint{Method:"GET",Route:"/a"},Path:"$.status",Kind:"read"},
{Consumer:ir.Consumer{Name:"c"},Endpoint:ir.Endpoint{Method:"GET",Route:"/a"},Path:"$.status",Kind:"compare",Operator:"eq",Operand:"ACTIVE"},
{Consumer:ir.Consumer{Name:"c"},Endpoint:ir.Endpoint{Method:"GET",Route:"/a"},Path:"$.status",Kind:"control",Result:&b},
}{e.Normalize();s.Add(e)}
cs,err:=Generate(s.List(),[]byte(`{"status":"ACTIVE","x":1}`));if err!=nil||len(cs)!=2{t.Fatalf("cases=%d err=%v",len(cs),err)}
}