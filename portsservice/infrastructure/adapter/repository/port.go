package repository

import (
	"context"
	"time"

	"github.com/prigotti/cargo/portsservice/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type portRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoDBPortRepository is a factory for the MongoDB Port Repository adapter
func NewMongoDBPortRepository(
	ctx context.Context,
	database *mongo.Database,
	portCollection string,
) (domain.PortRepository, error) {
	collection := database.Collection(portCollection)

	// This is idempotent, so should be fine
	ctxTO, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	im := mongo.IndexModel{Keys: bson.M{"id": 1}, Options: nil}
	_, err := collection.Indexes().CreateOne(ctxTO, im)
	if err != nil {
		return nil, err
	}

	return &portRepository{
		database:   database,
		collection: collection,
	}, nil
}

// Save is an upsert operation
func (r *portRepository) Save(ctx context.Context, p *domain.Port) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"id": p.ID},
		bson.M{"$set": p},
		options.Update().SetUpsert(true),
	)

	return err
}

// GetWithID uses the custom index we created on this adapter's factory
func (r *portRepository) GetWithID(ctx context.Context, id string) (*domain.Port, error) {
	res := r.collection.FindOne(ctx, bson.M{"id": id})
	if res.Err() == mongo.ErrNoDocuments {
		return nil, domain.ErrAggregateNotFound
	} else if res.Err() != nil {
		return nil, res.Err()
	}

	p := &domain.Port{}

	if err := res.Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}
