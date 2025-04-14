CREATE TABLE touring_tags (
    id VARCHAR(20) PRIMARY KEY,  -- レコードの一意なID
    touring_id VARCHAR(20) NOT NULL,  -- 紐づくTouringのID
    tag_id VARCHAR(20) NOT NULL,      -- 紐づくTagのID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- レコード作成日時
    UNIQUE (touring_id, tag_id),     -- 同じTouringに同じTagを二重登録させない制約
    FOREIGN KEY (touring_id) REFERENCES tourings (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
)
