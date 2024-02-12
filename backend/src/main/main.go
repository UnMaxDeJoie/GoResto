package main

import (
	"log"
	"net/http"
	"time"

	"NewGoResto/src/entities"
	"NewGoResto/src/handler"
	"NewGoResto/src/managers"
	// "github.com/Watsuk/go-food/src/entity"
	// "github.com/Watsuk/go-food/src/handler"
)

func main() {
	time.Sleep(5 * time.Second)
	db := managers.GetDBController()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var reference entities.Reference
	reference.User = &entities.User{}
	reference.Truck = &entities.Truck{}
	reference.Chart = &entities.Chart{}
	reference.Order = &entities.Order{}
	reference.OrderDetail = &entities.OrderDetail{}

	mux := handler.NewHandler(db, reference)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("could not listen on port 3000: %v", err)
		return
	}

}
