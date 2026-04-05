CREATE TABLE IF NOT EXISTS iptv_accounts (
    id               BIGSERIAL    PRIMARY KEY,
    customer_id      BIGINT       NOT NULL,
    provider_id      BIGINT       NOT NULL,
    package_id       BIGINT       NOT NULL,
    provider_user_id VARCHAR(255),
    username         VARCHAR(255) NOT NULL,
    password         VARCHAR(255) NOT NULL,
    server_url       TEXT         NOT NULL,
    m3u_url          TEXT,
    max_connections  INT          NOT NULL DEFAULT 1,
    status           VARCHAR(50),
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMPTZ,

    CONSTRAINT fk_iptv_accounts_customer
        FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE,
    CONSTRAINT fk_iptv_accounts_provider
        FOREIGN KEY (provider_id) REFERENCES iptv_providers (id) ON DELETE CASCADE,
    CONSTRAINT fk_iptv_accounts_package
        FOREIGN KEY (package_id) REFERENCES provider_packages (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_iptv_accounts_customer_id ON iptv_accounts (customer_id);
CREATE INDEX IF NOT EXISTS idx_iptv_accounts_provider_id ON iptv_accounts (provider_id);
CREATE INDEX IF NOT EXISTS idx_iptv_accounts_package_id  ON iptv_accounts (package_id);
CREATE INDEX IF NOT EXISTS idx_iptv_accounts_deleted_at  ON iptv_accounts (deleted_at);
