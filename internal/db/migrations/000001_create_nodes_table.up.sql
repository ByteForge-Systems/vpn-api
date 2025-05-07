CREATE TABLE IF NOT EXISTS nodes (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    port INTEGER NOT NULL,
    ssh_port INTEGER NOT NULL DEFAULT 22,
    country VARCHAR(3) NOT NULL,
    comment TEXT NOT NULL DEFAULT '',
    is_online BOOLEAN NOT NULL,
    username VARCHAR(100) NOT NULL,
    password TEXT
);