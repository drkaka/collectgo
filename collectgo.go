package collectgo

import (
	"sync"
	"time"
)

// input plugin
type input struct {
	tick *time.Ticker
	// Request when scheduled, returning name, value and error
	request *func() (string, float64, error)
}

type inputsHolder struct {
	all map[string]*input
	l   *sync.RWMutex
}

var (
	// inputs to hold all input plugins
	inputs inputsHolder
)

func init() {
	inputs = inputsHolder{
		all: make(map[string]*input, 0),
		l:   new(sync.RWMutex),
	}
}

// AddInput and start
func AddInput(name string, interval int, fn func() (float64, error)) {
	tk := time.NewTicker(time.Duration(interval) * time.Second)
	go func() {
		for {
			<-tk.C
			v, err := fn()
			if err != nil {
				// handle error
			} else {
				// handle v
			}
		}
	}()
}
