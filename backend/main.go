package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/CZnavody19/supply-chain/src/config"
	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/setup"
	"github.com/gorilla/mux"
	"github.com/nextap-solutions/goNextService"
	"github.com/nextap-solutions/goNextService/components"
	"github.com/rs/cors"
	"go.uber.org/zap"

	httpHandler "github.com/CZnavody19/supply-chain/src/http"
)

type ServerComponents struct {
	httpServer goNextService.Component
}

func main() {
	err := serve()
	if err != nil {
		fmt.Println("Error running the application: ", err)
		os.Exit(1)
	}
}

func serve() error {
	configuration := config.LoadConfig()

	api, err := setupService(configuration)
	if err != nil {
		return err
	}

	app := goNextService.NewApplications(api.httpServer)
	app.WithLogger(zap.S())

	return app.Run()
}

func setupService(configuration *config.Config) (*ServerComponents, error) {
	setup.InitLogger(*configuration)
	s, _ := json.MarshalIndent(configuration, "", "\t")
	zap.S().Info("Logger initialized successfully", string(s))

	dbConn, err := setup.SetupDb(&configuration.DBConfig)
	if err != nil {
		zap.S().Error("Error setting up db connection", err)
		return nil, err
	}

	dbStore := db.NewDatabaseStore(dbConn)

	if os.Getenv("SEED_DB") == "true" {
		if err := dbStore.SeedDatabase(context.Background()); err != nil {
			zap.S().Warn("Error seeding database", err)
		}
	}

	httpHandler := httpHandler.NewHttpHandler(dbStore)

	router := mux.NewRouter()

	setup.SetupHTTPHandlers(router, httpHandler)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(router)
	api := http.Server{
		Addr:              "0.0.0.0:" + configuration.Server.Port,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		IdleTimeout:       0,
		WriteTimeout:      configuration.Server.WriteTimeout,
		Handler:           handler,
	}
	httpComponent := components.NewHttpComponent(handler, components.WithHttpServer(&api))

	return &ServerComponents{
		httpServer: httpComponent,
	}, nil
}
