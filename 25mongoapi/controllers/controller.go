package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapis/models"
	"net/http"

	"github.com/gorilla/mux"                     // For handling HTTP routes
	"go.mongodb.org/mongo-driver/bson"           // For working with MongoDB documents
	"go.mongodb.org/mongo-driver/bson/primitive" // For MongoDB ObjectIDs
	"go.mongodb.org/mongo-driver/mongo"          // MongoDB driver
	"go.mongodb.org/mongo-driver/mongo/options"  // MongoDB connection options
)

// MongoDB connection details
const connectionString = "mongodb://localhost:27017/" // Address of the MongoDB server
const dbName = "firstgo"                              // Database name
const colName = "watchlist"                           // Collection name (like a table in SQL)

var collection *mongo.Collection // Global variable to hold our collection reference

// DATABASE CONNECTION SETUP
// --------------------------

// init runs automatically when the app starts, similar to a constructor in other languages
func init() {
	// Set up the connection options for MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// If connection fails, stop the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connection success")

	// Get a reference to our specific collection
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection reference is ready")
}

// DATABASE HELPER FUNCTIONS
// -------------------------

// insertOneMovie - Adds a new movie to the database
func insertOneMovie(movie models.Netflix) {
	// Insert the movie document into MongoDB
	inserted, err := collection.InsertOne(context.Background(), movie)

	// If insertion fails, stop the program
	if err != nil {
		log.Fatal(err)
	}

	// Print the ID of the newly inserted document
	fmt.Println(inserted.InsertedID)
}

// updateOneMovie - Marks a movie as watched by its ID
func updateOneMovie(movieId string) {
	// Convert string ID to MongoDB ObjectID
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	// Set up filter to find the movie by ID
	filter := bson.M{"_id": id}

	// Set up update to change "watched" field to true
	update := bson.M{"$set": bson.M{"watched": true}}

	// Update the document in MongoDB
	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	// Print ID of the document if it was newly inserted (which doesn't happen in this case)
	fmt.Println(updated.UpsertedID)
}

// deleteOneMovie - Removes a single movie by its ID
func deleteOneMovie(movieId string) {
	// Convert string ID to MongoDB ObjectID
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	// Set up filter to find the movie by ID
	filter := bson.M{"_id": id}

	// Delete the document from MongoDB
	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	// Print how many documents were deleted (should be 1)
	fmt.Println(deleted.DeletedCount)
}

// deleteAllMovie - Clears all movies from the collection
func deleteAllMovie() {
	// Empty filter {} means "match all documents"
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Print how many documents were deleted
	fmt.Println(deleted.DeletedCount)
}

// getAllMovies - Retrieves all movies from the database
func getAllMovies() []primitive.M {
	// Find all documents in the collection (empty filter {})
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice to hold all the movies
	var movies []primitive.M

	// Loop through each document in the result
	for cursor.Next(context.Background()) {
		var movie bson.M
		// Convert the document to a map
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		// Add the movie to our slice
		movies = append(movies, movie)
	}

	// Close the cursor when we're done
	defer cursor.Close(context.Background())
	return movies
}

// HTTP HANDLER FUNCTIONS
// ----------------------

// GetMyAllMovies - HTTP handler to get all movies
// Called when a GET request is made to the endpoint
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	// Set the response content type
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	// Get all movies from database
	allMovies := getAllMovies()

	// Send the movies as JSON in the response
	json.NewEncoder(w).Encode(allMovies)
}

// CreateMovie - HTTP handler to add a new movie
// Called when a POST request is made to the endpoint
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// Create an empty movie struct
	movie := models.Netflix{}

	// Parse the JSON from the request body into our movie struct
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Save the movie to database
	insertOneMovie(movie)

	// Send the movie back as JSON in the response
	json.NewEncoder(w).Encode(movie)
}

// MarkAsWatched - HTTP handler to update a movie's watched status
// Called when a PUT request is made to the endpoint
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// Get the movie ID from URL parameters
	params := mux.Vars(r)

	// Update the movie in database
	updateOneMovie(params["id"])

	// Send the movie ID back as JSON in the response
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteOneMovie - HTTP handler to remove a single movie
// Called when a DELETE request is made to the endpoint
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// Get the movie ID from URL parameters
	params := mux.Vars(r)

	// Delete the movie from database
	deleteOneMovie(params["id"])

	// Send the deleted movie ID back as JSON in the response
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteAllMovies - HTTP handler to clear all movies
// Called when a DELETE request is made to the endpoint
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// Delete all movies from database
	deleteAllMovie()

	// Send confirmation message as JSON in the response
	json.NewEncoder(w).Encode("allMovies are deleted")
}
