package util

import "sync"

type SelfPublisher[E interface{}] interface {
	SetPublisher(p Publisher[E])
}

type Publisher[E interface{}] interface {
	Subscribe(chan E)
	Publish()
	Read(func(E))
	Write(func(E))
}

type Observable[E interface{}] struct {
	itemMu      sync.RWMutex
	item        E
	subMu       sync.RWMutex
	subscribers []chan E
}

func (o *Observable[E]) Subscribe(s chan E) {
	o.subMu.Lock()
	defer o.subMu.Unlock()
	o.subscribers = append(o.subscribers, s)
}

func (o *Observable[E]) Publish() {
	o.subMu.RLock()
	defer o.subMu.RUnlock()
	for _, s := range o.subscribers {
		select {
		case s <- o.item:
		default:
		}
	}
}

func (o *Observable[E]) Read(ui func(E)) {
	o.itemMu.RLock()
	defer o.itemMu.RUnlock()
	ui(o.item)
	o.Publish()
}

func (o *Observable[E]) Write(ui func(E)) {
	o.itemMu.Lock()
	defer o.itemMu.Unlock()
	ui(o.item)
	o.Publish()
}

func NewObservable[E interface{}](o E) *Observable[E] {
	return &Observable[E]{
		itemMu:      sync.RWMutex{},
		item:        o,
		subscribers: []chan E{},
	}
}
