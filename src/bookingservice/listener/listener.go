package listener

import (
	"fmt"
	"log"

	"github.com/event-management/src/contracts"
	"github.com/event-management/src/lib/msgqueue"
	"github.com/event-management/src/lib/store"
	"gopkg.in/mgo.v2/bson"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      store.DatabaseHandler
}

func (p *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := p.EventListener.Listen("eventCreated")

	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			fmt.Printf("got event %T: %s\n", evt, evt)
			p.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)

		if !bson.IsObjectIdHex(e.ID) {
			log.Printf("event %v did not contain valid object ID", e)
			return
		}

		p.Database.AddEvent(store.Event{ID: bson.ObjectIdHex(e.ID), Name: e.Name})
	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %v", e.ID, e)
		// TODO: No store for locations, yet
	default:
		log.Printf("unknown event type: %T", e)
	}
}
