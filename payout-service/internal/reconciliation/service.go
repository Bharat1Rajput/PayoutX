package reconciliation

import (
	"context"
	"log"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/bank"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
)

type Service struct {
	repo       repository.PayoutRepository
	bankClient *bank.Client
}

func NewService(
	repo repository.PayoutRepository,
	bankClient *bank.Client,
) *Service {

	return &Service{
		repo:       repo,
		bankClient: bankClient,
	}
}

func (s *Service) Reconcile(
	ctx context.Context,
) error {

	payouts, err := s.repo.GetProcessing(ctx)
	if err != nil {
		return err
	}

	log.Printf(
		"found %d processing payouts",
		len(payouts),
	)

	for _, payout := range payouts {

		bankPayout, err := s.bankClient.GetPayout(
			payout.BankReference,
		)

		if err != nil {
			log.Printf(
				"failed to fetch bank status for payout=%s: %v",
				payout.BankReference,
				err,
			)
			continue
		}

		if payout.Status == bankPayout.Status {
			continue
		}

		log.Printf(
			"reconciliation updating payout=%s from %s to %s",
			payout.BankReference,
			payout.Status,
			bankPayout.Status,
		)

		err = s.repo.UpdateStatus(
			ctx,
			payout.BankReference,
			bankPayout.Status,
		)

		if err != nil {
			log.Printf(
				"failed updating payout=%s: %v",
				payout.BankReference,
				err,
			)
		}
	}

	return nil
}