# Copyright (c) 2026 UsedContract contributors
# SPDX-License-Identifier: BUSL-1.1
# Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

import json, urllib.request
class UsedContract:
    def __init__(self, collector, consumer, version="", sender=None):
        self.collector=collector.rstrip('/'); self.consumer=consumer; self.version=version; self.sender=sender or self._send
    def _send(self,event):
        data=json.dumps(event).encode(); req=urllib.request.Request(self.collector+'/v1/events',data=data,headers={'Content-Type':'application/json'},method='POST'); urllib.request.urlopen(req,timeout=2).read()
    def emit(self,method,route,path,kind,operator=None,operand=None,result=None,trace_id=None):
        e={'consumer':{'name':self.consumer,'version':self.version,'language':'python'},'endpoint':{'method':method,'route':route},'path':path,'kind':kind}
        if operator is not None:e['operator']=operator
        if operand is not None:e['operand']=operand
        if result is not None:e['result']=result
        if trace_id is not None:e['trace_id']=trace_id
        self.sender(e)
    def read(self,method,route,path): self.emit(method,route,path,'read')
    def compare(self,method,route,path,operator,operand,result): self.emit(method,route,path,'compare',operator,operand,result)
    def control(self,method,route,path,result): self.emit(method,route,path,'control',result=result)
