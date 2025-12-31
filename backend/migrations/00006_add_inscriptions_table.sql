-- +goose Up
-- +goose StatementBegin
CREATE TABLE inscriptions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  competition UUID NOT NULL REFERENCES competitions(id),
  participant UUID NOT NULL UNIQUE REFERENCES users(id),
  points INT NOT NULL DEFAULT 0 CHECK ( points >= 0 ),
  active BOOLEAN NOT NULL DEFAULT True,

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inscriptions;
-- +goose StatementEnd
