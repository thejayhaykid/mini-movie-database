package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie with all of the relevant information from tmdb.com that I need
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}
