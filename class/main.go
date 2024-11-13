package main

import (
	"20241113/class/database"
	"20241113/class/router"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen("20241113")
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
