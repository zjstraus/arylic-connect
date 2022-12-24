package tcp

import (
	"context"
	"log"
	"time"
)

func (t *Transport) asyncWriteLoop() {
	for {
		select {
		case <-t.listenerCloser:
			return
		case msg := <-t.outgoingQueue:
			err := t.writeMessage(msg)
			if err != nil {
				log.Println(err.Error())
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func (t *Transport) SendMessage(ctx context.Context, message string) error {
	select {
	case t.outgoingQueue <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
