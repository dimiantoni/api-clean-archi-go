package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID entity ID
type ID = primitive.ObjectID

// NewID create a new entity ID
func NewID() ID {
	return ID(primitive.NewObjectID())
}

// StringToID convert a string to an entity ID
func StringToID(s string) ID {
	c, _ := primitive.ObjectIDFromHex(s)
	return c
}
