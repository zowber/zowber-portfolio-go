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

func goDotEnv(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	return os.Getenv(key)
}

func getCaseStudy(caseStudyId int) (CaseStudy, error) {

	opts := options.Client().ApplyURI(goDotEnv("DB_URI"))
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("portfolioitems").Collection("casestudies")
	query := bson.M{"id": caseStudyId}
	options := options.FindOne()

	var caseStudy CaseStudy

	if err := collection.FindOne(context.Background(), query, options).Decode(&caseStudy); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("getCaseStudy(caseStudyId int): No documents for case study with id %d\n", caseStudyId)
			return caseStudy, err
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("getCaseStudy(caseStudyId int): Got case study with name %s\n", caseStudy.Name)

	return caseStudy, err
}

func getAllCaseStudies() ([]CaseStudy, error) {

	opts := options.Client().ApplyURI(goDotEnv("DB_URI"))

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		fmt.Printf("getAllCaseStudies(): %s\n", err)
		panic(nil)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			fmt.Printf("getAllCaseStudies(): %s\n", err)
			log.Fatal(err)
		}
	}()

	collection := client.Database("portfolioitems").Collection("casestudies")
	query := bson.M{}
	options := options.Find()

	sortFilter := bson.D{{Key: "id", Value: -1}}
	options.SetSort(sortFilter)

	cursor, err := collection.Find(context.Background(), query, options)
	if err != nil {
		fmt.Printf("getAllCaseStudies(): %s\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var caseStudies []CaseStudy
	for cursor.Next(context.Background()) {
		var caseStudy CaseStudy
		if err := cursor.Decode(&caseStudy); err != nil {
			fmt.Printf("getAllCaseStudies(): %s\n", err)
			return nil, err
		}
		caseStudies = append(caseStudies, caseStudy)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("getAllCaseStudies(): %s\n", err)
		return nil, err
	}

	fmt.Printf("getAllCaseStudies(): Got case studies\n")

	return caseStudies, err
}
