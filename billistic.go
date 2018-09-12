package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	sa := option.WithCredentialsFile("./private/firekey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	bill := getBill()

	log.Print(bill)

	result, err := client.Collection("sampleData").Doc("Bills").Set(context.Background(), bill)

	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result)

	defer client.Close()
}

func getBill() *Bill {
	// what's this shit?
	bill, err := {
		"name": "usbank",
		"blah": "foo",
	}

	return
}
