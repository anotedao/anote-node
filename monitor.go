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

	if total > 0 {
		total -= 2 * AnoteFee

		amountFee := int(float64(total) * float64(fee.Value) / float64(1000))
		sendAnote(NetworkNode, amountFee)

		amountOwner := total - amountFee
		sendAnote(conf.OwnerAddress, amountOwner)
	}
}

func (m *Monitor) isGeneratingNode() bool {
	bar, err := gowaves.WNC.BlocksAt(uint(m.height - 5))
	if err != nil {
		log.Println(err.Error())
	}
	return bar.Generator == NodeAddress
}

func (m *Monitor) run() {
	m.height = m.getHeight()
	for {
		if m.getHeight() > m.height {
			m.height = m.getHeight()
			if m.isGeneratingNode() {
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
