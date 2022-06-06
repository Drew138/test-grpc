package database

import (
	"log"
)

func insertImage(userId int, url string) {
	db, err := createConnection()
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}
	_, err = db.Exec("INSERT INTO public.image (user, url) VALUES ($1, $2)", userId, url)
	if err != nil {
		log.Fatalf("Error while inserting image: %v", err)
	}
}

func insertUser(userId int, provider string) {
	db, err := createConnection()
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}
	_, err = db.Exec("INSERT INTO public.user (user, provider) VALUES ($1, $2)", userId, provider)
	if err != nil {
		log.Fatalf("Error while inserting image: %v", err)
	}
}
