package contracts

import "github.com/event-management/src/lib/store"

// LocationCreatedEvent is emitted whenever a location is created
type LocationCreatedEvent struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Address string       `json:"address"`
	Country string       `json:"country"`
	Halls   []store.Hall `json:"halls"`
}

// EventName returns the event's name
func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}
