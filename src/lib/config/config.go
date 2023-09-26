package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/event-management/src/lib/store/dblayer"
)

var (
	DBTypeDefault              = dblayer.DBTYPE("mongodb")
	EventDBConnection          = "mongodb://127.0.0.1/events"
	BookingDBConnection        = "mongodb://127.0.0.1/bookings"
	EventRestfulEndpoint       = "localhost:8181"
	BookingRestfulEndpoint     = "localhost:8282"
	RestfulTLSEPDefault        = "localhost:9191"
	MessageBrokerTypeDefault   = "amqp"
	AMQPMessageBrokerDefault   = "amqp://guest:guest@localhost:5672"
	KafkaMessageBrokersDefault = []string{"localhost:9092"}
)

type ServiceConfig struct {
	Databasetype        dblayer.DBTYPE `json:"databasetype"`
	EventDBConnection   string         `json:"eventdbconnection"`
	BookingDBConnection string         `json:"bookingdbconnection"`

	EventRestfulEndpoint   string `json:"event_api_endpoint"`
	BookingRestfulEndpoint string `json:"booking_api_endpoint"`

	RestfulTLSEndPint   string   `json:"restfulapi-tlsendpoint"`
	MessageBrokerType   string   `json:"message_broker_type"`
	AMQPMessageBroker   string   `json:"amqp_message_broker"`
	KafkaMessageBrokers []string `json:"kafka_message_brokers"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		EventDBConnection,
		BookingDBConnection,
		EventRestfulEndpoint,
		BookingRestfulEndpoint,
		RestfulTLSEPDefault,
		MessageBrokerTypeDefault,
		AMQPMessageBrokerDefault,
		KafkaMessageBrokersDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}
	json.NewDecoder(file).Decode(&conf)

	if v := os.Getenv("LISTEN_URL_EVENT"); v != "" {
		conf.EventRestfulEndpoint = v
	}
	if v := os.Getenv("LISTEN_URL_BOOKING"); v != "" {
		conf.BookingRestfulEndpoint = v
	}
	if v := os.Getenv("MONGO_URL_EVENT"); v != "" {
		conf.Databasetype = "mongodb"
		conf.EventDBConnection = v
	}
	if v := os.Getenv("MONGO_URL_BOOKING"); v != "" {
		conf.Databasetype = "mongodb"
		conf.BookingDBConnection = v
	}
	if v := os.Getenv("AMQP_BROKER_URL"); v != "" {
		conf.MessageBrokerType = "amqp"
		conf.AMQPMessageBroker = v
	} else if v := os.Getenv("KAFKA_BROKER_URLS"); v != "" {
		conf.MessageBrokerType = "kafka"
		conf.KafkaMessageBrokers = strings.Split(v, ",")
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}

/*
  openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
*/
