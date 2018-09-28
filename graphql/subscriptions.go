package graphql

import (
	"context"
	"fmt"
	"sync"
)

func (r subscriptionResolver) MessageAdded(ctx context.Context, id string) (<-chan Message, error) {
	c := r.ls.addListener(ctx, id)

	return c, nil
}

type listenerPool struct {
	mtx sync.Mutex
	ls  map[string]*listener
}

type listener struct {
	mc chan Message
}

func (p *listenerPool) addListener(ctx context.Context, key string) <-chan Message {
	p.mtx.Lock()
	ls := p.ls[key]
	if ls == nil {
		// defaults the listener and insert one into the map.
		ls = &listener{mc: make(chan Message, 1)}
	}
	p.ls[key] = ls
	p.mtx.Unlock()

	go func() {
		<-ctx.Done()
		p.mtx.Lock()
		delete(p.ls, key)
		p.mtx.Unlock()
	}()

	return ls.mc
}

func (p *listenerPool) sendMessage(receiverID string, m Message) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	fmt.Println(receiverID)
	ls := p.ls[receiverID]
	if ls == nil {
		fmt.Println("nil")
		// listener seems to have gone online, bail.
		return
	}

	ls.mc <- m
}
