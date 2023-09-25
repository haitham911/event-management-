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
	DBConnectionDefault        = "mongodb://127.0.0.1"
	RestfulEPDefault           = "localhost:8181"
	RestfulTLSEPDefault        = "localhost:9191"
	MessageBrokerTypeDefault   = "amqp"
	AMQPMessageBrokerDefault   = "amqp://guest:guest@localhost:5672"
	KafkaMessageBrokersDefault = []string{"localhost:9092"}
)

type ServiceConfig struct {
	Databasetype        dblayer.DBTYPE `json:"databasetype"`
	DBConnection        string         `json:"dbconnection"`
	RestfulEndpoint     string         `json:"restfulapi_endpoint"`
	RestfulTLSEndPint   string         `json:"restfulapi-tlsendpoint"`
	MessageBrokerType   string         `json:"message_broker_type"`
	AMQPMessageBroker   string         `json:"amqp_message_broker"`
	KafkaMessageBrokers []string       `json:"kafka_message_brokers"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
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

	if v := os.Getenv("LISTEN_URL"); v != "" {
		conf.RestfulEndpoint = v
	}

	if v := os.Getenv("MONGO_URL"); v != "" {
		conf.Databasetype = "mongodb"
		conf.DBConnection = v
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
