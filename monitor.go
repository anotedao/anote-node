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

func (m *Monitor) payToNetwork() {
	amount := SatInBTC
	amount = amount - (amount / 10)
	sendAnote(NetworkNode, int(amount))
}

func (m *Monitor) isGeneratingNode() bool {
	bar, err := gowaves.WNC.BlocksAt(uint(m.height))
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
				m.payToNetwork()
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
