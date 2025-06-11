package repos

import "context"

const getPaginationMetaQuery = `
	SELECT
  	COUNT(*) AS total,
  	COUNT(*) FILTER (WHERE removed = true) AS removed_count
	FROM goods
`

func (r *Goods) GetPaginationMeta(ctx context.Context) (int, int, error) {
	var total, removed int
	err := r.db.QueryRowContext(ctx, getPaginationMetaQuery).Scan(&total, &removed)
	return total, removed, err
}
