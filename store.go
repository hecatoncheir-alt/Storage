package storage

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	dataBaseClient "github.com/dgraph-io/dgo"
	dataBaseAPI "github.com/dgraph-io/dgo/protos/api"
)

// Store is a object with database resource
type Store struct {
	GraphAddress  string
	GraphGRPCHost string
	GraphGRPCPort int

	Client *dataBaseClient.Dgraph
}

// New is a constructor for Store objects
func New(host string, port int) *Store {
	storage := &Store{}

	storage.GraphGRPCHost = host
	storage.GraphGRPCPort = port
	storage.GraphAddress = fmt.Sprintf("%v:%v", host, port)

	return storage
}

func (store *Store) prepareDataBaseClient() (*dataBaseClient.Dgraph, error) {
	conn, err := grpc.Dial(store.GraphAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	baseClient := dataBaseAPI.NewDgraphClient(conn)
	databaseGraph := dataBaseClient.NewDgraphClient(baseClient)

	return databaseGraph, nil
}
