#!/usr/bin/env python3

import sys
import zlib
import hashlib


def checksum(p):
    f = open(p, "rb")
    o = zlib.decompress(f.read())
    s = hashlib.sha1(o).hexdigest()
    print("%s  %s" % (s, p))


for arg in sys.argv[1:]:
    checksum(arg)
