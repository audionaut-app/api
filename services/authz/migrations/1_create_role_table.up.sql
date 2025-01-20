CREATE DOMAIN user_id AS TEXT CHECK (VALUE ~ '^user_[0-9A-Za-z]{27}$');

CREATE TABLE IF NOT EXISTS role_ (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT role__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT role__ak02 CHECK (version > 0),
    CONSTRAINT role__ak03 CHECK (char_length(name) <= 256),
    CONSTRAINT role__ak04 CHECK (description IS NULL OR char_length(description) <= 1024)
);

CREATE TABLE IF NOT EXISTS user_role_xref_ (
    user_id user_id,
    role_id TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (role_id) REFERENCES role_ (id) ON DELETE CASCADE,
    CONSTRAINT user_role_xref__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT user_role_xref__ak02 CHECK (version > 0)
);
