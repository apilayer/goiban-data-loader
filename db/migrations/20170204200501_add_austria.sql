
-- +goose Up
INSERT INTO DATA_SOURCE (id, name) VALUES (6, "AT");


-- +goose Down
DELETE FROM DATA_SOURCE where id = 6;

