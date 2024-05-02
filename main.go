package main

import (
	"context"
	"fmt"
	application "internal/consumer/application"
	database "internal/consumer/database"
	initializers "internal/consumer/initializers"
	"log"

	_ "github.com/lib/pq"
)

func main(){

	err := initializers.LoadEnv(".env")
	if err != nil{
		log.Fatal(err)
	}

	_,err = database.ConnectToDataBase()
	if err != nil{
		log.Fatal(err)
	}

	app := application.New()
	err = app.Start(context.TODO())
	if err !=nil{
		fmt.Println("failed to start app:", err)
	}

}