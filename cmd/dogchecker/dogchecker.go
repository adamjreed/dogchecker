package main

import (
	"context"
	"dogchecker/internal/app"
	"dogchecker/internal/handlers"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	appContext, err := app.CreateContext()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("booting...")

	ctx := context.WithValue(context.Background(), app.AppContext, appContext)

	lambda.StartWithContext(ctx, handlers.CheckForDogs)
}
