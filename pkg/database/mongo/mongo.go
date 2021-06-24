package mongo

import (
	"context"
	"fmt"
	"github.com/MuZaZaVr/account-service/internal/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const timeout = 10

func NewMongo(ctx context.Context, cfg config.MongoConfig) (*mongo.Database, error) {
	cfg.URI = fmt.Sprintf("%s://%s:%v", cfg.Dialect, cfg.Host, cfg.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, errors.Wrap(err, "can't create new mongo client")
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = checkConnection(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "cant connect to database")
	}

	return client.Database(cfg.Name), nil
}

func checkConnection(ctx context.Context, client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "can't ping to database")
	}

	return nil
}
