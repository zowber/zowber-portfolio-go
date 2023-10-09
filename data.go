package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func DotEnv(key string) string {

	if os.Getenv(key) == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	return os.Getenv(key)
}

func newPortfolioDbClient() (*MongoDBClient, error) {

	connectionStr := DotEnv("DB_URI")
	dbName := "portfolioitems"
	collectionName := "casestudies"

	ctx := context.Background()
	logLvl := options.LogLevel(5)
	loggerOpts := options.Logger().SetComponentLevel(options.LogComponentAll, logLvl)
	clientOpts := options.
		Client().
		ApplyURI(connectionStr).
		SetLoggerOptions(loggerOpts)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Print(err)
		// TODO: Recover from this error
		panic(err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoDBClient{client, collection}, err
}

func (m *MongoDBClient) getAllCaseStudies() ([]CaseStudy, error) {
	ctx := context.Background()
	filter := bson.M{}
	opts := options.Find()
	opts.SetSort(bson.M{"id": -1})

	var caseStudies []CaseStudy

	cursor, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println(err)
		return caseStudies, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &caseStudies); err != nil {
		log.Println(err)
		return caseStudies, err
	}

	return caseStudies, err
}

func (m *MongoDBClient) getCaseStudy(caseStudyId int) (CaseStudy, error) {
	ctx := context.Background()
	filter := bson.M{"id": caseStudyId}

	var caseStudy CaseStudy

	err := m.collection.FindOne(ctx, filter).Decode(&caseStudy)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			return CaseStudy{}, err
		}
	}

	return caseStudy, err
}
