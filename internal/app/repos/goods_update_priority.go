package repos

import (
	"context"
	"database/sql"
	"fmt"
	"test/internal/app/models"
)

const getPriorityQuery = `
	SELECT priority
	FROM goods 
	WHERE id = $1 AND project_id = $2
`

const updatePriorityQuery = `
	UPDATE goods
	SET priority = CASE
			WHEN id = $1 THEN $3
			WHEN $3::int < $4::int AND priority >= $3::int AND priority < $4::int THEN priority + 1
			WHEN $3::int > $4::int AND priority <= $3::int AND priority > $4::int THEN priority - 1
			ELSE priority
	END
	WHERE project_id = $2
	RETURNING id, priority
	ORDER BY priority
`

func (r *Goods) UpdatePriority(ctx context.Context, id, projectID, newPriority int) ([]*models.ReprioritizedGoods, error) {
	tx, _ := r.db.BeginTx(ctx, nil)

	oldPriority, err := r.getPriority(ctx, tx, id, projectID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	updated, err := r.reprioritize(ctx, tx, id, projectID, newPriority, oldPriority)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return updated, err
}

func (r *Goods) reprioritize(ctx context.Context, tx *sql.Tx, id, projectID, newPriority, oldPriority int) ([]*models.ReprioritizedGoods, error) {
	result := make([]*models.ReprioritizedGoods, 0)

	rows, err := tx.QueryContext(ctx, updatePriorityQuery, id, projectID, newPriority, oldPriority)
	if err != nil {
		fmt.Println("Here", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goods models.ReprioritizedGoods
		err := rows.Scan(&goods.Id, &goods.Priority)
		if err != nil {
			fmt.Println("Here 2", err)
			return nil, err
		}
		result = append(result, &goods)
	}

	return result, nil
}

func (r *Goods) getPriority(ctx context.Context, tx *sql.Tx, id, projectID int) (int, error) {
	var priority int

	err := tx.QueryRowContext(ctx, getPriorityQuery, id, projectID).Scan(&priority)
	if err != nil {
		return -1, err
	}

	return priority, nil
}
