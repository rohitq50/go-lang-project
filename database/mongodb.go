package database

import (
	"context"
	"time"

	"github.com/rohitq50/go-lang-project/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *mongo.Client
var err error

func GetInstance(c *config.Configuration) *mongo.Client {
	if mongoInstance == nil {
		// ctx will be used to set deadline for process, here
		// deadline will of 30 seconds.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		mongoInstance, err = mongo.Connect(ctx, options.Client().ApplyURI(c.DBHost))
		if err != nil {
			panic(err)
		}

	}
	return mongoInstance
}
