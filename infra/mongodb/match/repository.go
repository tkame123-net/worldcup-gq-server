package match

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
	"time"
	"tkame123-net/worldcup-gq-server/adapter"
	"tkame123-net/worldcup-gq-server/domain"
)

type repository struct {
	client adapter.MongoClient
}

func NewRepository(client adapter.MongoClient) adapter.MongodbMatchRepository {
	return &repository{
		client,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.Match, error) {
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

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var i entity
		err := cur.Decode(&i)
		if err != nil {
			return nil, fmt.Errorf("message: %w", err)
		}
		entities = append(entities, i)
	}

	items := make([]*domain.Match, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Match/GetAll")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}

func (r *repository) GetAllByYear(ctx context.Context, year string, filterType domain.FilterType) ([]*domain.Match, error) {
	var entities []entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	//cur, err := col.Find(ctx, bson.M{})

	var match primitive.M

	switch filterType {
	case domain.FilterType_EQ:
		_year, _ := strconv.Atoi(year)
		match = bson.M{"Year": _year}
	case domain.FilterType_REGEX:
		//_year, _ := strconv.Atoi(year)
		selector := primitive.Regex{Pattern: year}
		match = bson.M{"Year": selector}
	}

	cur, err := col.Find(ctx, match)

	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var i entity
		err := cur.Decode(&i)
		if err != nil {
			return nil, fmt.Errorf("message: %w", err)
		}
		entities = append(entities, i)
	}

	items := make([]*domain.Match, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Match/GetAll")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}
