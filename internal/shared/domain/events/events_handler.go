package shared_events

import "sync"

type EventHandler interface {
	Handler(event any)
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Register(eventName string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[eventName] = append(d.handlers[eventName], handler)
}

func (d *EventDispatcher) Dispatch(eventName string, event any) {
	d.mu.Lock()
	defer d.mu.Unlock()

	handlers, ok := d.handlers[eventName]

	if !ok {
		return
	}

	for _, h := range handlers {
		go func(h EventHandler) {
			defer func() {
				if r := recover(); r != nil {
					println("panic recovered in handler for", eventName)
				}
			}()
		}(h)
	}
}
