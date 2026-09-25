# Copyright (c) 2026 UsedContract contributors
# SPDX-License-Identifier: BUSL-1.1
# Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

import unittest
from auto_usedcontract import Context,loads
class T(unittest.TestCase):
 def test_auto_read_compare_control(self):
  ev=[]; c=Context("checkout","1","GET","/account",ev.append)
  x=loads('{"status":"ACTIVE","enabled":true}',c)
  self.assertTrue(x["status"]=="ACTIVE")
  self.assertTrue(x["enabled"])
  self.assertEqual([e["kind"] for e in ev],["read","compare","read","control"])
if __name__=="__main__":unittest.main()
