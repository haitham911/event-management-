package main

import (
	"flag"

	"github.com/event-management/src/bookingservice/api"
	"github.com/event-management/src/bookingservice/listener"
	"github.com/event-management/src/lib/config"
	"github.com/event-management/src/lib/msgqueue"
	msgqueue_amqp "github.com/event-management/src/lib/msgqueue/amqp"
	"github.com/event-management/src/lib/store/dblayer"
	"github.com/streadway/amqp"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var eventListener msgqueue.EventListener
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", "./configuration/config.json", "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := config.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		panicIfErr(err)

		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
		panicIfErr(err)

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		panicIfErr(err)

	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	dbhandler, _ := dblayer.NewSoreHandler(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{EventListener: eventListener, Database: dbhandler}
	go processor.ProcessEvents()

	api.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
}
