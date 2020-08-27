package player

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"tkame123-net/worldcup-gq-server/adapter"
	"tkame123-net/worldcup-gq-server/domain"
)

type repository struct {
	client adapter.MongoClient
}

func NewRepository(client adapter.MongoClient) adapter.MongodbPlayerRepository {
	return &repository{
		client,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.Player, error) {
	var entities []entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	groupStage := bson.D{
		{"$group", bson.M{"_id": "$Player Name", "MatchIDList": bson.M{"$push": "$MatchID"}}},
	}
	lookupStage := bson.D{
		{"$lookup", bson.M{"from": "matches", "localField": "MatchIDList", "foreignField": "MatchID", "as": "MatchList"}},
	}
	cur, err := col.Aggregate(ctx, mongo.Pipeline{groupStage, lookupStage})
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		log.Printf("cur: %v", cur)
		var i entity
		err := cur.Decode(&i)
		if err != nil {
			return nil, fmt.Errorf("message: %w", err)
		}
		entities = append(entities, i)
	}

	items := make([]*domain.Player, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Player/GetAll")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}
