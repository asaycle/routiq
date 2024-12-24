CREATE TABLE route_tag_counts (
    id VARCHAR(20) PRIMARY KEY,
    route_id VARCHAR(20) NOT NULL,
    tag_id VARCHAR(20) NOT NULL,
    count int DEFAULT 0,
    UNIQUE (route_id, tag_id),
    FOREIGN KEY (route_id) REFERENCES routes (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
)
