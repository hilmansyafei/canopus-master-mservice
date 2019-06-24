package api

import "gopkg.in/mgo.v2"

// Handler connection database
type Handler struct {
	DB *mgo.Database
}
