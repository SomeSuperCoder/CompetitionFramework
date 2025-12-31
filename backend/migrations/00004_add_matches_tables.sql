-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  name VARCHAR NOT NULL CHECK (name != ''),
  competition UUID NOT NULL REFERENCES competitions(id),
  start_time TIMESTAMP WITH TIME ZONE NOT NULL,
  end_time TIMESTAMP WITH TIME ZONE NOT NULL CHECK (end_time > start_time),
  winner UUID REFERENCES users(id),
  user1 UUID NOT NULL REFERENCES users(id),
  user2 UUID REFERENCES users(id) CHECK (user1 != user2), -- can be null
  prev UUID REFERENCES matches(id) CHECK (prev != id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
