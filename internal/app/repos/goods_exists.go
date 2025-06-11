package repos

import "context"

const existsQuery = `
	SELECT EXISTS (
		SELECT 1 FROM goods 
		WHERE id = $1 AND project_id = $2
	)
`

func (r *Goods) Exists(ctx context.Context, id, projectID int) (bool, error) {
	var result bool
	err := r.db.QueryRowContext(ctx, existsQuery, id, projectID).Scan(&result)
	return result, err
}
