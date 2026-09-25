// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

export class UsedContract {
  constructor({collector, consumer, version="", fetchImpl=globalThis.fetch}) { this.collector=collector.replace(/\/$/,""); this.consumer=consumer; this.version=version; this.fetch=fetchImpl; }
  async emit({method,route,path,kind,operator,operand,result,trace_id}) { const e={consumer:{name:this.consumer,version:this.version,language:"node"},endpoint:{method,route},path,kind,operator,operand,result,trace_id}; const r=await this.fetch(`${this.collector}/v1/events`,{method:"POST",headers:{"content-type":"application/json"},body:JSON.stringify(e)}); if(!r.ok) throw new Error(`collector ${r.status}`); }
  read(ctx,path){return this.emit({...ctx,path,kind:"read"})}
  compare(ctx,path,operator,operand,result){return this.emit({...ctx,path,kind:"compare",operator,operand,result})}
  control(ctx,path,result){return this.emit({...ctx,path,kind:"control",result})}
}
