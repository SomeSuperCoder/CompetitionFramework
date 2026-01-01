-- +goose Up
-- +goose StatementBegin
CREATE TABLE rounds (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  task UUID NOT NULL REFERENCES tasks(id),
  match UUID NOT NULL REFERENCES matches(id),
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL CHECK (end_time > start_time),
  winner UUID REFERENCES users(id),
  status unit_status NOT NULL DEFAULT 'awaiting',

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rounds;
-- +goose StatementEnd
