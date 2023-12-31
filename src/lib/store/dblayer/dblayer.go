package dblayer

import (
	"github.com/event-management/src/lib/store"
	"github.com/event-management/src/lib/store/mongolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

func NewSoreHandler(options DBTYPE, connection string) (store.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
