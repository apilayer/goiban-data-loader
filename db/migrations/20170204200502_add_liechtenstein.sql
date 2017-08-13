
-- +goose Up
INSERT INTO DATA_SOURCE (id, name) VALUES (7, "LI");


-- +goose Down
DELETE FROM DATA_SOURCE where id = 7;

