package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"module31/internal/entity"
	"strconv"
)

var collection *mongo.Collection
var ctx = context.TODO()

type mongodb struct {
	index     int
	usersById map[int]*entity.User
}

func NewMongodb() (*mongodb, error) {
	return &mongodb{
		usersById: make(map[int]*entity.User),
	}, nil
}

func connectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func disconnectDB(client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

//CreateUser accepts new user, adds to the database and return user id
func (r *mongodb) CreateUser(user *entity.User) (int, error) {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")

	var u *entity.User
	opt := options.FindOne().SetSort(bson.D{{"_id", -1}})
	err := collection.FindOne(
		ctx,
		bson.D{{}},
		opt,
	).Decode(&u)
	if err != nil {
		user.Id = 1
		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err)
		}
		disconnectDB(client)
		return user.Id, nil
	}
	r.index = u.Id
	r.index++
	user.Id = r.index
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	disconnectDB(client)
	return user.Id, nil
}

//DeleteUser accepts user id, delete from database and return user name
func (r *mongodb) DeleteUser(id int) (string, error) {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")
	var user *entity.User
	cur := collection.FindOne(ctx, bson.D{{"_id", id}})
	err := cur.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{"Friends", strconv.Itoa(id)}}

	update := bson.D{
		{"$pull", bson.D{
			{"Friends", strconv.Itoa(id)},
		}},
	}
	_, err = collection.UpdateMany(ctx, filter, update)

	_, err = collection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
	}
	disconnectDB(client)
	return user.Name, nil
}

//GetUsers return all users from database
func (r *mongodb) GetUsers() map[int]*entity.User {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var user *entity.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		r.usersById[user.Id] = user
	}
	err = cur.Close(ctx)
	disconnectDB(client)
	return r.usersById
}

//UpdateAge accepts user id and new age, update user age into database
func (r *mongodb) UpdateAge(id int, newAge int) error {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$set", bson.D{
			{"Age", newAge},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	disconnectDB(client)
	return nil
}

//MakeFriends accepts target and source id, adds to the slice of friends each other and returns users names
func (r *mongodb) MakeFriends(target int, source int) (string, string, error) {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")
	cur, _ := collection.Find(ctx, bson.D{{
		"_id",
		bson.D{{
			"$in",
			bson.A{target, source},
		}},
	}})
	var user *entity.User
	var u1 string
	var u2 string
	for cur.Next(ctx) {
		_ = cur.Decode(&user)
		if user.Id == target {
			u1 = user.Name
			filter := bson.D{{"_id", target}}

			update := bson.D{
				{"$push", bson.D{
					{"Friends", strconv.Itoa(source)},
				}},
			}
			_, _ = collection.UpdateOne(ctx, filter, update)
		} else if user.Id == source {
			u2 = user.Name
			filter := bson.D{{"_id", source}}

			update := bson.D{
				{"$push", bson.D{
					{"Friends", strconv.Itoa(target)},
				}},
			}
			_, _ = collection.UpdateOne(ctx, filter, update)
		}
	}

	disconnectDB(client)
	return u1, u2, nil
}

//GetFriends adds user id, return slice of friends names
func (r *mongodb) GetFriends(userId int) ([]string, error) {
	client := connectDB()
	collection = client.Database("usersDB").Collection("users")

	var user *entity.User

	f, err := collection.Find(ctx, bson.D{{"Friends", strconv.Itoa(userId)}})
	if err != nil {
		log.Fatal(err)
	}
	var friends []string
	for f.Next(ctx) {
		_ = f.Decode(&user)
		friends = append(friends, user.Name)
	}
	disconnectDB(client)
	return friends, nil
}
