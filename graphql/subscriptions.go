package graphql

import (
	"context"
	"sync"
)

func (r subscriptionResolver) MessageAdded(ctx context.Context, id string) (<-chan Message, error) {
	ls := r.ls.addListener(ctx, id)
	// Not sure if I need to lock here, better safe then sorry.
	r.ls.mtx.Lock()
	ls.mc = make(chan Message, 1)
	r.ls.mtx.Unlock()

	return ls.mc, nil
}

type listenerPool struct {
	mtx sync.Mutex
	ls  map[string]*listener
}

type listener struct {
	mc chan Message
}

func (p *listenerPool) addListener(ctx context.Context, key string) *listener {
	p.mtx.Lock()
	ls := p.ls[key]
	if ls == nil {
		// defaults the listener and insert one into the map.
		ls = &listener{}
	}
	p.ls[key] = ls
	p.mtx.Unlock()

	go func() {
		// When user ends their connection we remove him from the pool.
		<-ctx.Done()
		p.mtx.Lock()
		delete(p.ls, key)
		p.mtx.Unlock()
	}()

	return ls
}

func (p *listenerPool) sendMessage(receiverID string, m Message) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	ls := p.ls[receiverID]
	if ls == nil {
		// listener seems to have gone offline, bail.
		return
	}

	ls.mc <- m
}
