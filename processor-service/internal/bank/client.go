package bank

import (
	"bytes"
	"encoding/json"
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

type CreatePayoutRequest struct {
	PayoutID string `json:"payout_id"`
	Amount   int64  `json:"amount"`
}

type CreatePayoutResponse struct {
	BankReference string `json:"bank_reference"`
	Status        string `json:"status"`
}

func (c *Client) CreatePayout(
	req CreatePayoutRequest,
) (*CreatePayoutResponse, error) {

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}


	resp, err := http.Post(
		c.baseURL+"/payouts",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response CreatePayoutResponse

	err = json.NewDecoder(resp.Body).
		Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
