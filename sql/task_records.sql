CREATE TABLE IF NOT EXISTS task_records(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    task_id BIGINT NOT NULL REFERENCES user_tasks(id),
    progress_before SMALLINT NOT NULL, -- 1 == to-do
    progress_after SMALLINT NOT NULL,
    target_start TIMESTAMP,
    actual_start TIMESTAMP,
    target_done TIMESTAMP,
    actual_done TIMESTAMP
);