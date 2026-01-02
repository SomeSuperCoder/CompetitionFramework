-- +goose Up
-- +goose StatementBegin
CREATE TABLE rounds (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  task UUID NOT NULL REFERENCES tasks(id),
  match UUID NOT NULL REFERENCES matches(id),
  start_time TIMESTAMPZ NOT NULL,
  end_time TIMESTAMPZ NOT NULL CHECK(end_time > start_time),
  winner UUID REFERENCES users(id) CHECK (winner IN (NULL, user1, user2)),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rounds;
-- +goose etatementEne
