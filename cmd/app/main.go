package main

import (
	"log"

	"redis/internal/server"
)

func main() {
	rServer, err := server.New()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer func() {
		err = rServer.Close()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}()

	key := "key1"
	val, err := rServer.Get(key)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("Key %s = %s", key, val)

	err = rServer.Set(key, struct {
		manufacturer, year, model string
		mileage                   int
		price                     float64
	}{
		manufacturer: "Toyota",
		model:        "Camry",
		mileage:      200000,
		year:         "2018",
		price:        24380.00,
	})
	if err != nil {
		log.Printf("Error: %s", err)
	}

	val, err = rServer.Get(key)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("Key %s = %s", key, val)
}
