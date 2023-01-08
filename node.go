package main

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func setScript() error {
	var networkByte = byte(55)
	var nodeURL = AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	gs, _ := base64.StdEncoding.DecodeString(generatorScript)

	tr := proto.NewUnsignedSetScriptWithProofs(
		2,
		networkByte,
		sender,
		gs,
		AnoteFee*2,
		ts)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func callScript() error {
	var networkByte = byte(55)
	var nodeURL = AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	rec, err := proto.NewRecipientFromString(NodeAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	args := proto.Arguments{}
	args.Append(proto.NewStringArgument(OwnerAddress))

	call := proto.FunctionCall{
		Name:      "constructor",
		Arguments: args,
	}

	payments := proto.ScriptPayments{}

	fa := proto.OptionalAsset{}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	tr := proto.NewUnsignedInvokeScriptWithProofs(
		2,
		networkByte,
		sender,
		rec,
		call,
		payments,
		fa,
		RewardFee,
		ts)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
