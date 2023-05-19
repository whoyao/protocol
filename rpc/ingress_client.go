package rpc

import (
	"github.com/livekit/psrpc"
	"github.com/whoyao/protocol/livekit"
)

type IngressClient interface {
	IngressInternalClient
	IngressHandlerClient
}

type ingressClient struct {
	IngressInternalClient
	IngressHandlerClient
}

func NewIngressClient(nodeID livekit.NodeID, bus psrpc.MessageBus) (IngressClient, error) {
	if bus == nil {
		return nil, nil
	}

	clientID := string(nodeID)
	internalClient, err := NewIngressInternalClient(clientID, bus)
	if err != nil {
		return nil, err
	}
	handlerClient, err := NewIngressHandlerClient(clientID, bus)
	if err != nil {
		return nil, err
	}
	return &ingressClient{
		IngressInternalClient: internalClient,
		IngressHandlerClient:  handlerClient,
	}, nil
}
