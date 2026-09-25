// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
// Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

package evidence
import("crypto/sha256";"encoding/hex";"encoding/json";"time";"github.com/usedcontract/usedcontract/internal/inference")
type Report struct{Schema string `json:"schema"`;GeneratedAt time.Time `json:"generated_at"`;ContractSHA256 string `json:"contract_sha256"`;Requirements int `json:"requirements"`}
func New(reqs []inference.Requirement) Report {b,_:=json.Marshal(reqs);h:=sha256.Sum256(b);return Report{"usedcontract.evidence.v1",time.Now().UTC(),hex.EncodeToString(h[:]),len(reqs)}}
