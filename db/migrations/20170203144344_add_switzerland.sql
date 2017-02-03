
-- +goose Up
INSERT INTO DATA_SOURCE (id, name) VALUES (5, "CH");


-- +goose Down
DELETE FROM DATA_SOURCE where id = 5;

