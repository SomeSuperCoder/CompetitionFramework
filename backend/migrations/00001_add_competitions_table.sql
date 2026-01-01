-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE unit_status AS ENUM (
  'awaiting',
  'running',
  'completed'
);

CREATE TABLE competitions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL CHECK (name != ''),
  status unit_status NOT NULL DEFAULT 'awaiting',

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE competitions;
DROP TYPE unit_status;
-- +goose StatementEnd
