CREATE TABLE IF NOT EXISTS subscriptions (
    id              BIGSERIAL   PRIMARY KEY,
    iptv_account_id BIGINT      NOT NULL,
    duration_months INT         NOT NULL,
    start_date      TIMESTAMPTZ NOT NULL,
    end_date        TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_subscriptions_iptv_account
        FOREIGN KEY (iptv_account_id) REFERENCES iptv_accounts (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_subscriptions_iptv_account_id ON subscriptions (iptv_account_id);
CREATE INDEX IF NOT EXISTS idx_subscriptions_end_date        ON subscriptions (end_date);
CREATE INDEX IF NOT EXISTS idx_subscriptions_deleted_at      ON subscriptions (deleted_at);
