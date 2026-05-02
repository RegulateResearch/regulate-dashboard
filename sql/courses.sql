CREATE TABLE IF NOT EXISTS courses(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    course_year SMALLINT NOT NULL,
    semester SMALLINT NOT NULL
);