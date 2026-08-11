CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    email VARCHAR UNIQUE,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR,
    display_name VARCHAR NOT NULL,
    user_role SMALLINT NOT NULL,
    civitas_id VARCHAR UNIQUE,
    has_sso_login BOOLEAN NOT NULL DEFAULT FALSE
);

ALTER TABLE users
ADD COLUMN IF NOT EXISTS
academic_role SMALLINT NOT NULL DEFAULT 1 -- 1 is student
;