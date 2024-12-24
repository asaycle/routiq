CREATE TABLE user_profiles (
    id VARCHAR(20) PRIMARY KEY,               -- プロフィールID (XID)
    user_id VARCHAR(20) NOT NULL,             -- usersテーブルのユーザーID (外部キー)
    custom_id VARCHAR(50) UNIQUE,              -- ユーザーのエイリアス (一意制約)
    name VARCHAR(255),                        -- ユーザーの表示名
    bio TEXT,                                 -- 自己紹介文
    avatar_url VARCHAR(512),                  -- プロフィール画像のURL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- プロフィール作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- プロフィール更新日時
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);