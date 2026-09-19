CREATE TABLE IF NOT EXISTS user_tasks(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    user_id BIGINT NOT NULL REFERENCES users(id),
    item_id BIGINT NOT NULL REFERENCES course_items(id),
    grade DECIMAL(10,5),
    progress SMALLINT NOT NULL DEFAULT 1, -- 1 == to-do
    target_start TIMESTAMP,
    actual_start TIMESTAMP,
    target_done TIMESTAMP,
    actual_done TIMESTAMP
);