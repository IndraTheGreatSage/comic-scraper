package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Comics   *mongo.Collection
	Chapters *mongo.Collection
}

func NewDatabase(uri, dbName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &Database{
		Client:   client,
		Comics:   db.Collection("comics"),
		Chapters: db.Collection("chapters"),
	}, nil
}

func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}

func (db *Database) SaveComic(comic Comic) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"endpoint": comic.Endpoint}
	update := bson.M{"$set": comic}
	opts := options.Update().SetUpsert(true)

	_, err := db.Comics.UpdateOne(ctx, filter, update, opts)
	return err
}

func (db *Database) GetComicByEndpoint(endpoint string) (*Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comic Comic
	err := db.Comics.FindOne(ctx, bson.M{"endpoint": endpoint}).Decode(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func (db *Database) GetAllComics() ([]Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := db.Comics.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comics []Comic
	if err = cursor.All(ctx, &comics); err != nil {
		return nil, err
	}
	return comics, nil
}

func (db *Database) SaveChapterDetail(endpoint string, detail ChapterDetail) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"endpoint": endpoint}
	update := bson.M{"$set": detail}
	opts := options.Update().SetUpsert(true)

	_, err := db.Chapters.UpdateOne(ctx, filter, update, opts)
	return err
}

func (db *Database) GetComicsPaginated(offset, limit int64) ([]Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.Find().SetSkip(offset).SetLimit(limit)
	cursor, err := db.Comics.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comics []Comic
	if err = cursor.All(ctx, &comics); err != nil {
		return nil, err
	}
	return comics, nil
}

func (db *Database) GetTotalComics() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return db.Comics.CountDocuments(ctx, bson.M{})
}

func (db *Database) SearchComics(query string) ([]Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": query, "$options": "i"}},
			{"desc": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := db.Comics.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comics []Comic
	if err = cursor.All(ctx, &comics); err != nil {
		return nil, err
	}
	return comics, nil
}

func (db *Database) GetComicsByType(comicType string, offset, limit int64) ([]Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"type": comicType}
	opts := options.Find().SetSkip(offset).SetLimit(limit)

	cursor, err := db.Comics.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comics []Comic
	if err = cursor.All(ctx, &comics); err != nil {
		return nil, err
	}
	return comics, nil
}

func (db *Database) GetComicsByTypeStats() (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$group": bson.M{"_id": "$type", "count": bson.M{"$sum": 1}}},
	}

	cursor, err := db.Comics.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []struct {
		ID    string `bson:"_id"`
		Count int64  `bson:"count"`
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	stats := make(map[string]int64)
	for _, result := range results {
		stats[result.ID] = result.Count
	}

	return stats, nil
}
