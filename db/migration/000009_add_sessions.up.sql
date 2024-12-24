CREATE TABLE sessions (
    id VARCHAR(20) PRIMARY KEY,           -- セッションID
    user_id VARCHAR(20) NOT NULL,         -- ユーザーID (usersテーブル参照)
    refresh_token VARCHAR(255) NOT NULL UNIQUE,   -- リフレッシュトークン
    expires_at TIMESTAMP NOT NULL,        -- セッション有効期限
    created_at TIMESTAMP DEFAULT NOW(),   -- セッション作成日時
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
