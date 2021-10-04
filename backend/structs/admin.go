package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AdminID  string             `bson:"admin_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
