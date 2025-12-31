-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  competition UUID NOT NULL REFERENCES competitions(id),
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL CHECK (end_time > start_time),
  winner UUID REFERENCES users(id),
  user1 UUID NOT NULL REFERENCES users(id),
  user2 UUID REFERENCES users(id) CHECK (user1 != user2), -- can be null
  prev UUID REFERENCES matches(id) CHECK (prev != id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (competition, user1, user2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
