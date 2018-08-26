package storage

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"fmt"
	dataBaseClient "github.com/dgraph-io/dgo"
	dataBaseAPI "github.com/dgraph-io/dgo/protos/api"
)

// Store is a object with database resource
type Store struct {
	DatabaseGateway string
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

func (store *Store) Mutate(setJson []byte) (uid string, err error) {
	client, err := store.PrepareDataBaseClient(store.DatabaseGateway)
	if err != nil {
		return "", err
	}

	mutation := dataBaseAPI.Mutation{
		SetJson:   setJson,
		CommitNow: true}

	transaction := client.NewTxn()
	assigned, err := transaction.Mutate(context.Background(), &mutation)
	if err != nil {
		return "", err
	}

	uid = assigned.Uids["blank-0"]

	return uid, nil
}

func (store *Store) SetNQuads(subject, predicate, object string) error {

	final := fmt.Sprintf(`<%s> <%s> %s .`, subject, predicate, object)

	mutation := dataBaseAPI.Mutation{
		SetNquads: []byte(final),
		CommitNow: true}

	client, err := store.PrepareDataBaseClient(store.DatabaseGateway)
	if err != nil {
		return err
	}

	transaction := client.NewTxn()
	_, err = transaction.Mutate(context.Background(), &mutation)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) DeleteJSON(encodedJSON []byte) error {

	mutation := dataBaseAPI.Mutation{
		DeleteJson: encodedJSON,
		CommitNow:  true}

	client, err := store.PrepareDataBaseClient(store.DatabaseGateway)
	if err != nil {
		return err
	}

	transaction := client.NewTxn()
	_, err = transaction.Mutate(context.Background(), &mutation)
	if err != nil {
		return err
	}

	return nil
}
