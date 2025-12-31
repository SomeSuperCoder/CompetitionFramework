-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  name VARCHAR NOT NULL CHECK (name != ''),
  details TEXT NOT NULL CHECK (details != ''),
  points INT NOT NULL CHECK (points > 0),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
