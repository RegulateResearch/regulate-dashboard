-- for docker
\c frascati

-- table def
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    user_role SMALLINT NOT NULL
);

CREATE TABLE samples(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    data_str VARCHAR NOT NULL
);

CREATE TABLE records(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR NOT NULL,
    rand_num BIGINT NOT NULL,
    description VARCHAR
);