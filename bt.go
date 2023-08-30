package bt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func BoardcastTransaction(payload BoardcastPayload) (string, error) {
	u := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(u, "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	resData := hashResp{}
	if err = json.NewDecoder(resp.Body).Decode(&resData); err != nil {
		return "", err
	}
	return resData.Hash, nil
}

type BoardcastPayload struct {
	Symbol string `json:"symbol"`
	Price  uint64 `json:"price"`

	// unix time in type 'uint64'
	Timestamp uint64 `json:"timestamp"`
}

type hashResp struct {
	Hash string `json:"tx_hash"`
}

func GetTransactionStatus(hash string) (string, error) {
	// add hash to the path.
	u, err := url.Parse("https://mock-node-wgqbnxruha-as.a.run.app")
	u = u.JoinPath("check", hash)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(u.String())
	resData := transactionStatusResp{}
	if err = json.NewDecoder(resp.Body).Decode(&resData); err != nil {
		return "", nil
	}
	return resData.Status, nil
}

type transactionStatusResp struct {
	Status string `json:"tx_status"`
}

const (
	StatusConfirmed    = "CONFIRMED"
	StatusFailed       = "FAILED"
	StatusPending      = "PENDING"
	StatusDoesNotExist = "DNE"
)
