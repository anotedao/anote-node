package main

import (
	"log"

	"github.com/anonutopia/gowaves"
)

func sendAnote(recipient string, amount int) {
	atr := &gowaves.AssetsTransferRequest{
		Recipient: recipient,
		Amount:    amount,
		Fee:       AnoteFee,
		Sender:    NodeAddress,
	}
	_, err := gowaves.WNC.AssetsTransfer(atr)
	if err != nil {
		log.Println(err.Error())
	}
}
