package DAL

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	FName string
	LName string
}

func AddUser(user User) {

	session := Connect()

	collection := session.Database("demo").Collection("usersdemo")

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	log.Println(result)
	log.Println(err)

}

func GetUser() []User {

	var users []User

	session := Connect()

	collection := session.Database("demo").Collection("usersdemo")
	cur, _ := collection.Find(context.TODO(), bson.M{})

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var user User

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	return users

}

func UpdateUserByID(_id string, user User) {

	session := Connect()

	collection := session.Database("demo").Collection("usersdemo")

	idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"_id": idPrimitive}
	update := bson.M{"$set": bson.M{"fname": user.FName, "lname": user.LName}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	fmt.Println(result, err)

}

func DeleteUserByID(_id string) {

	session := Connect()

	collection := session.Database("demo").Collection("usersdemo")

	idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"_id": idPrimitive}

	result, err := collection.DeleteOne(context.TODO(), filter)

	fmt.Println(result, err)

}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected...")
	return client
}
