package main

import (
	"log"
	"time"
)

type Monitor struct {
	balance int
}

func (m *Monitor) run() {
	m.balance = balance()
	for {
		if balance() > m.balance {
			amount := balance() - m.balance
			amount = amount - (amount / 500)
			sendAINT(NetworkNode, amount)
			m.balance = balance()
			log.Printf("Mined: %d", amount)
		}
		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() *Monitor {
	m := &Monitor{}
	m.run()
	return m
}
