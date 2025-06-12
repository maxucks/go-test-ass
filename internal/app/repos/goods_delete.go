package repos

import (
	"context"
	"test/internal/app/models"
)

const removeQuery = `
	UPDATE goods
	SET
		removed = true
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

func (r *Goods) Remove(ctx context.Context, id, projectID int) (*models.Goods, error) {
	var goods models.Goods
	err := r.db.QueryRowContext(ctx, removeQuery, id, projectID).
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
