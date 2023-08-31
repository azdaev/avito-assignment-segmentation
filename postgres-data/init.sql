CREATE TABLE IF NOT EXISTS segments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users_segments (
    user_id INT NOT NULL,
    segment_id INT NOT NULL,

    PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (segment_id) REFERENCES segments (id) ON DELETE CASCADE
);