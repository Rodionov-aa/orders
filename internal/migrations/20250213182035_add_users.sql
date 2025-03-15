-- +goose Up
INSERT INTO clients (inn, is_trash) VALUES
    ('1010', true);


-- +goose Down
