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
			log.Println("Mined")
			m.balance = balance()
		}
		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() *Monitor {
	m := &Monitor{}
	m.run()
	return m
}
