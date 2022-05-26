package mev_boost

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type Repository interface {
	Get(ctx context.Context, ID int64) (*entity.Payload, error)
	Create(ctx context.Context) (*int64, error)
}