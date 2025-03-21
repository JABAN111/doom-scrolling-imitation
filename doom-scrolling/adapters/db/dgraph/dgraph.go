package dgraph

import (
	"context"
	"log/slog"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/db/graph"

	"github.com/dgraph-io/dgo/v240/protos/api"
)

type DgraphDb struct {
	log          *slog.Logger
	dgraphClient *graph.DgraphClient
}

func New(log *slog.Logger, cfg config.Config) (*DgraphDb, error) {
	dgraphClient, err := graph.InitClient(cfg, log)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}

	op := &api.Operation{
		Schema: `
			id: string @index(exact) .
			username: string @index(exact) @upsert .
			follows: uid @reverse .
			likes: uid .
		`,
	}
	err = dgraphClient.Dg.Alter(ctx, op)
	if err != nil {
		log.Error("Fail to set schema", "error", err)
		return nil, err
	}

	return &DgraphDb{log: log, dgraphClient: dgraphClient}, nil
}

func (db *DgraphDb) FollowUser(ctx context.Context, followerID, followingID string) error {
	query := `
	{
		follower as var(func: eq(id, "` + followerID + `"))
		following as var(func: eq(id, "` + followingID + `"))
	}`

	mu := &api.Mutation{
		SetNquads: []byte(`
			uid(follower) <follows> uid(following) .
		`),
		CommitNow: true,
	}

	req := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	_, err := db.dgraphClient.Dg.NewTxn().Do(ctx, req)
	if err != nil {
		db.log.Error("Error of following", "error", err)
		return err
	}

	db.log.Info("User followed", "follower", followerID, "following", followingID)
	return nil
}

func (db *DgraphDb) LikePost(ctx context.Context, userID, postID string) error {
	query := `
	{
		user as var(func: eq(id, "` + userID + `"))
		post as var(func: eq(id, "` + postID + `"))
	}`

	mu := &api.Mutation{
		SetNquads: []byte(`
			uid(user) <likes> uid(post) .
		`),
		CommitNow: true,
	}

	req := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	_, err := db.dgraphClient.Dg.NewTxn().Do(ctx, req)
	if err != nil {
		db.log.Error("Error while liking", "error", err)
		return err
	}

	db.log.Debug("User liked", "user", userID, "post", postID)
	return nil
}

func (db *DgraphDb) GetFeed(ctx context.Context, userID string) (string, error) {
	query := `
	{
		var(func: eq(id, "` + userID + `")) {
			followed as follows
		}

		feed(func: uid(followed)) {
			posts {
				id
				caption
				image_url
				likes
				created_at
			}
		}
	}`

	resp, err := db.dgraphClient.Dg.NewTxn().Query(ctx, query)
	if err != nil {
		db.log.Error("Error request to feed", "error", err)
		return "", err
	}

	db.log.Debug("User feed", "userID", userID, "data", resp.String())
	return resp.String(), nil
}

func (db *DgraphDb) CreateUser(ctx context.Context, username string) error {
	mutation := &api.Mutation{
		SetJson: []byte(`{
			"username": "` + username + `"
		}`),
		CommitNow: true,
	}

	req := &api.Request{
		Mutations: []*api.Mutation{mutation},
		CommitNow: true,
	}

	_, err := db.dgraphClient.Dg.NewTxn().Do(ctx, req)
	if err != nil {
		db.log.Error("Error creating user in Dgraph", "error", err)
		return err
	}

	db.log.Debug("User created in Dgraph", "username", username)
	return nil
}

func (db *DgraphDb) CreatePost(ctx context.Context, postId string) error {
	mutation := &api.Mutation{
		SetJson: []byte(`{
			"id": "` + postId + `"
		}`),
		CommitNow: true,
	}
	req := &api.Request{
		Mutations: []*api.Mutation{mutation},
		CommitNow: true,
	}

	_, err := db.dgraphClient.Dg.NewTxn().Do(ctx, req)
	if err != nil {
		db.log.Error("Error creating post in Dgraph", "error", err)
		return err
	}

	db.log.Debug("Post created in Dgraph", "postID", postId)
	return nil
}
