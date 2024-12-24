CREATE TABLE tags (
    id CHAR(20) PRIMARY KEY,         -- 20文字の文字列を主キーとして使用
    name VARCHAR UNIQUE NOT NULL, -- タグ名（例: "絶景", "歴史", "初心者向け"）
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);