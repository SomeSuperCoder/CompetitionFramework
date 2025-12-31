-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE role AS ENUM (
  'participant',
  'admin'
);

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL,
  email CITEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  role role NOT NULL,
 
  competition UUID NOT NULL REFERENCES competitions(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  -- Basic email regex check
  CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE role;
-- +goose StatementEnd
