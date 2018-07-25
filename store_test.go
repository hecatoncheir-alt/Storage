package storage

import (
	"testing"

	"github.com/hecatoncheir/Configuration"
)

func TestIntegrationStorageCanConnectToDatabase(test *testing.T) {
	config := configuration.New()
	store := New(config.Development.Database.Host, config.Development.Database.Port)

	_, err := store.prepareDataBaseClient()
	if err != nil {
		test.Fail()
	}
}
