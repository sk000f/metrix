package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Config struct {
// 	Username     string
// 	Password     string
// 	DatabaseName string
// 	URI          string
// }

type MongoDB struct {
	c *mongo.Client
}

func New(conn string) (*MongoDB, error) {

	db := &MongoDB{}

	client, err := db.GetClient(conn)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("mongodb.New")
		return nil, err
	}

	db.c = client

	return db, nil
}

func (m *MongoDB) GetClient(conn string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(conn)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("mongodb.GetClient")
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("mongodb.GetClient")
		return nil, err
	}

	return client, nil
}

func (m *MongoDB) GetAll() ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) Update(d []domain.Deployment) error {

	collection := m.c.Database("metrix").Collection("deployments")
	updateOpts := options.Update().SetUpsert(true)

	for _, dep := range d {
		filter := bson.M{"deployment_id": dep.ID}

		update := bson.M{
			"$set": bson.M{
				"deployment_id":     dep.ID,
				"status":            dep.Status,
				"environment_name":  dep.EnvironmentName,
				"project_id":        dep.ProjectID,
				"project_name":      dep.ProjectName,
				"project_path":      dep.ProjectPath,
				"project_namespace": dep.ProjectNamespace,
				"pipeline_id":       dep.PipelineID,
				"finished_at":       dep.FinishedAt,
				"duration":          dep.Duration,
			},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update, updateOpts)
		if err != nil {
			log.Error().Stack().Err(err).
				Msg("mongodb.Update")
			return err
		}
	}

	return nil
}
