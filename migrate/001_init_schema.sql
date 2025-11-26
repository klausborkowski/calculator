-- Создание таблицы package
CREATE TABLE IF NOT EXISTS package (
    id SERIAL PRIMARY KEY,
    size INTEGER NOT NULL
);

-- Инициализация с 3 записями (1, 2, 3)
INSERT INTO package (id, size) VALUES 
    (1, 1),
    (2, 2),
    (3, 3)
ON CONFLICT (id) DO NOTHING;

