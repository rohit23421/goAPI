package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohit23421/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB connection string
const connectionString = "mongodb+srv://rohit23421:rohit92929@cluster0.h3nggix.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//taking a reference of the mongoDB collection here
var collection *mongo.Collection

//connect with mongoDB
func init() {
	//init runs only once in Golang
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//firing a connection request by connecting to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB conection success")

	collection = client.Database(dbName).Collection(colName)

	//if collection instance/reference is ready
	fmt.Println("Collection instance/reference is ready!!!")
}

//mongodb helpers

// insert one record into DB
// the passed data will be a movie variable of model.Netflix type
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one record in movie DB with id: ", inserted.InsertedID)
}

//update a record in DB
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	//filtering using mongoDB bson method
	filter := bson.M{"_id": id}
	//providing the update data to the model
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)

}

//delete a record in DB
//in mongodb we have to provide a filter so that go search for this filter and
// delete the record that matches
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record successfully Deleted with deletecount: ", deleteCount)

}

//delete all records from mongodb
func deleteAllMovie() int64 {
	filter := bson.D{{}} // not passing any value, means everything is selected
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All records deleted successfully, count ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all movie records from DB
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
		// if record present we append it into the moviews primitive.M we created
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

//controllers
func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "appliation/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

//creating a movie into DB, this is controller
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

//marking movie as watched, this is controller
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	fmt.Println("Reached inside markaswatched")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//deleteing one movie from DB, this is controller
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//deleting all movies from DB using this controller
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
