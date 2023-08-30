package bt

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func BoardcastTransaction(payload BoardcastPayload) (string, error) {
	u := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	resp, err := http.Post(u, "application/json", nil)
	if err != nil {
		return "", err
	}
	resData := hashResp{}
	if err = json.NewDecoder(resp.Body).Decode(&resData); err != nil {
		return "", err
	}
	return "", nil
}

type BoardcastPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type hashResp struct {
	Hash string `json:"tx_hash"`
}

func GetTransactionStatus(hash string) (TransactionStatus, error) {
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
	return TransactionStatus(resData.Status), nil
}

type transactionStatusResp struct {
	Status string `json:"tx_status"`
}

type TransactionStatus string

const (
	StatusConfirmed    TransactionStatus = "CONFIRMED"
	StatusFailed       TransactionStatus = "FAILED"
	StatusPending      TransactionStatus = "PENDING"
	StatusDoesNotExist TransactionStatus = "DNE"
)
