CREATE TABLE IF NOT EXISTS operation_ (
    id TEXT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT operation__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT operation__ak02 CHECK (version > 0),
    CONSTRAINT operation__ak03 CHECK (description IS NULL OR char_length(description) <= 1024)
);

CREATE TABLE IF NOT EXISTS resource_ (
    id TEXT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT resource__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT resource__ak02 CHECK (version > 0),
    CONSTRAINT resource__ak03 CHECK (description IS NULL OR char_length(description) <= 1024)
);

CREATE TABLE IF NOT EXISTS context_ (
    id TEXT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT context__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT context__ak02 CHECK (version > 0),
    CONSTRAINT context__ak03 CHECK (description IS NULL OR char_length(description) <= 1024)
);

CREATE TABLE IF NOT EXISTS permission_ (
    operation_id TEXT NOT NULL REFERENCES operation_ (id) ON DELETE CASCADE,
    resource_id TEXT NOT NULL REFERENCES resource_ (id) ON DELETE CASCADE,
    context_id TEXT NOT NULL REFERENCES context_ (id) ON DELETE CASCADE,
    description TEXT,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (operation_id, resource_id, context_id),
    CONSTRAINT permission__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT permission__ak02 CHECK (version > 0),
    CONSTRAINT permission__ak03 CHECK (description IS NULL OR char_length(description) <= 1024)
);

CREATE TABLE IF NOT EXISTS role_permission_xref_ (
    role_id TEXT,
    operation_id TEXT,
    resource_id TEXT,
    context_id TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (role_id, operation_id, resource_id, context_id),
    FOREIGN KEY (role_id) REFERENCES role_ (id) ON DELETE CASCADE,
    FOREIGN KEY (operation_id) REFERENCES operation_ (id) ON DELETE CASCADE,
    FOREIGN KEY (resource_id) REFERENCES resource_ (id) ON DELETE CASCADE,
    FOREIGN KEY (context_id) REFERENCES context_ (id) ON DELETE CASCADE,
    CONSTRAINT role_permission_xref__ak01 CHECK (updated_at >= created_at),
    CONSTRAINT role_permission_xref__ak02 CHECK (version > 0)
);
