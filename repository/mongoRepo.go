package repository

import (
	"context"
	"fmt"
	"log"
	"myHttpServer/models"

	"os"
	"time"

	"myHttpServer/utils"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global Mongo Repo with init
var MongoRepo = NewMongoRepo()

type mongoRepo struct {
	client           *mongo.Client
	database         *string
	userCollection   *string
	schemaCollection *string
	dataCollection   *string
	taskCollection   *string
}

// Create a new repository
func NewMongoRepo() *mongoRepo {
	//Get Password from dotenv
	godotenv.Load()
	username := os.Getenv(utils.MONGO_USERNAME)
	password := os.Getenv(utils.MONGO_PASSWORD)
	uri := fmt.Sprintf(os.Getenv(utils.MONGO_URI), username, password)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	//Localhost connection
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return &mongoRepo{
		client:           client,
		database:         utils.StringPointer(os.Getenv(utils.DATABASE)),
		userCollection:   utils.StringPointer(os.Getenv(utils.USER_COLLECTION)),
		schemaCollection: utils.StringPointer(os.Getenv(utils.SCHEMA_COLLECTION)),
		dataCollection:   utils.StringPointer(os.Getenv(utils.DATA_COLLECTION)),
		taskCollection:   utils.StringPointer(os.Getenv(utils.TASK_COLLECTION)),
	}
}

// Get all schemas
func (r *mongoRepo) GetSchemas() []models.Schema {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.schemaCollection)
	//Find all schemas
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(ctx)
	//Create a slice of schemas
	var schemas []models.Schema
	//Iterate through the cursor
	for cur.Next(ctx) {
		var schema models.Schema
		err := cur.Decode(&schema)
		if err != nil {
			log.Println(err)
		}
		schemas = append(schemas, schema)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return schemas
}

// Get Schema by ID
func (r *mongoRepo) GetSchemaByID(id string) *models.Schema {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.schemaCollection)
	//Find the schema
	var schema models.Schema
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectID}}).Decode(&schema)
	if err != nil {
		log.Println(err)
	}
	return &schema
}

// Save a schema
func (r *mongoRepo) SaveSchema(schema *models.Schema) {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.schemaCollection)
	//Insert the schema
	insertedId, err := collection.InsertOne(ctx, schema)
	if err != nil {
		log.Println(err)
	}
	schema.ID = insertedId.InsertedID.(primitive.ObjectID)
}

// Get all users
func (r *mongoRepo) GetUsers() []models.User {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Find all users
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(ctx)
	//Create a slice of users
	var users []models.User
	//Iterate through the cursor
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return users
}

// Get User by ID
func (r *mongoRepo) GetUserByID(id string) *models.User {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Find the user
	var user models.User
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectID}}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &user
}

// Get User by username
func (r *mongoRepo) GetUserByUsername(username string) *models.User {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Find the user
	var user models.User
	err := collection.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		//Warn log
		log.Println(err)
		return nil
	}
	return &user
}

// Save a user
func (r *mongoRepo) SaveUser(user *models.User) {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Insert the user
	insertedId, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
	user.ID = insertedId.InsertedID.(primitive.ObjectID)
}

// Update a user
func (r *mongoRepo) UpdateUser(user *models.User) {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Update the user
	updateResult, err := collection.UpdateByID(ctx, user.ID, bson.M{"$set": user})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// Delete a user
func (r *mongoRepo) DeleteUser(id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.userCollection)
	//Delete the user
	_, err = collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		log.Println(err)
	}
}

// Save Data to mongo
func (r *mongoRepo) SaveData(data *models.Data) {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.dataCollection)
	//Insert the data
	insertedId, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	data.ID = insertedId.InsertedID.(primitive.ObjectID)
}

// Save Task to mongo
func (r *mongoRepo) SaveTask(task *models.Task) {
	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Get the collection
	collection := r.client.Database(*r.database).Collection(*r.taskCollection)
	//Insert the data
	insertedId, err := collection.UpdateOne(ctx, bson.D{{Key: "username", Value: task.Username},{Key:"date", Value: task.Date}},task)
	if err != nil {
		log.Println(err)
	}
	task.ID = insertedId.UpsertedID.(primitive.ObjectID)
}
