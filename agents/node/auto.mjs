// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

// Experimental automatic field-consumption observer for Node.
// It transparently observes property reads while returning original primitive values.
// JavaScript strict equality on primitives cannot be intercepted by Proxy; semantic
// comparison tracking therefore requires build-time instrumentation in a later phase.
export function trackJSON(value, ctx, path="$") {
  if (value === null || typeof value !== "object") return value;
  return new Proxy(value, {
    get(target, prop, receiver) {
      const v = Reflect.get(target, prop, receiver);
      if (typeof prop === "symbol") return v;
      const child = Array.isArray(target) ? `${path}[${String(prop)}]` : `${path}.${String(prop)}`;
      ctx.emit({schema:"usedcontract.event.v1", consumer:{name:ctx.consumer,version:ctx.version,language:"node"},
        endpoint:{method:ctx.method,route:ctx.route}, path:child, kind:"read"});
      return trackJSON(v, ctx, child);
    }
  });
}
