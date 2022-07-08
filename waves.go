package main

import (
	"fmt"
	"log"

	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func initWaves() {
	pk := crypto.MustPublicKeyFromBase58(conf.PublicKey)
	a, err := proto.NewAddressFromPublicKey(55, pk)
	if err != nil {
		log.Println(err.Error())
	}
	NodeAddress = a.String()
	fmt.Printf("Node Address: %s\n", NodeAddress)
}
