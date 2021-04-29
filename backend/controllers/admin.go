package controllers

import (
	"backend/db"
	"database/sql"
)

// AdminController is a controller and is defined here.
type AdminController struct {
	DB      *sql.DB
	Queries *db.Queries
}

// NewAdminController returns pointer to AdminController.
func NewAdminController(db *sql.DB, queries *db.Queries) *AdminController {
	return &AdminController{
		DB:      db,
		Queries: queries,
	}
}
