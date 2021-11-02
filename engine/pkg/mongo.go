package pkg

import (
	"context"
	"fmt"
	"github.com/kuno989/cert_plugin"
	"github.com/kuno989/cert_plugin/engine/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongo struct {
	Config cert_plugin.Config
	client *mongo.Client
}
func NewMongo(ctx context.Context, cfg cert_plugin.Config) (*Mongo, error){
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		return nil, err
	}
	ctxTime, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := client.Connect(ctxTime); err != nil {
		return nil, fmt.Errorf("connect mongo: %w", err)
	}

	return &Mongo{
		Config: cfg,
		client: client,
	}, nil
}

func (m *Mongo) ApiKeyCheck(ctx context.Context, apiKey string) (schema.Key, error){
	coll := m.client.Database(m.Config.DB).Collection(m.Config.Collection)
	var key schema.Key
	err := coll.FindOne(ctx, bson.M{"api_key": apiKey}).Decode(&key)
	return key, err
}