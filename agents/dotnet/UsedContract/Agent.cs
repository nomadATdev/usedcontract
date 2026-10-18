// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

using System.Net.Http.Json;
namespace UsedContract;
public sealed class Agent {
 private readonly HttpClient _http; private readonly string _consumer,_version;
 public Agent(HttpClient http,string consumer,string version=""){_http=http;_consumer=consumer;_version=version;}
 public async Task EmitAsync(string method,string route,string path,string kind,string? op=null,object? operand=null,bool? result=null,CancellationToken ct=default){
  var e=new {consumer=new{name=_consumer,version=_version,language="dotnet"},endpoint=new{method,route},path,kind,@operator=op,operand,result};
  using var r=await _http.PostAsJsonAsync("v1/events",e,ct);r.EnsureSuccessStatusCode();
 }
 public Task ReadAsync(string m,string r,string p,CancellationToken ct=default)=>EmitAsync(m,r,p,"read",ct:ct);
 public Task CompareAsync(string m,string r,string p,string op,object operand,bool result,CancellationToken ct=default)=>EmitAsync(m,r,p,"compare",op,operand,result,ct);
 public Task ControlAsync(string m,string r,string p,bool result,CancellationToken ct=default)=>EmitAsync(m,r,p,"control",result:result,ct:ct);
}
