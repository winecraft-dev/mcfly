package mcproxy

import (
	"context"

	"github.com/jnorman-us/mcfly/mcserver/manager"
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

type MCProxy struct {
	*proxy.Proxy
	servers manager.ServerManager
}

func NewMCProxy(p *proxy.Proxy, s manager.ServerManager) *MCProxy {
	return &MCProxy{
		Proxy:   p,
		servers: s,
	}
}

func (p *MCProxy) Init(ctx context.Context) error {
	// event.Subscribe(p.Event(), 0, p.HandlePreLogin)
	event.Subscribe(p.Event(), 0, p.HandlePlayerChooseInitialServer)
	event.Subscribe(p.Event(), 0, p.HandlePlayerDisconnected)
	event.Subscribe(p.Event(), 0, func(e *proxy.PreLoginEvent) {
		// Kicks every player
		e.Deny(&component.Text{Content: "Sorry, the server is in maintenance."})
	})

	return nil
}
