package database

import (
	"context"

	"github.com/Edilberto-Vazquez/website-services/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	coll *mongo.Collection
}

func NewMongoDBRepository(dbName, dbCollection, uri string) (*MongoDBRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	coll := db.Collection(dbCollection)
	return &MongoDBRepository{coll: coll}, nil
}

func (db *MongoDBRepository) GetProfile(ctx context.Context, lang string) (profile models.Profile, err error) {
	filter := bson.M{"lang": lang}
	err = db.coll.FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return
}

func (db *MongoDBRepository) GetProjects(ctx context.Context, lang string) (projects []models.Project, err error) {
	filter := bson.M{"lang": lang}
	err = db.coll.FindOne(ctx, filter).Decode(&projects)
	if err != nil {
		return projects, err
	}
	return
}
func (db *MongoDBRepository) GetResume(ctx context.Context, lang string) (resume models.Resume, err error) {
	filter := bson.M{"lang": lang}
	err = db.coll.FindOne(ctx, filter).Decode(&resume)
	if err != nil {
		return resume, err
	}
	return
}
func (db *MongoDBRepository) GetJobs(ctx context.Context, lang string) (jobs []models.Job, err error) {
	filter := bson.M{"lang": lang}
	err = db.coll.FindOne(ctx, filter).Decode(&jobs)
	if err != nil {
		return jobs, err
	}
	return
}
