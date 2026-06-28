package bank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

type GetPayoutResponse struct {
	PayoutID      string `json:"payout_id"`
	Status        string `json:"status"`
	BankReference string `json:"bank_reference,omitempty"`
}

func (c *Client) GetPayout(
	bankReference string,
) (*GetPayoutResponse, error) {

	resp, err := http.Get(
		fmt.Sprintf("%s/payouts/%s", c.baseURL, bankReference),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"bank returned status %d",
			resp.StatusCode,
		)
	}

	var payout GetPayoutResponse

	err = json.NewDecoder(resp.Body).Decode(
		&payout,
	)

	if err != nil {
		return nil, err
	}

	return &payout, nil
}