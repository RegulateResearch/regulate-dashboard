-- for docker
SELECT 'CREATE DATABASE regulatedb'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'regulatedb')
\gexec
;

\c regulatedb;

-- table def
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    user_role SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS samples(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    data_str VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS records(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR NOT NULL,
    rand_num BIGINT NOT NULL,
    description VARCHAR
);