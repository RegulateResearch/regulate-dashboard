CREATE TABLE IF NOT EXISTS task_records(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    task_id BIGINT NOT NULL REFERENCES user_tasks(id),
    progress_before SMALLINT NOT NULL,
    progress_after SMALLINT NOT NULL,
    target_start_before TIMESTAMP,
    target_start_after TIMESTAMP,
    target_done_before TIMESTAMP,
    target_done_after TIMESTAMP
);