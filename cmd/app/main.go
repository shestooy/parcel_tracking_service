package main

import (
	"log"
	a "parcel_tracking_service/internal/app"
)

func main() {
	err := a.Start()
	if err != nil {
		log.Fatal(err)
	}
}
