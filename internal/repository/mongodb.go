package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"module31/internal/entity"
)

type mongodb struct {
	client *mongo.Client
}

func NewMongodb() (*mongodb, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return &mongodb{client: client}, nil

}

func DisconnectDB(client *mongodb) {
	err := client.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}

//CreateUser accepts new user, adds to the database and return user id
func (r *mongodb) CreateUser(user *entity.User) (string, error) {

	collection := r.client.Database("usersDB").Collection("users")

	u, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	id := u.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

//DeleteUser accepts user id, delete from database and return user name
func (r *mongodb) DeleteUser(id string) (string, error) {

	collection := r.client.Database("usersDB").Collection("users")

	userID, err := primitive.ObjectIDFromHex(id)

	var d bson.M
	_ = collection.FindOneAndDelete(context.TODO(), bson.D{{
		"_id",
		userID,
	}}).Decode(&d)

	name := d["Name"].(string)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{"Friends", id}}
	fmt.Println(name)
	update := bson.D{
		{"$pull", bson.D{
			{"Friends", id},
		}},
	}
	_, err = collection.UpdateMany(context.TODO(), filter, update)

	return name, nil
}

//GetUsers return all users from database
func (r *mongodb) GetUsers(user *entity.User) []*entity.User {

	collection := r.client.Database("usersDB").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var allUsers []*entity.User
	for cur.Next(context.TODO()) {

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		allUsers = append(allUsers, user)
	}
	err = cur.Close(context.TODO())

	return allUsers
}

//UpdateAge accepts user id and new age, update user age into database
func (r *mongodb) UpdateAge(id string, newAge int) error {

	collection := r.client.Database("usersDB").Collection("users")
	userID, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", userID}}

	update := bson.D{
		{"$set", bson.D{
			{"Age", newAge},
		}},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

//MakeFriends accepts target and source id, adds to the slice of friends each other and returns users names
func (r *mongodb) MakeFriends(target string, source string) (string, string, error) {

	collection := r.client.Database("usersDB").Collection("users")
	targetID, _ := primitive.ObjectIDFromHex(target)
	sourceID, _ := primitive.ObjectIDFromHex(source)
	opt := bson.D{
		{"_id", 0},
		{"Name", 1},
	}
	cur, _ := collection.Find(context.TODO(), bson.D{{
		"_id",
		bson.D{{
			"$in",
			bson.A{targetID, sourceID},
		}},
	}}, options.Find().SetProjection(opt))
	var n bson.M
	var names []string

	for cur.Next(context.TODO()) {
		_ = cur.Decode(&n)
		names = append(names, n["Name"].(string))
	}

	filter := bson.D{{"_id", targetID}}

	update := bson.D{
		{"$push", bson.D{
			{"Friends", source},
		}},
	}
	_, _ = collection.UpdateOne(context.TODO(), filter, update)

	filter = bson.D{{"_id", sourceID}}

	update = bson.D{
		{"$push", bson.D{
			{"Friends", target},
		}},
	}
	_, _ = collection.UpdateOne(context.TODO(), filter, update)

	return names[0], names[1], nil
}

//GetFriends accepts user id, return slice of friends names
func (r *mongodb) GetFriends(userId string) ([]string, error) {

	collection := r.client.Database("usersDB").Collection("users")

	var user bson.M

	f, err := collection.Find(context.TODO(), bson.D{{"Friends", userId}})
	if err != nil {
		log.Fatal(err)
	}
	var friends []string
	for f.Next(context.TODO()) {
		_ = f.Decode(&user)
		friends = append(friends, user["Name"].(string))
	}

	return friends, nil
}
