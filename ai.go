package ai

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	idFieldName  = "id"
	seqFieldName = "seq"
)

type (
	AI struct {
		client       *mongo.Client
		collection   *mongo.Collection
		ctx          context.Context
		idFieldName  string
		seqFieldName string
	}
)

// Create new instance of AI
func Create(c *mongo.Collection, fieldNames ...string) *AI {
	ai := &AI{
		collection:   c,
		client:       c.Database().Client(),
		idFieldName:  "id",
		seqFieldName: "seq",
	}

	if len(fieldNames) > 0 {
		ai.idFieldName = fieldNames[0]
	}
	if len(fieldNames) > 1 {
		ai.seqFieldName = fieldNames[1]
	}
	return ai
}

func (ai *AI) Next(name string) (seq int32) {

	ai.connectionCheck()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetBypassDocumentValidation(true)

	filter := bson.M{ai.idFieldName: name}
	update := bson.M{"$set": bson.M{ai.idFieldName: name}, "$inc": bson.M{ai.seqFieldName: 1}}

	result := ai.collection.FindOneAndUpdate(context.TODO(), filter, update, opts)

	reg, err := result.DecodeBytes()
	if err == nil {
		seq = reg.Index(2).Value().Int32()
	} else {
		reg, err = ai.collection.FindOneAndUpdate(context.TODO(), filter, update, opts).DecodeBytes()
		if err == nil {
			seq = reg.Index(2).Value().Int32()
		}
	}

	return seq

}

func (ai *AI) Cancel(name string) {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetBypassDocumentValidation(true)

	filter := bson.M{ai.idFieldName: name}
	update := bson.M{"$set": bson.M{ai.idFieldName: name}, "$inc": bson.M{ai.seqFieldName: -1}}

	ai.collection.FindOneAndUpdate(context.TODO(), filter, update, opts)

}

func (ai *AI) connectionCheck() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := ai.client.Ping(ctx, readpref.Primary()); err != nil {
		ai.client.Connect(ctx)
	}
}
