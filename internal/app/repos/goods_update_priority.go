package repos

import (
	"context"
	"test/internal/app/models"
)

const updatePriorityQuery = `
	UPDATE goods 
		SET priority = $3
	WHERE id >= $1 AND project_id = $2
	RETURNING 
		id, 
		priority
`

func (r *Goods) UpdatePriority(ctx context.Context, id, projectID, priority int) ([]*models.ReprioritizedGoods, error) {
	result := make([]*models.ReprioritizedGoods, 0)

	rows, err := r.db.QueryContext(ctx, updatePriorityQuery, id, projectID, priority)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goods models.ReprioritizedGoods
		err := rows.Scan(&goods.Id, &goods.Priority)
		if err != nil {
			return nil, err
		}
		result = append(result, &goods)
	}

	return result, err
}
