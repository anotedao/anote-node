package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

type Monitor struct {
	height int
}

func (m *Monitor) getHeight() int {
	bhr, err := gowaves.WNC.BlocksHeight()
	if err != nil {
		log.Println(err.Error())
	}
	return bhr.Height
}

func (m *Monitor) sendRewards() {
	fee, err := gowaves.WNC.AddressesDataKey(NetworkNode, "fee")
	if err != nil {
		log.Println(err.Error())
	}

	total := balance()

	if total >= int(SatInBTC) {
		total -= (2 * AnoteFee)

		amountOwner := int(float64(total) * float64(1000-fee.Value) / float64(1000))
		log.Printf("Amount owner: %d\n", amountOwner)
		if amountOwner > AnoteFee {
			sendAnote(conf.OwnerAddress, amountOwner)
		}

		amountFee := total - amountOwner
		log.Printf("Amount fee: %d\n", amountFee)
		if amountFee > AnoteFee {
			sendAnote(NetworkNode, amountFee)
		}
	}
}

func (m *Monitor) run() {
	m.height = m.getHeight()
	for {
		if m.getHeight() > m.height {
			m.height = m.getHeight()
			if balance() > AnoteFee {
				m.sendRewards()
				log.Println("Mined.")
			}
		}
		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() *Monitor {
	m := &Monitor{}
	m.run()
	return m
}
