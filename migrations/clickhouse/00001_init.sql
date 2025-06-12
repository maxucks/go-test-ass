-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    Id INT,
    ProjectId INT,
    Name String,
    Description String,
    Priority INT,
    Removed Bool,
    EventTime DateTime
) 
ENGINE = MergeTree()
ORDER BY (Id, ProjectId, Name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
