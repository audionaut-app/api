INSERT INTO operation_ (id, description, created_at, updated_at, version)
VALUES
    ('create', NULL, '2006-01-02', '2006-01-02', 1),
    ('read', NULL, '2006-01-02', '2006-01-02', 1),
    ('update', NULL, '2006-01-02', '2006-01-02', 1),
    ('delete', NULL, '2006-01-02', '2006-01-02', 1)
ON CONFLICT (id) DO NOTHING
;

INSERT INTO resource_ (id, description, created_at, updated_at, version)
VALUES
    ('tokens', NULL, '2006-01-02', '2006-01-02', 1)
ON CONFLICT (id) DO NOTHING
;

INSERT INTO context_ (id, description, created_at, updated_at, version)
VALUES
    ('global', NULL, '2006-01-02', '2006-01-02', 1),
    ('self', NULL, '2006-01-02', '2006-01-02', 1)
ON CONFLICT (id) DO NOTHING
;

INSERT INTO permission_ (operation_id, resource_id, context_id, description, active, created_at, updated_at, version)
VALUES
    ('create', 'tokens', 'global', NULL, TRUE, '2006-01-02', '2006-01-02', 1),
    ('create', 'tokens', 'self', NULL, TRUE, '2006-01-02', '2006-01-02', 1)
ON CONFLICT (operation_id, resource_id, context_id) DO NOTHING
;

INSERT INTO role_ (id, name, description, created_at, updated_at, version)
VALUES
    ('user:basic', 'Basic', NULL, '2006-01-02', '2006-01-02', 1),
    ('user:premium', 'Premium', NULL, '2006-01-02', '2006-01-02', 1),
    ('editor', 'Editor', NULL, '2006-01-02', '2006-01-02', 1),
    ('admin', 'Admin', NULL, '2006-01-02', '2006-01-02', 1)
ON CONFLICT (id) DO NOTHING
;

-- Assign permissions for user:basic.
-- INSERT INTO role_permission_xref_ (role_id, operation_id, resource_id, context_id, created_at, updated_at, version)
-- VALUES
--     ()
-- ON CONFLICT (role_id, operation_id, resource_id, context_id) DO NOTHING
-- ;

-- Assign permissions for user:premium.
INSERT INTO role_permission_xref_ (role_id, operation_id, resource_id, context_id, created_at, updated_at, version)
VALUES
    ('user:premium', 'create', 'tokens', 'self', '2006-01-02', '2006-01-02', 1)
ON CONFLICT (role_id, operation_id, resource_id, context_id) DO NOTHING
;

-- Assign permissions for editor.
-- INSERT INTO role_permission_xref_ (role_id, operation_id, resource_id, context_id, created_at, updated_at, version)
-- VALUES
--     ()
-- ON CONFLICT (role_id, operation_id, resource_id, context_id) DO NOTHING
-- ;

-- Assign permissions for admin.
INSERT INTO role_permission_xref_ (role_id, operation_id, resource_id, context_id, created_at, updated_at, version)
VALUES
    ('user:premium', 'create', 'tokens', 'global', '2006-01-02', '2006-01-02', 1)
ON CONFLICT (role_id, operation_id, resource_id, context_id) DO NOTHING
;

INSERT INTO user_role_xref_ (user_id, role_id, created_at, updated_at, version)
VALUES
    ('user_2ptpubMyPtlQy34bvW7z0ohvXW5', 'user:basic', '2006-01-02', '2006-01-02', 1),

    ('user_2ptq4Ucfz0uhqfhUribcZyXWdQQ', 'user:basic', '2006-01-02', '2006-01-02', 1),
    ('user_2ptq4Ucfz0uhqfhUribcZyXWdQQ', 'user:premium', '2006-01-02', '2006-01-02', 1),

    ('user_2ptq7WTFYGipSxKUdSmAXatZxdK', 'user:basic', '2006-01-02', '2006-01-02', 1),
    ('user_2ptq7WTFYGipSxKUdSmAXatZxdK', 'user:premium', '2006-01-02', '2006-01-02', 1),
    ('user_2ptq7WTFYGipSxKUdSmAXatZxdK', 'editor', '2006-01-02', '2006-01-02', 1),

    ('user_2ptqA4h1eQoeSD7TI9fbL4apFEd', 'user:basic', '2006-01-02', '2006-01-02', 1),
    ('user_2ptqA4h1eQoeSD7TI9fbL4apFEd', 'user:premium', '2006-01-02', '2006-01-02', 1),
    ('user_2ptqA4h1eQoeSD7TI9fbL4apFEd', 'editor', '2006-01-02', '2006-01-02', 1),
    ('user_2ptqA4h1eQoeSD7TI9fbL4apFEd', 'admin', '2006-01-02', '2006-01-02', 1)
ON CONFLICT (user_id, role_id) DO NOTHING
;
