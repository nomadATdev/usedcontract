// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package io.usedcontract;
import java.net.URI; import java.net.http.*; import java.nio.charset.StandardCharsets;
public final class Agent {
 private final String collector,consumer,version; private final HttpClient client;
 public Agent(String collector,String consumer,String version){this.collector=collector.replaceAll("/$","");this.consumer=consumer;this.version=version;this.client=HttpClient.newHttpClient();}
 private static String q(String s){return "\""+s.replace("\\","\\\\").replace("\"","\\\"")+"\"";}
 public void emit(String method,String route,String path,String kind,String operator,String operand,Boolean result)throws Exception{
  StringBuilder b=new StringBuilder("{\"consumer\":{\"name\":").append(q(consumer)).append(",\"version\":").append(q(version)).append(",\"language\":\"jvm\"},\"endpoint\":{\"method\":").append(q(method)).append(",\"route\":").append(q(route)).append("},\"path\":").append(q(path)).append(",\"kind\":").append(q(kind));
  if(operator!=null)b.append(",\"operator\":").append(q(operator)); if(operand!=null)b.append(",\"operand\":").append(q(operand)); if(result!=null)b.append(",\"result\":").append(result); b.append("}");
  HttpRequest r=HttpRequest.newBuilder(URI.create(collector+"/v1/events")).header("content-type","application/json").POST(HttpRequest.BodyPublishers.ofString(b.toString(),StandardCharsets.UTF_8)).build(); HttpResponse<Void> x=client.send(r,HttpResponse.BodyHandlers.discarding()); if(x.statusCode()/100!=2)throw new IllegalStateException("collector "+x.statusCode());
 }
 public void read(String m,String r,String p)throws Exception{emit(m,r,p,"read",null,null,null);} public void compare(String m,String r,String p,String op,String operand,boolean result)throws Exception{emit(m,r,p,"compare",op,operand,result);} public void control(String m,String r,String p,boolean result)throws Exception{emit(m,r,p,"control",null,null,result);}
}
