// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

import assert from "node:assert/strict"; import {UsedContract} from "./index.mjs";
const calls=[]; const u=new UsedContract({collector:"http://x",consumer:"web",fetchImpl:async(u,o)=>{calls.push(JSON.parse(o.body));return {ok:true}}}); await u.read({method:"GET",route:"/a"},"$.status"); assert.equal(calls[0].consumer.language,"node"); assert.equal(calls[0].kind,"read");
