package mongodb

import (
	"budgeteer/internal/model"
	"budgeteer/internal/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func New(ctx context.Context) (r *Repository, err error) {
	uri := util.GetEnv("MONGODB_URI", false, "", true)
	dbName := util.GetEnv("MONGODB_DATABASE", false, "", true)
	collName := util.GetEnv("MONGODB_COLLECTION", false, "", true)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	r.client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = r.client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	r.database = r.client.Database(dbName)

	var result bson.M
	if err = r.database.RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}

	r.collection = r.database.Collection(collName)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!") // TODO: USE A LOGGER

	return r, nil
}

func (r *Repository) Disconnect(ctx context.Context) {
	if err := r.client.Disconnect(ctx); err != nil {
		fmt.Println(err) // TODO: USE A LOGGER
	}
}

func (r *Repository) GetExpenses() ([]*model.Expense, error) {

	return nil, nil
}
