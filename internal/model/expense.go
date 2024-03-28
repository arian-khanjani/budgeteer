package model

import "time"

// TODO: give JSON tags

type Expense struct {
	ID        string    `bson:"_id"`
	Timestamp time.Time `bson:"timestamp"`
	Name      string    `bson:"name"`
	Tags      []*Tag    `bson:"tags"`
	Value     string    `bson:"value"`
}

type Tag struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}
