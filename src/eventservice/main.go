package main

import (
	"flag"
	"fmt"

	"github.com/event-management/config"
	"github.com/event-management/src/eventservice/api"
	"github.com/event-management/src/lib/msgqueue"
	msgqueue_amqp "github.com/event-management/src/lib/msgqueue/amqp"
	"github.com/event-management/store/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `./config/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := config.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}

	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewSoreHandler(config.Databasetype, config.DBConnection)

	fmt.Println("Serving API")
	//RESTful API start
	err := api.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
	if err != nil {
		panic(err)
	}
}
