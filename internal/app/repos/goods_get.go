package repos

import (
	"context"
	"test/internal/app/models"
)

const getQuery = `
	SELECT 
		id, 
		project_id, 
		name, 
		description, 
		priority, 
		removed, 
		created_at
	FROM goods
	ORDER BY priority DESC
	LIMIT $1 
	OFFSET $2
`

func (r *Goods) Get(ctx context.Context, limit, offset int) ([]*models.Goods, error) {
	result := make([]*models.Goods, 0)

	rows, err := r.db.QueryContext(ctx, getQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goods models.Goods
		err := rows.Scan(
			&goods.Id,
			&goods.ProjectId,
			&goods.Name,
			&goods.Description,
			&goods.Priority,
			&goods.Removed,
			&goods.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, &goods)
	}

	return result, err
}
