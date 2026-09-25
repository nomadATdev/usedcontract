# Copyright (c) 2026 UsedContract contributors
# SPDX-License-Identifier: BUSL-1.1
# Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

"""Experimental automatic Python JSON-value observer.

Wrap decoded JSON with `track_json(...)`. Property/index reads and scalar
comparisons emit language-neutral UsedContract events without manual emit calls.
This is a feasibility adapter, not whole-program taint tracking.
"""
from dataclasses import dataclass
import json, time, uuid

@dataclass
class Context:
    consumer: str
    version: str
    method: str
    route: str
    emit: callable

def _event(ctx,path,kind,operator=None,operand=None,result=None):
    e={"schema":"usedcontract.event.v1","event_id":str(uuid.uuid4()),
       "timestamp":time.strftime("%Y-%m-%dT%H:%M:%SZ",time.gmtime()),
       "consumer":{"name":ctx.consumer,"version":ctx.version,"language":"python"},
       "endpoint":{"method":ctx.method,"route":ctx.route},"path":path,"kind":kind}
    if operator is not None:e["operator"]=operator
    if operand is not None:e["operand"]=operand
    if result is not None:e["result"]=bool(result)
    ctx.emit(e)

def _unwrap(x): return x.value if isinstance(x, TrackedScalar) else x

class TrackedScalar:
    def __init__(self,value,path,ctx): self.value,self.path,self.ctx=value,path,ctx
    def _cmp(self,op,other,fn):
        other=_unwrap(other); result=fn(self.value,other)
        _event(self.ctx,self.path,"compare",op,other,result); return result
    def __eq__(self,o): return self._cmp("eq",o,lambda a,b:a==b)
    def __ne__(self,o): return self._cmp("ne",o,lambda a,b:a!=b)
    def __lt__(self,o): return self._cmp("lt",o,lambda a,b:a<b)
    def __le__(self,o): return self._cmp("le",o,lambda a,b:a<=b)
    def __gt__(self,o): return self._cmp("gt",o,lambda a,b:a>b)
    def __ge__(self,o): return self._cmp("ge",o,lambda a,b:a>=b)
    def __bool__(self):
        r=bool(self.value); _event(self.ctx,self.path,"control",result=r); return r
    def __str__(self): return str(self.value)
    def __repr__(self): return repr(self.value)
    def unwrap(self): return self.value

class TrackedDict:
    def __init__(self,value,path,ctx): self.value,self.path,self.ctx=value,path,ctx
    def __getitem__(self,key):
        path=f"{self.path}.{key}"; _event(self.ctx,path,"read")
        return _wrap(self.value[key],path,self.ctx)
    def get(self,key,default=None):
        path=f"{self.path}.{key}"; _event(self.ctx,path,"read")
        if key not in self.value:return default
        return _wrap(self.value[key],path,self.ctx)

class TrackedList:
    def __init__(self,value,path,ctx): self.value,self.path,self.ctx=value,path,ctx
    def __getitem__(self,i):
        path=f"{self.path}[{i}]"; _event(self.ctx,path,"read")
        return _wrap(self.value[i],path,self.ctx)

def _wrap(v,path,ctx):
    if isinstance(v,dict): return TrackedDict(v,path,ctx)
    if isinstance(v,list): return TrackedList(v,path,ctx)
    return TrackedScalar(v,path,ctx)

def track_json(value,ctx): return _wrap(value,"$",ctx)
def loads(text,ctx): return track_json(json.loads(text),ctx)
