// file: pkg/tasks/notify.go
package tasks

import "sync"

// subscribers stores active channels for task updates.
var (
	subMu       sync.Mutex
	subscribers = map[chan TaskSnapshot]struct{}{}
)

// Subscribe returns a channel that receives task updates.
func Subscribe() chan TaskSnapshot {
	ch := make(chan TaskSnapshot, 1)
	subMu.Lock()
	subscribers[ch] = struct{}{}
	subMu.Unlock()
	return ch
}

// Unsubscribe removes the channel from receiving updates and closes it.
func Unsubscribe(ch chan TaskSnapshot) {
	subMu.Lock()
	delete(subscribers, ch)
	subMu.Unlock()
	close(ch)
}

// broadcast sends the task snapshot to all subscribers.
func broadcast(snapshot TaskSnapshot) {
	subMu.Lock()
	for ch := range subscribers {
		select {
		case ch <- snapshot:
		default:
		}
	}
	subMu.Unlock()
}
