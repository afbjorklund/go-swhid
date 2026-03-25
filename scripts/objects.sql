.load ./git
select lower(hex(oid)), git_object_type(type), size from objects order by oid;
