CREATE TABLE IF NOT EXISTS customers (
    id              BIGSERIAL    PRIMARY KEY,
    user_id         BIGINT       NOT NULL,
    email           VARCHAR(255),
    whatsapp_number VARCHAR(255),
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_customers_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_customers_user_id    ON customers (user_id);
CREATE INDEX IF NOT EXISTS idx_customers_email      ON customers (email);
CREATE INDEX IF NOT EXISTS idx_customers_deleted_at ON customers (deleted_at);
