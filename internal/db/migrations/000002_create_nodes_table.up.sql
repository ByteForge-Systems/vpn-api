CREATE TABLE IF NOT EXISTS nodes (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    port INTEGER NOT NULL,
    public_key TEXT,
    private_key TEXT,
    country VARCHAR(3) NOT NULL,
    comment TEXT NOT NULL DEFAULT '',
    is_online BOOLEAN NOT NULL,
    ssh_port INTEGER NOT NULL DEFAULT 22,
    ssh_username VARCHAR(100) NOT NULL,
    ssh_password TEXT
);
