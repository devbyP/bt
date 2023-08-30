package main

import (
	"log"
	"time"

	"github.com/devbyP/bt"
)

func main() {
	hash, err := bt.BoardcastTransaction(bt.BoardcastPayload{
		Symbol: "BTC",
		Price:  100000,
		// use unix time and convert to uint64
		Timestamp: uint64(time.Now().Unix()),
	})
	if err != nil {
		log.Fatal(err)
	}
	// do something with the returned hash.

	// check the status of transaction using hash.
	status, err := bt.GetTransactionStatus(hash)
	if err != nil {
		log.Fatal(err)
	}
	// handle code by status.
	switch status {
	case bt.StatusConfirmed:
		// handle confirmed transaction.
	case bt.StatusPending:
		// handle pending transaction.
	case bt.StatusFailed:
		// handle failed transaction.
	case bt.StatusDoesNotExist:
		// handle non exist transaction.
	default:
		// handle unknown status.
	}
	log.Println(status)
}
