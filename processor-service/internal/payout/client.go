package payout

import (
	"bytes"
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

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type UpdateBankReferenceRequest struct {
	BankReference string `json:"bank_reference"`
}

func (c *Client) UpdateStatus(
	payoutID string,
	status string,
) error {

	reqBody := UpdateStatusRequest{
		Status: status,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/payouts/%s/status", c.baseURL, payoutID),
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"failed to update payout status: %d",
			resp.StatusCode,
		)
	}

	return nil
}

func (c *Client) UpdateBankReference(
	payoutID string,
	bankReference string,
) error {

	reqBody := UpdateBankReferenceRequest{
		BankReference: bankReference,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/payouts/%s/bank-reference", c.baseURL, payoutID),
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"failed to update bank reference: %d",
			resp.StatusCode,
		)
	}

	return nil
}