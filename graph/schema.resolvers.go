package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/karansinghgit/gqlgen-demo/graph/generated"
	"github.com/karansinghgit/gqlgen-demo/graph/model"
)

var es *elasticsearch.Client

func init() {
	var err error
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	idx := "feelr"
	ctx := context.Background()

	_, err = esapi.IndicesDeleteRequest{Index: []string{idx}}.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error on deleting Index: %s", err)
	}
	_, err2 := esapi.IndicesCreateRequest{Index: idx}.Do(ctx, es)
	if err2 != nil {
		log.Fatalf("Error on creating Index: %s", err)
	}
}

func (r *mutationResolver) SendTextMessage(ctx context.Context, chatID string, text string) (model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFeelr(ctx context.Context, chatID string, feelrID string, answer string) (*model.Feelr, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendFeelrResponse(ctx context.Context, chatID string, feelrID string, answer string) (*model.Feelr, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTopFeelrs(ctx context.Context, top *int) ([]*model.Feelr, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMessages(ctx context.Context, chatID string, last *int) ([]model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, chatID string) (<-chan model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
