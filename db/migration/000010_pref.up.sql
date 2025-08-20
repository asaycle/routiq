CREATE TABLE prefs (
    id CHAR(6) PRIMARY KEY,           -- 都道府県ID
    display_name VARCHAR(100) NOT NULL,   -- 表示名 (例: "東京都")
    created_at TIMESTAMP DEFAULT NOW(),   -- 作成日時
    updated_at TIMESTAMP DEFAULT NOW()  -- 更新日時
)