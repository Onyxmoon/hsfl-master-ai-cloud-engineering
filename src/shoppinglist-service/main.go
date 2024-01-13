package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/router/middleware/auth"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/rpc/user"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/api/router"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/config"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingList"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingListEntry"
	"log"
	"net/http"
)

func main() {
	var configuration = loadConfiguration()

	var shoppingListRepository userShoppingList.Repository = userShoppingList.NewRQLiteRepository(configuration.Database.GetConnectionString())
	var shoppingListController userShoppingList.Controller = userShoppingList.NewDefaultController(shoppingListRepository)

	var shoppingListEntryRepository userShoppingListEntry.Repository = userShoppingListEntry.NewRQLiteRepository(configuration.Database.GetConnectionString())
	var shoppingListEntryController userShoppingListEntry.Controller = userShoppingListEntry.NewDefaultController(shoppingListEntryRepository, shoppingListRepository)

	startHTTPServer(configuration, &shoppingListController, &shoppingListEntryController)
}

func loadConfiguration() *config.ServiceConfiguration {
	godotenv.Load()

	serviceConfiguration := &config.ServiceConfiguration{}
	if err := env.Parse(serviceConfiguration); err != nil {
		log.Fatalf("couldn't parse configuration from environment: %s", err.Error())
	}

	return serviceConfiguration
}

func startHTTPServer(configuration *config.ServiceConfiguration, shoppingListController *userShoppingList.Controller, shoppingListEntryController *userShoppingListEntry.Controller) {
	// Create client for user service for token validation
	userConn, err := grpc.Dial(configuration.GrpcUserServiceTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer userConn.Close()
	grpcUserServiceClient := user.NewUserServiceClient(userConn)

	authMiddleware := auth.CreateAuthMiddleware(grpcUserServiceClient)
	handler := router.New(shoppingListController, shoppingListEntryController, authMiddleware)
	server := &http.Server{Addr: fmt.Sprintf(":%d", configuration.HttpPort), Handler: handler}

	log.Println("Starting HTTP server: ", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP Server Shutdown Failed:%v", err)
	}
}
