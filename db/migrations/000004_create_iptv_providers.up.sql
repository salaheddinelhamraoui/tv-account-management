CREATE TABLE IF NOT EXISTS iptv_providers (
    id         BIGSERIAL    PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    api_url    TEXT         NOT NULL,
    api_token  TEXT         NOT NULL,
    type       VARCHAR(50)  NOT NULL DEFAULT 'xtream',
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_iptv_providers_deleted_at ON iptv_providers (deleted_at);
