package stock

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/utils"

	model "github.com/BigNutJaa/user-service/internals/model/stock"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
	Get(ctx context.Context, request *model.FitterReadStock) (*model.ReadResponseStock, error)
	Delete(ctx context.Context, input *model.FitterDeleteStock) (*model.DeleteResponseStock, error)
	Update(ctx context.Context, request *model.FitterUpdateStock) (*model.UpdateResponseStock, error)
	List(ctx context.Context, request *model.FitterListStock, pagination *utils.PaginationOptions) (*utils.Pagination, error)
}
