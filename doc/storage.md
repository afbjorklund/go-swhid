
### .swh directory

The .swh directory is similar to a .git directory.

It contains "objects" and "refs", and a "HEAD" file.

The HEAD contains the default revision, if any.
(otherwise a dummy value "ref: refs/heads/master")

The objects directory contain all the SWH objects,
compressed with zlib and indexed by two characters:

```
.swh/
├── HEAD
├── objects
│   ├── 00
│   │   └── 00000000000000000000000000000000000000
│   ├── ...
│   └── ff
│       └── ffffffffffffffffffffffffffffffffffffff
└── refs
```

It can be used as a bare git repository, `$GIT_DIR`.

To list all objects in the directory:

```sh
file -z .swh/objects/*/*
```

### .swh database

The .swh database file contains all the SWH objects:

```sql
CREATE TABLE objects (oid CHAR(20) PRIMARY KEY, type TEXT, size INT, data BLOB /*compressed*/);
CREATE TABLE refs (name TEXT PRIMARY KEY, oid CHAR(20), symbolic TEXT); -- either: oid|symbolic
```

The data is compressed with the `compress()` function.
It uses zlib compression, with a leading varint size.

The "type" and "size" are needed, with the uncompressed
"data", when calculating the checksum that is the "oid".

To list all objects in the database:

```sql
SELECT hex(oid), type, size FROM objects ORDER BY oid;
```
