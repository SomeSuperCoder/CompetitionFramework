-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL CHECK (name != ''),
  details TEXT NOT NULL CHECK (details != ''),
  competition UUID NOT NULL REFERENCES competitions(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
