CREATE TABLE IF NOT EXISTS broadcasts (
    id         BIGSERIAL    PRIMARY KEY,
    subject    VARCHAR(255) NOT NULL,
    message    TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
