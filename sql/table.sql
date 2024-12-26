CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
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