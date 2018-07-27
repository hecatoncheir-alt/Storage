package storage

import (
	"testing"

	"github.com/hecatoncheir/Configuration"
)

func TestIntegrationStorageCanConnectToDatabase(test *testing.T) {
	config := configuration.New()
	store := New(config.Development.Database.Host, config.Development.Database.Port)

	_, err := store.PrepareDataBaseClient(store.GraphAddress)
	if err != nil {
		test.Fail()
	}
}
