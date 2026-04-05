CREATE TABLE IF NOT EXISTS email_logs (
    id            BIGSERIAL   PRIMARY KEY,
    user_id       BIGINT      NOT NULL,
    type          VARCHAR(50),
    subject       TEXT,
    status        VARCHAR(50),
    error_message TEXT,
    sent_at       TIMESTAMPTZ,

    CONSTRAINT fk_email_logs_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_email_logs_user_id ON email_logs (user_id);
