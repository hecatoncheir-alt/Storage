package storage

import (
	"context"
	"google.golang.org/grpc"
	"log"

	dataBaseClient "github.com/dgraph-io/dgo"
	dataBaseAPI "github.com/dgraph-io/dgo/protos/api"
)

// Store is a object with database resource
type Store struct {
	DatabaseGateway string
}

// New is a constructor for Store objects
func New(databaseGateway string) *Store {
	storage := &Store{DatabaseGateway: databaseGateway}
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

	return databaseGraph, nil
}

func (store *Store) Query(request string) (response []byte, err error) {
	client, err := store.PrepareDataBaseClient(store.DatabaseGateway)
	if err != nil {
		return nil, err
	}

	transaction := client.NewTxn()
	resp, err := transaction.Query(context.Background(), request)
	if err != nil {
		return response, err
	}

	return resp.GetJson(), nil
}
