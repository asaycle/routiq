CREATE TABLE ridings (
    id            VARCHAR(20) PRIMARY KEY, -- XID
    title         TEXT NOT NULL,           -- ツーリングタイトル
    note          TEXT,                    -- 詳細
    date          DATE NOT NULL,           -- ツーリングの日付
    score SMALLINT NOT NULL DEFAULT 0,
    route_id      VARCHAR(20) NOT NULL,    -- 紐付けるルート
    created_at    TIMESTAMP DEFAULT now(), -- 作成日時
    updated_at    TIMESTAMP DEFAULT now(), -- 更新日時
    FOREIGN KEY (route_id) REFERENCES routes(id) -- Routeとの関連付け
);