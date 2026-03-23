.load ./sha1
.load ./compress
-- verify that the oid matches the checksum of the (uncompressed) data, with the git TLV header
select lower(hex(oid)), sha1(type || ' ' || length || char(0) || uncompress(data)) from objects;
