
The SWH objects can be written to a directory
or to a database, just like the Git objects.

Each object has an object id (an "oid") SHA,
which is the checksum of Type-Length-Value:

`<type><space><size><nul><data>`

A SHA-1 checksum is 20 bytes long.

Git repositories can be stored in git bundles,
which contains git packfiles, more efficiently
than the simple directory structure shown here.
But the loose objects are simpler to understand.

Git repositories can also be stored in libgit2
databases, which uses integer type and has data
which is not explicitly compressed (just BLOB).
There are different libgit2 backends for storage.

The integer values for types are:

0) `none` (not used by swh/git)
1) `commit`
2) `tree`
3) `blob`
4) `tag`
5) `snapshot` (not used by git)

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
