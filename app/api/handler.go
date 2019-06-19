package api

import "github.com/zebresel-com/mongodm"

// Handler connection database
type Handler struct {
	DB *mongodm.Connection
}
