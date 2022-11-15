package delta

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/delta"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
