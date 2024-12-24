CREATE TABLE routes (
    id VARCHAR(20) PRIMARY KEY,         -- プライマリーキー
    name VARCHAR NOT NULL,              -- ルート名
    description TEXT,                   -- ルートの説明
    distance FLOAT,                     -- 距離（km単位、オプション）
    feature JSONB NOT NULL,
    user_id VARCHAR(25) NOT NULL,            -- 作成者（ユーザーID、外部キー）
    created_at TIMESTAMP DEFAULT now(), -- 作成日時
    updated_at TIMESTAMP DEFAULT now() -- 更新日時
    -- CONSTRAINT fk_created_by FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);