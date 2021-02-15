package competition

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *repository) GetByYear(ctx context.Context, year *int) (*domain.Competition, error) {
	var entity entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	e := col.FindOne(ctx, bson.M{"Year": year}).Decode(&entity)
	if e != nil {
		// todo: エラーハンドリング（Nodocだけ分岐させたい）
		log.Printf("log: %v", e)
		return nil, nil
	}

	log.Println("[info] infra/mongodb/Competition/GetByYear")
	log.Println("[info] ", *year)
	log.Println("[info] ", entity)

	return entity.toDomain(), nil
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

func (r *repository) Get(ctx context.Context, id domain.CompetitionID) (*domain.Competition, error) {
	var entity entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	objectID, _ := primitive.ObjectIDFromHex(string(id))

	e := col.FindOne(ctx, bson.M{"_id": objectID}).Decode(&entity)
	if e != nil {
		// todo: エラーハンドリング（Nodocだけ分岐させたい）
		log.Printf("log: %v", e)
		return nil, nil
	}

	log.Println("[info] infra/mongodb/Competition/Get")
	log.Println("[info] ", id)
	log.Println("[info] ", entity)

	return entity.toDomain(), nil
}

func (r *repository) GetCursorsToEdges(ctx context.Context, after *string, before *string) ([]*domain.Competition, error) {
	var entities []entity
	var filter bson.M

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	if after == nil && before == nil {
		filter = bson.M{}
	}
	if after != nil && before != nil {
		objectIDAfter, _ := primitive.ObjectIDFromHex(*after)
		objectIDBefore, _ := primitive.ObjectIDFromHex(*before)
		filter = bson.M{"_id": bson.M{"$gt": objectIDAfter, "$lt": objectIDBefore}}
	}
	if after != nil && before == nil {
		objectIDAfter, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": objectIDAfter}}
	}
	if after == nil && before != nil {
		objectIDBefore, _ := primitive.ObjectIDFromHex(*before)
		filter = bson.M{"_id": bson.M{"$lte": objectIDBefore}}
	}
	opts := options.FindOptions{
		Sort: bson.M{"_id": 1},
	}
	cur, err := col.Find(ctx, filter, &opts)
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

	items := make([]*domain.Competition, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Competition/GetMultiByRange")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}

func (r *repository) GetMultiByRange(ctx context.Context, limit *int, cursor *string, asc *bool) ([]*domain.Competition, error) {
	var entities []entity
	sort := 1 // default: asc
	if *asc == false {
		sort = -1
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	objectID, _ := primitive.ObjectIDFromHex(*cursor)
	filter := bson.M{"_id": bson.M{"$gte": objectID}}
	if sort == -1 {
		filter = bson.M{"_id": bson.M{"$lte": objectID}}
	}
	i := int64(*limit)

	opts := options.FindOptions{
		Sort:  bson.M{"_id": sort},
		Limit: &i,
	}
	cur, err := col.Find(ctx, filter, &opts)
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

	items := make([]*domain.Competition, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Competition/GetMultiByRange")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}

func (r *repository) GetAllByCountry(ctx context.Context, country *string) ([]*domain.Competition, error) {

	var entities []entity

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)

	t := *country

	selector := primitive.Regex{Pattern: t}

	cur, err := col.Find(ctx, bson.M{"Country": selector})
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

	items := make([]*domain.Competition, 0, len(entities))
	for _, v := range entities {
		items = append(items, v.toDomain())
	}

	log.Println("[info] infra/mongodb/Competition/GetAllByCountry")
	for _, i := range items {
		log.Println("[info] ", i)
	}

	return items, nil
}

func (r *repository) Exists(ctx context.Context, id *domain.CompetitionID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	c, err := r.client(ctx)
	if err != nil {
		return false, fmt.Errorf("message: %w", err)
	}
	defer c.Client.Disconnect(ctx)
	col := c.Client.Database(c.Database).Collection(collection)
	objectID, _ := primitive.ObjectIDFromHex(string(*id))

	i, err := col.CountDocuments(ctx, bson.M{"_id": objectID})
	if err != nil {
		return false, fmt.Errorf("message: %w", err)
	}

	log.Println("[info] infra/mongodb/Competition/Exists")
	log.Println("[info] ", id)
	log.Println("[info] ", i)

	return i > 0, nil
}
