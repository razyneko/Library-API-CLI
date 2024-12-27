package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Library-API-CLI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Declare global variables for mongo client and mongo collection
var client *mongo.Client
var bookCollection *mongo.Collection

// setting mongoURL
var mongoURL = "mongodb://localhost:27017"

// Initialising mongoDB connection
func init() {
	//set client options
	clientOptions := options.Client().ApplyURI(mongoURL)

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	bookCollection = client.Database("library").Collection("books")
}

// Handler Functions

// Get All Books
func GetAllBooks(c *gin.Context) {
	// for returning books use a slice of type Book struct
	// i must initialise it even if its empty because serialization keeps [] -> []
	// but if i dont initialise it it is nil after serialisation it became null 
	var books = []models.Book{}
	

	// find all books in mongodb
	// for getting all entries use bson.D{{"empty -- here"}}
	cursor, err := bookCollection.Find(c, bson.M{})
	// can also pass context.Bacgtound() in place of c
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error while searching for books"})
		return
	}

	// for ensuring that db is closed after the request
	defer cursor.Close(c)

	// loop through the cursor to decode each document into books slice
	// for cursor.Next(context.Background()) {
	// 	var book models.Book
	// 	if err := cursor.Decode(&book); err != nil {
	// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	// appending decoded document to books slice
	// 	books = append(books, book)
	// }

	// much simpler way
	if err := cursor.All(c, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error decoding books"})
		return
	}
	// Return books slice
	c.IndentedJSON(http.StatusOK, books)
}

// Get Book By Id

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}

	var book models.Book

	if err := bookCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&book); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// Add a new book to library
func AddBook(c *gin.Context) {
	var newBook models.Book

	// converting json to struct data
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// creating new id to store in mongoDB
	newBook.ID = primitive.NewObjectID()

	// Inserting into DB
	_, err := bookCollection.InsertOne(c, newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error while inerting into DB"})
		return
	}
	c.IndentedJSON(http.StatusOK, newBook)
}

// checkout a book
func CheckoutBook(c *gin.Context) {
	id := c.Param("id")
	// check for missing id
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing ID parameter"})
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	// find and update the book in the database
	filter := bson.M{"_id": objectID, "stock": bson.M{"$gt": 0}} // searching for stock greater than 0
	update := bson.M{"$inc": bson.M{"stock": -1}}                // increment by -1 = decrement by 1

	// finding and updating according to filter

	result := bookCollection.FindOneAndUpdate(c, filter, update)

	if result.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Book is out of stock :("})
	}

	var updatedBook models.Book
	if err := result.Decode(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"Error in updating book stock"})
		return
	}
	c.IndentedJSON(http.StatusOK, "Book checked out successfully")
}

// Return Handler

func ReturnBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing Id Parameter"})
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	// find and update the book in db
	filter := bson.M{"_id": objectID}
	update := bson.M{"$inc": bson.M{"stock": 1}}

	result := bookCollection.FindOneAndUpdate(c, filter, update)

	if result.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Book Not Found"})
		return
	}

	var updatedBook models.Book

	if err := result.Decode(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}
	c.IndentedJSON(http.StatusOK, "Book returned successfully")
}
