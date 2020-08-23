package competition

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
	"tkame123-net/worldcup-gq-server/adapter"
	"tkame123-net/worldcup-gq-server/domain"
)

type repository struct {
	client adapter.MongoClient
}

func NewRepository(client adapter.MongoClient) adapter.MongodbCompetitionRepository {
	return &repository{
		client,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.Competition, error) {
	var entities []entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}

	log.Printf("col: %s", collection)

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var i entity
		err := cur.Decode(&i)
		if err != nil {
			return nil, fmt.Errorf("message: %w", err)
		}
		entities = append(entities, i)
	}

	items := make([]*domain.Competition, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Competition/GetAll")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}
