Simple scripts for verifying directory or database

* "objects" - list objects

  `objects.sh .swh/objects/*/*`

  `objects.py .swh/objects/*/*`

  `sqlite3 swh.db <objects.sql`

* "checksum" - calculate checksum

  `checksum.sh .swh/objects/*/*`

  `checksum.py .swh/objects/*/*`

  `sqlite3 swh.db <checksum.sql`
