package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	model "github.com/kmabhinai/go-mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = os.Getenv("DBURL")
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db connected")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

//Mongo Db helpers

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single movie: %v\n", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie Got deleted", deleteCount)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies Deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual Controllers

func GetMyAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(res).Encode(allMovies)
}

func CreateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(req.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(res).Encode(movie)
}

func MarkAsWatched(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(req)

	updateOneMovie(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DeleteAMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(req)
	deleteOneMovie(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DeleteAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result := deleteAllMovies()
	json.NewEncoder(res).Encode(result)
}
