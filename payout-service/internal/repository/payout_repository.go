package repository

import (
	"context"

	"github.com/Bharat1Rajput/payout-service/internal/model"
)

type PayoutRepository interface {
	Create(
		ctx context.Context, payout *model.Payout,
	) error 
}
