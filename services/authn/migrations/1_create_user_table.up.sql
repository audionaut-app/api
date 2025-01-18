CREATE DOMAIN user_id AS TEXT CHECK (VALUE ~ '^user_[0-9A-Za-z]{27}$');

CREATE TABLE IF NOT EXISTS user_ (
    id user_id PRIMARY KEY,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT user__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT user__ak02 CHECK (version > 0)
);
