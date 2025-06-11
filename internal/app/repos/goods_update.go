package repos

import (
	"context"
	"test/internal/app/models"
)

const updateQuery = `
	UPDATE goods
	SET
		name = $3,
		description = COALESCE($4, description)
	WHERE id = $1 AND project_id = $2
	RETURNING
		id, 
		project_id, 
		name, 
		description, 
		priority, 
		removed, 
		created_at
`

func (r *Goods) Update(ctx context.Context, id, projectID int, name string, description *string) (*models.Goods, error) {
	var goods models.Goods
	err := r.db.QueryRowContext(ctx, updateQuery, id, projectID, name, description).
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
