INSERT INTO user_ (id, created_at, updated_at, version)
VALUES
    ('user_2ptpubMyPtlQy34bvW7z0ohvXW5', '2006-01-02', '2006-01-02', 1),
    ('user_2ptq4Ucfz0uhqfhUribcZyXWdQQ', '2006-01-02', '2006-01-02', 1),
    ('user_2ptq7WTFYGipSxKUdSmAXatZxdK', '2006-01-02', '2006-01-02', 1),
    ('user_2ptqA4h1eQoeSD7TI9fbL4apFEd', '2006-01-02', '2006-01-02', 1)
ON CONFLICT (id) DO NOTHING
;
