package main

import (
	"time"
)

type Monitor struct {
}

func (m *Monitor) sendRewards() {
	total := balance()

	if total > RewardFee {
		total -= RewardFee
		sendAnote(total)
	}
}

func (m *Monitor) run() {
	for {
		balance := balance()
		if balance > RewardFee && balance == balanceC() {
			m.sendRewards()
		}
		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() *Monitor {
	m := &Monitor{}
	m.run()
	return m
}
