package database

import (
	"context"

	"github.com/Edilberto-Vazquez/website-services/constants"
	"github.com/Edilberto-Vazquez/website-services/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	coll *mongo.Collection
}

type ProjectsList struct {
	Projects []models.Project `json:"projects" bson:"projects"`
}

func NewMongoDBRepository(url string) (*MongoDBRepository, error) {
	opts := options.Client()
	opts.ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	db := client.Database(constants.DB_NAME)
	coll := db.Collection(constants.DB_COLLECTION)
	return &MongoDBRepository{coll: coll}, nil
}

func (mongo *MongoDBRepository) GetProfile(ctx context.Context, lang string) (profile models.Profile, err error) {
	filter := bson.M{"lang": lang}
	err = mongo.coll.FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return
}

func (mongo *MongoDBRepository) GetProjects(ctx context.Context, lang string) (projects models.Projects, err error) {
	filter := bson.M{"lang": lang}
	err = mongo.coll.FindOne(ctx, filter).Decode(&projects)
	if err != nil {
		return projects, err
	}
	return
}

func (mongo *MongoDBRepository) GetJobs(ctx context.Context, lang string) (jobs models.Jobs, err error) {
	filter := bson.M{"lang": lang}
	err = mongo.coll.FindOne(ctx, filter).Decode(&jobs)
	if err != nil {
		return jobs, err
	}
	return
}
