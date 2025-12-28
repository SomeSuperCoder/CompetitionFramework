-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  name VARCHAR NOT NULL CHECK (name != ''),
  competition UUID NOT NULL REFERENCES competitions(id),
  start_time TIMESTAMP WITH TIME ZONE NOT NULL,
  end_time TIMESTAMP WITH TIME ZONE NOT NULL CHECK (end_time > start_time),
  winner UUID REFERENCES users(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
