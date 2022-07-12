package main

import (
	"fmt"
	"log"
	"net/http"
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
	go func() {
		for {
			m.ping()
			time.Sleep(time.Second * PingTick)
		}
	}()

	for {
		balance := balance()
		if balance > RewardFee && balance == balanceC() {
			m.sendRewards()
		}
		time.Sleep(time.Second * MonitorTick)
	}
}

func (m *Monitor) ping() {
	url, err := joinUrl(MasterNodeUrl, fmt.Sprintf("/ping/%s/%s", conf.OwnerAddress, NodeAddress))
	if err != nil {
		log.Println(err.Error())
	}

	res, err := http.Get(url.String())
	if err != nil {
		log.Println(err.Error())
	}
	res.Body.Close()
}

func initMonitor() *Monitor {
	m := &Monitor{}
	m.run()
	return m
}
