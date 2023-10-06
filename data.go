package main

import (
	"context"
	"fmt"
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

func goDotEnv(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	return os.Getenv(key)
}

func newMongoDBClient() (*MongoDBClient, error) {

	connectionStr := goDotEnv("DB_URI")
	dbName := "portfolioitems"
	collectionName := "casestudies"

	ctx := context.Background()
	opts := options.Client().ApplyURI(connectionStr)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoDBClient{client, collection}, err

}

func (m *MongoDBClient) getAllCaseStudies() ([]CaseStudy, error) {
	ctx := context.Background()
	filter := bson.M{}

	var caseStudies []CaseStudy

	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return caseStudies, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &caseStudies); err != nil {
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
			log.Println("No case study found matching filter:", filter)
			return CaseStudy{}, err
		}
	}

	fmt.Println("Found a case study with name:", caseStudy.Name)
	return caseStudy, nil
}
