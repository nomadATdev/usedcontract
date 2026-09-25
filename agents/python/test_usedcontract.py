# Copyright (c) 2026 UsedContract contributors
# SPDX-License-Identifier: BUSL-1.1
# Change Date: 2030-09-25; Change License: Apache-2.0. See LICENSE.

import unittest
from usedcontract import UsedContract
class T(unittest.TestCase):
 def test_emit(self):
  x=[];u=UsedContract('http://x','svc',sender=x.append);u.read('GET','/a','$.id');self.assertEqual(x[0]['consumer']['language'],'python');self.assertEqual(x[0]['kind'],'read')
if __name__=='__main__':unittest.main()
