CREATE TABLE users (
    id VARCHAR(20) PRIMARY KEY,               -- Motify内で管理する一意ユーザーID (XID)
    google_sub VARCHAR(255) UNIQUE NOT NULL,  -- Googleのsub (一意識別子)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- レコード作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- レコード更新日時
);