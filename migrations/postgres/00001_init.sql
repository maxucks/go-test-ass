-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS goods (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  project_id int REFERENCES projects(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  name varchar(100) NOT NULL,
  description text NOT NULL,
  priority int NOT NULL,
  removed boolean DEFAULT false,
  created_at timestamp NOT NULL DEFAULT now()
);

CREATE INDEX idx_goods_project_id ON goods(project_id);
CREATE INDEX idx_goods_name ON goods(name);

INSERT INTO projects (name) VALUES ('Первая запись');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goods;
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
