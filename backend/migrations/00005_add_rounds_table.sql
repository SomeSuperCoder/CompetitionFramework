-- +goose Up
-- +goose StatementBegin
CREATE TABLE rounds (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  name VARCHAR NOT NULL CHECK (name != ''),
  match UUID NOT NULL REFERENCES matches(id),
  -- TODO: check if the winner is one of the match players
  winner UUID REFERENCES users(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rounds;
-- +goose StatementEnd
