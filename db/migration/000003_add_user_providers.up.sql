CREATE TABLE user_providers (
    user_id VARCHAR(25) NOT NULL,                     -- `users`テーブルの外部キー
    provider VARCHAR(50) NOT NULL,            -- OAuthプロバイダー名
    provider_user_id VARCHAR(255) NOT NULL,   -- プロバイダー内の一意なユーザーID
    access_token TEXT,                        -- アクセストークン
    refresh_token TEXT,                       -- リフレッシュトークン
    expires_at TIMESTAMP,                     -- トークンの有効期限
    created_at TIMESTAMP DEFAULT now(),       -- 作成日時
    updated_at TIMESTAMP DEFAULT now(),       -- 更新日時
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, provider),          -- 複合主キー
    UNIQUE (provider, provider_user_id)       -- プロバイダー内での一意性を保証
);