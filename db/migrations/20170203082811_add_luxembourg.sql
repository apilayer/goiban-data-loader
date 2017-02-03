
-- +goose Up
INSERT INTO DATA_SOURCE (id, name) VALUES (4, "LU");


-- +goose Down
DELETE FROM DATA_SOURCE where id = 4;

