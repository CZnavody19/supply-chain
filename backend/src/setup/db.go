package setup

import (
	"context"

	"github.com/CZnavody19/supply-chain/src/config"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func SetupDb(config *config.DBConfig) (*neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(config.ConnectionURI, neo4j.BasicAuth(config.Username, config.Password, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		return nil, err
	}

	return &driver, nil
}
