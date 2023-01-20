package serialmedia

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
)

func (wrapper *SerialMediaWrapper) MetadataChanges(ctx context.Context, target string) (*rpc.Subscription, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	channelContext, channelContextCancel := context.WithCancel(context.Background())
	incomingChannel := connection.MetadataChangeChannel(channelContext)
	sub := notifier.CreateSubscription()

	go func() {
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

func (wrapper *SerialMediaWrapper) VolumeChanges(ctx context.Context, target string) (*rpc.Subscription, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	channelContext, channelContextCancel := context.WithCancel(context.Background())
	incomingChannel := connection.VolumeChannel(channelContext)
	sub := notifier.CreateSubscription()

	go func() {
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

func (wrapper *SerialMediaWrapper) MuteChanges(ctx context.Context, target string) (*rpc.Subscription, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	channelContext, channelContextCancel := context.WithCancel(context.Background())
	incomingChannel := connection.MuteChannel(channelContext)
	sub := notifier.CreateSubscription()

	go func() {
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

func (wrapper *SerialMediaWrapper) MediaReady(ctx context.Context, target string) (*rpc.Subscription, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	channelContext, channelContextCancel := context.WithCancel(context.Background())
	incomingChannel := connection.MediaReadyChannel(channelContext)
	sub := notifier.CreateSubscription()

	go func() {
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
