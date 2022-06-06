package database

import (
	"log"
)

func getUserHistory(userId int) {

	db, err := createConnection()
	if err != nil {
		log.Fatalf("Error while connecting to daabase: %v", err)
	}
	rows, err := db.Query("SELECT FROM public.image WHERE (user = $1)", userId)
	if err != nil {
		log.Fatalf("Error while inserting image: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var url string
		var pk, user int
		if err := rows.Scan(&pk, &user, &url); err != nil {
			log.Fatalf("Error while scanning user history rows: %v", err)
			break
		}
		// TODO - add data to proto msg response
	}
}
