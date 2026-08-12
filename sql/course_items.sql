CREATE TABLE IF NOT EXISTS course_items(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    course_id BIGINT NOT NULL REFERENCES courses(id),
    name VARCHAR(100) NOT NULL,
    item_type SMALLINT NOT NULL,
    item_url VARCHAR,
    start_time TIMESTAMP,
    due_time TIMESTAMP,
    description VARCHAR
);