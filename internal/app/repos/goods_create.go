package repos

import (
	"context"
	"test/internal/app/models"
)

const createQuery = `
	INSERT INTO goods (project_id, name, priority, description) 
	VALUES (
		$1, 
		$2,
		COALESCE((SELECT MAX(priority) FROM goods), 0) + 1,
		''
	)
	RETURNING 
		id, 
		project_id, 
		name, 
		description, 
		priority, 
		removed, 
		created_at
`

func (r *Goods) Create(ctx context.Context, projectID int, name string) (*models.Goods, error) {
	var goods models.Goods
	err := r.db.QueryRowContext(ctx, createQuery, projectID, name).
		Scan(
			&goods.Id,
			&goods.ProjectId,
			&goods.Name,
			&goods.Description,
			&goods.Priority,
			&goods.Removed,
			&goods.CreatedAt,
		)
	return &goods, err
}
