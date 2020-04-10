# mongo-driver-autoincrement (ai)
Package **ai** implements AutoIncrement methods for mongo-driver(golang) 

## How To Install

```
go get github.com/autotrend/mongo-ai
```

## Getting Started

```go
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ai "github.com/autotrend/mongo-ai"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("example-db").Collection("counters")

	ai := ai.Create(collection)

    client.Database("example-db").Collection("users").InsertOne(ctx, bson.M{
		"_id":   ai.Next("sequenc"),
		"login": "test",
		"age":   32,
	})
}

```

## License
           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                   Version 1, Abril 2020

Copyright (C) 2020 Dennys Freire <dennysvf@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
  TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
