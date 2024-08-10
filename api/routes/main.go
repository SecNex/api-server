package routes

import "github.com/secnex/server/db"

type Router struct {
	Database *db.Connection
}

func NewRouter(cnx *db.Connection) *Router {
	return &Router{
		Database: cnx,
	}
}
