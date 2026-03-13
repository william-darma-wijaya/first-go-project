-- +goose Up
CREATE TABLE IF NOT EXISTS user_sickness (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    sickness_id TEXT NOT NULL,
    diagnosed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_sickness
        FOREIGN KEY (sickness_id)
        REFERENCES sickness(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_sickness;