package repos

import (
	"context"
	"test/internal/app/models"
)

const deleteQuery = `
	DELETE FROM goods 
	WHERE id = $1 AND project_id = $2
	RETURNING 
		id, 
		project_id, 
		removed
`

func (r *Goods) Delete(ctx context.Context, id, projectID int) (*models.ShortGoods, error) {
	var deleted models.ShortGoods
	err := r.db.QueryRowContext(ctx, deleteQuery, id, projectID).
		Scan(
			&deleted.Id,
			&deleted.ProjectId,
			&deleted.Removed,
		)
	return &deleted, err
}
