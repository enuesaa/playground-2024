CREATE TABLE aaa (
    id BLOB PRIMARY KEY CHECK(is_uuid_v7(id)) DEFAULT (uuid_v7()) NOT NULL,
    new_1 TEXT DEFAULT ''
) STRICT;