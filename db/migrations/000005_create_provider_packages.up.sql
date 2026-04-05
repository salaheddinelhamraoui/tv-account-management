CREATE TABLE IF NOT EXISTS provider_packages (
    id                  BIGSERIAL    PRIMARY KEY,
    provider_id         BIGINT       NOT NULL,
    display_name        VARCHAR(255) NOT NULL,
    external_package_id VARCHAR(255) NOT NULL,
    duration_months     INT          NOT NULL,
    created_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_provider_packages_provider
        FOREIGN KEY (provider_id) REFERENCES iptv_providers (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_provider_packages_provider_id ON provider_packages (provider_id);
CREATE INDEX IF NOT EXISTS idx_provider_packages_deleted_at  ON provider_packages (deleted_at);
