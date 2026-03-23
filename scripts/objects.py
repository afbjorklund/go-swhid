#!/usr/bin/env python3

import sys
import zlib
import re


def objecttype(p):
    f = open(p, "rb")
    o = zlib.decompress(f.read())
    m = re.match(b"([a-z]+) ([0-9]+)\0(.*)", o)
    print("%s: Git %s %s (zlib compressed data)" % (p, m[1].decode(), m[2].decode()))


for arg in sys.argv[1:]:
    objecttype(arg)
