-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  competition UUID NOT NULL REFERENCES competitions(id),
  winner UUID REFERENCES users(id),
  user1 UUID NOT NULL REFERENCES users(id),
  user2 UUID REFERENCES users(id) CHECK (user1 != user2), -- can be null
  next UUID REFERENCES matches(id) CHECK (next != id),
  status unit_status NOT NULL DEFAULT 'awaiting',

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (competition, user1, user2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
