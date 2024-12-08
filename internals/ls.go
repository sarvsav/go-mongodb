package internals

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/sarvsav/go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OptionsLsFunc func(c *models.LsOptions) error

func Ls(lsOptions ...OptionsLsFunc) error {
	mongodbURI := os.Getenv("MONGODB_URI")
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
		Logger:      slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return err
		}
	}

	lsCmd.Logger.Debug("provided command with options", "longListing", lsCmd.LongListing, "color", lsCmd.Color, "args", lsCmd.Args)
	// List all the databases in the cluster running on connected server

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Use a filter to only select non-empty databases.
	result, err := client.ListDatabaseNames(
		context.TODO(),
		bson.D{},
	//	bson.D{{"empty", false}}
	)
	if err != nil {
		log.Panic(err)
	}

	for _, db := range result {
		database := client.Database(db)
		collections, err := database.ListCollectionNames(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Collections:")
		for _, collection := range collections {
			fmt.Println(collection)
		}

	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	return nil
	// err = client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
