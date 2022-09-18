package database

import (
	"context"
	"time"

	"github.com/Edilberto-Vazquez/website-services/src/constants"
	"github.com/Edilberto-Vazquez/website-services/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	coll *mongo.Collection
}

type projectList struct {
	Projects []*models.Project `json:"projects" bson:"projects"`
}

type jobList struct {
	Jobs []*models.Job `json:"jobs" bson:"jobs"`
}

func NewMongoDBRepository(url string) (*MongoDBRepository, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	db := client.Database(constants.DB_NAME)
	coll := db.Collection(constants.DB_COLLECTION)
	return &MongoDBRepository{coll: coll}, nil
}

func (mongo *MongoDBRepository) GetProfile(ctx context.Context, lang string) (profile *models.Profile, err error) {
	filter := bson.M{"lang": lang}
	err = mongo.coll.FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return
}

func (mongo *MongoDBRepository) GetProjects(ctx context.Context, lang string) ([]*models.Project, error) {
	filter := bson.M{"lang": lang}
	var res projectList
	err := mongo.coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res.Projects, nil
}

func (mongo *MongoDBRepository) GetJobs(ctx context.Context, lang string) ([]*models.Job, error) {
	filter := bson.M{"lang": lang}
	var res jobList
	err := mongo.coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res.Jobs, nil
}
