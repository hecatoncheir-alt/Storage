package storage

import (
	"log"

	"google.golang.org/grpc"

	dataBaseClient "github.com/dgraph-io/dgo"
	dataBaseAPI "github.com/dgraph-io/dgo/protos/api"
)

// Store is a object with database resource
type Store struct {
	DatabaseGateway string
	Client          *dataBaseClient.Dgraph
}

// New is a constructor for Store objects
func New(databaseGateway string) *Store {
	storage := &Store{
		DatabaseGateway: databaseGateway}

	return storage
}

func (store *Store) PrepareDataBaseClient(databaseGateway string) (*dataBaseClient.Dgraph, error) {
	conn, err := grpc.Dial(databaseGateway, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	baseClient := dataBaseAPI.NewDgraphClient(conn)
	databaseGraph := dataBaseClient.NewDgraphClient(baseClient)

	store.Client = databaseGraph

	return databaseGraph, nil
}
