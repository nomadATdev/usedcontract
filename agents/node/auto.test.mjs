// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

import assert from "node:assert/strict"; import {trackJSON} from "./auto.mjs";
const e=[];const x=trackJSON({status:"ACTIVE",plan:{limit:5}},{consumer:"c",version:"1",method:"GET",route:"/a",emit:v=>e.push(v)});
assert.equal(x.status,"ACTIVE");assert.equal(x.plan.limit,5);assert.deepEqual(e.map(x=>x.path),["$.status","$.plan","$.plan.limit"]);