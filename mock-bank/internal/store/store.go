package store

import "sync"

type Payout struct {
	PayoutID      string
	Status        string
	BankReference string
}

var (
	payouts = make(map[string]Payout)
	mu      sync.RWMutex
)

func Save(payout Payout) {
	mu.Lock()
	defer mu.Unlock()

	payouts[payout.BankReference] = payout
}

func Get(
	bankReference string,
) (Payout, bool) {

	mu.RLock()
	defer mu.RUnlock()

	payout, ok := payouts[bankReference]

	return payout, ok
}
