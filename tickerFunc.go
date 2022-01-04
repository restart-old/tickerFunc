package tickerFunc

import (
	"time"
)

type Ticker struct {
	*time.Ticker
	f     func()
	close chan bool
}

func NewTicker(t time.Duration, f func()) *Ticker {
	return &Ticker{Ticker: time.NewTicker(t), f: f, close: make(chan bool)}
}

func (t *Ticker) Start() {
	for {
		select {
		case _, running := <-t.close:
			if !running {
				return
			}
		case <-t.C:
			t.f()
		}
	}
}

func (t *Ticker) Stop() {
	t.Ticker.Stop()
	close(t.close)
}
