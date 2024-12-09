package ticker

import (
	"fmt"
	"time"
)

type Ticker struct {
	I        int64
	N        int64
	Interval time.Duration
	Handler  func(i int64)
}

func NewTicker(n int64, interval time.Duration) *Ticker {
	return &Ticker{
		I:        0,
		N:        n,
		Interval: interval,
		Handler: func(i int64) {
			fmt.Printf("%d\n", i)
		},
	}
}

func (t *Ticker) Start() {
	ticker := time.NewTicker(t.Interval)

	for {
		<-ticker.C
		t.Handler(t.I)
		t.I++
		if t.I >= t.N {
			ticker.Stop()
			return
		}
	}
}
