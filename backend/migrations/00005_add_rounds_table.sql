-- +goose Up
-- +goose StatementBegin
CREATE TABLE rounds (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  name VARCHAR NOT NULL CHECK (name != ''),
  match UUID NOT NULL REFERENCES matches(id),
  user1 UUID NOT NULL REFERENCES users(id),
  user2 UUID REFERENCES users(id) CHECK (user1 != user2), -- can be null
  winner UUID REFERENCES users(id) CHECK (winner IN (NULL, user1, user2)),
  prev UUID REFERENCES rounds(id) CHECK (prev != id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rounds;
-- +goose StatementEnd
