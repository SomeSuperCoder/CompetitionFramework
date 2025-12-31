-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_orders (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  competition UUID NOT NULL REFERENCES competitions(id),
  task UUID NOT NULL REFERENCES tasks(id),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (competition, task)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_orders;
-- +goose StatementEnd
