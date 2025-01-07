CREATE TABLE riding_tags (
    id VARCHAR(20) PRIMARY KEY,  -- レコードの一意なID
    riding_id VARCHAR(20) NOT NULL,  -- 紐づくTouringのID
    tag_id VARCHAR(20) NOT NULL,      -- 紐づくTagのID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- レコード作成日時
    UNIQUE (riding_id, tag_id),     -- 同じRidingに同じTagを二重登録させない制約
    FOREIGN KEY (riding_id) REFERENCES ridings (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
)
