package extWebsocket

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
)

func (wrapper *ExternalWebsocketWrapper) StatusChanges(ctx context.Context, target string) (*rpc.Subscription, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.HttpMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	channelContext, channelContextCancel := context.WithCancel(context.Background())
	incomingChannel := connection.StatusChangeChannel(channelContext)
	sub := notifier.CreateSubscription()

	go func() {
		initialStatus, initialErr := connection.GetStatus(channelContext)
		if initialErr == nil {
			notifier.Notify(sub.ID, initialStatus)
		}
		for {
			select {
			case change := <-incomingChannel:
				notifier.Notify(sub.ID, change)
			case <-sub.Err():
				channelContextCancel()
				return
			}
		}
	}()

	return sub, nil
}
