package ports

import (
	"context"
	"test/internal/app/models"
)

type Cache interface {
	CacheGoodsMetadata(ctx context.Context, meta models.PaginationMeta) error
	GetGoodsMetadata(ctx context.Context) (*models.PaginationMeta, error)
	ClearGoodsMetadata(ctx context.Context) error
	CacheGoods(ctx context.Context, offset, limit int, goods []*models.Goods) error
	GetGoods(ctx context.Context, offset, limit int) ([]*models.Goods, error)
	ClearGoods(ctx context.Context) error
}

type QuePublisher interface {
	PublishGoods(goods *models.Goods)
}

type GoodsRepo interface {
	Exists(ctx context.Context, id, projectID int) (bool, error)
	Get(ctx context.Context, limit, offset int) ([]*models.Goods, error)
	GetPaginationMeta(ctx context.Context) (int, int, error)
	Create(ctx context.Context, projectID int, name string) (*models.Goods, error)
	Update(ctx context.Context, id, projectID int, name string, description *string) (*models.Goods, error)
	UpdatePriority(ctx context.Context, id, projectID, priority int) ([]*models.ReprioritizedGoods, error)
	Remove(ctx context.Context, id, projectID int) (*models.Goods, error)
}
