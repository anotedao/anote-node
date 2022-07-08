package main

import (
	"fmt"
	"log"

	"github.com/anonutopia/gowaves"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func initWaves() {
	if conf.PublicKey != "PUBLICKEY" {
		gowaves.WNC.Host = "http://localhost"
		gowaves.WNC.Port = 6869
		gowaves.WNC.ApiKey = conf.ApiKey

		pk := crypto.MustPublicKeyFromBase58(conf.PublicKey)
		a, err := proto.NewAddressFromPublicKey(55, pk)
		if err != nil {
			log.Println(err.Error())
		}
		NodeAddress = a.String()
		fmt.Printf("Node Address: %s\n", NodeAddress)
	}
}
