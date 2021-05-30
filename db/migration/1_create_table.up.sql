CREATE TABLE cache (
    key varchar(10) primary key,
    value varchar not null,
    ttl int default(-2)
);

-- create index on cache (key);
