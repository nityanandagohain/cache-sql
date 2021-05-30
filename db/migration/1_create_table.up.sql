CREATE TABLE cache (
    cache_key varchar(10) primary key,
    value varchar(10) not null,
    ttl int default -2
);