-- +goose Up
-- +goose StatementBegin
CREATE TABLE inscriptions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  competition UUID NOT NULL REFERENCES competitions(id),
  participant UUID NOT NULL REFERENCES users(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (competition, participant)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inscriptions;
-- +goose StatementEnd
