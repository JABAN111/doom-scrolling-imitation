package dgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/internal/db/graph"

	"github.com/dgraph-io/dgo/v240"

	"github.com/dgraph-io/dgo/v240/protos/api"
)

type DbDgraph struct {
	log          *slog.Logger
	dgraphClient *graph.DgraphClient
}

type response struct {
	People []struct {
		UID string `json:"uid"`
	} `json:"people"`
}

func New(log *slog.Logger, cfg config.Config) (*DbDgraph, error) {
	dgraphClient, err := graph.InitClient(cfg, log)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}

	op := &api.Operation{
		Schema: `
			name: string @index(exact) @upsert .
			id: string @index(exact) .
			username: string @index(exact) @upsert .
			follows: uid @reverse .
			likes: uid @reverse .
		`,
	}
	err = dgraphClient.Dg.Alter(ctx, op)
	if err != nil {
		log.Error("Fail to set schema", "error", err)
		return nil, err
	}

	return &DbDgraph{log: log, dgraphClient: dgraphClient}, nil
}

func (db *DbDgraph) FollowUser(_ context.Context, followerName, followingName string) error {
	txn := db.dgraphClient.Dg.NewTxn()
	defer func(txn *dgo.Txn, ctx context.Context) {
		_ = txn.Discard(ctx)
	}(txn, context.Background())

	queryFollower := `
	{
		people(func: eq(username, "` + followingName + `")) {
			uid
		}
	}`
	var structFollower response
	var structFollowing response

	queryFollowing := `
	{
		people(func: eq(username, "` + followerName + `")) {
			uid
		}
	}`

	respFollower, err := txn.Query(context.Background(), queryFollower)
	if err != nil {
		db.log.Error("Query failed", "error", err)
		return err
	}

	if err := json.Unmarshal(respFollower.Json, &structFollower); err != nil {
		db.log.Error("Error unmarshaling JSON", "error", err)
		return err
	}

	respFollowing, err := txn.Query(context.Background(), queryFollowing)
	if err != nil {
		db.log.Error("Query failed", "error", err)
		return err
	}
	if err := json.Unmarshal(respFollowing.Json, &structFollowing); err != nil {
		db.log.Error("Error unmarshaling JSON", "error", err)
		return err
	}

	if structFollower.People == nil || structFollowing.People == nil {
		db.log.Error("User not found", "follower", followerName, "following", followingName)
		return fmt.Errorf("user not found")
	}

	mutation := &api.Mutation{
		SetJson: []byte(`
			{
				"uid": "` + structFollowing.People[0].UID + `",
				"follows": [{
					"uid": "` + structFollower.People[0].UID + `"
				}]
			}
		`),
		CommitNow: true,
	}
	_, err = txn.Mutate(context.Background(), mutation)
	if err != nil {
		db.log.Error("Mutation failed", "error", err)
		return err
	}
	return nil
}

func (db *DbDgraph) LikePost(ctx context.Context, userID, postID string) error {
	txn := db.dgraphClient.Dg.NewTxn()
	defer func(txn *dgo.Txn, ctx context.Context) {
		_ = txn.Discard(ctx)
	}(txn, context.Background())

	userUid := `
	{
		people(func: eq(username, "` + userID + `")) {
			uid
		}
	}`

	postUid := `
	{
		people(func: eq(id, "` + postID + `")) {
			uid
		}
	}`

	var user response
	var post response

	respUser, err := txn.Query(context.Background(), userUid)
	if err != nil {
		db.log.Error("Query failed", "error", err)
		return err
	}
	respPost, err := txn.Query(context.Background(), postUid)
	if err != nil {
		db.log.Error("Query failed", "error", err)
		return err
	}

	if err := json.Unmarshal(respUser.Json, &user); err != nil {
		db.log.Error("Error unmarshaling JSON", "error", err)
		return err
	}
	if err := json.Unmarshal(respPost.Json, &post); err != nil {
		db.log.Error("Error unmarshaling JSON", "error", err)
		return err
	}

	mutation := &api.Mutation{
		SetJson: []byte(`
			{
				"uid": "` + post.People[0].UID + `",
				"likes": {
					"uid": "` + user.People[0].UID + `"
				}
			}
		`),
		CommitNow: true,
	}
	_, err = txn.Mutate(ctx, mutation)
	if err != nil {
		db.log.Error("Mutation failed", "error", err)
		return err
	}
	return nil
}

func (db *DbDgraph) GetFeed(ctx context.Context, userID string) ([]string, error) {
	txn := db.dgraphClient.Dg.NewTxn()
	defer func(txn *dgo.Txn, ctx context.Context) {
		_ = txn.Discard(ctx)
	}(txn, context.Background())

	query := fmt.Sprintf(`
	{
		people(func: eq(username, "%s")) {
			uid
			follows {
				uid
			}
		}
	}`, userID)

	queryResponse, err := txn.Query(ctx, query)
	if err != nil {
		db.log.Error("Query failed", "error", err)
		return nil, err
	}

	var userResponse struct {
		People []struct {
			UID     string `json:"uid"`
			Follows []struct {
				UID string `json:"uid"`
			} `json:"follows"`
		} `json:"people"`
	}
	if err := json.Unmarshal(queryResponse.Json, &userResponse); err != nil {
		db.log.Error("Error unmarshalling JSON (GetFeed)", "error", err)
		return nil, err
	}
	if len(userResponse.People) == 0 {
		db.log.Error("User not found", "username", userID)
		return nil, fmt.Errorf("user not found")
	}
	return nil, nil
}

func (db *DbDgraph) CreateUser(ctx context.Context, username string) error {
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

func (db *DbDgraph) CreatePost(ctx context.Context, postId, userId string) error {
	mutation := &api.Mutation{
		SetJson: []byte(`{
			"id": "` + postId + `",
			"user_id": "` + userId + `"
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

func (db *DbDgraph) READ() {
	txn := db.dgraphClient.Dg.NewTxn()
	defer func(txn *dgo.Txn, ctx context.Context) {
		_ = txn.Discard(ctx)
	}(txn, context.Background())

	query := `
	{
		people(func: has(name)) {
			name
			age
    	follows{
				name
        age
      }
  }
}`

	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		panic(err)

	}

	db.log.Info("Query result:", "response", resp.Json)

}

func (db *DbDgraph) RANDOMBULLSHITGO() {
	txn := db.dgraphClient.Dg.NewTxn()
	defer func(txn *dgo.Txn, ctx context.Context) {
		_ = txn.Discard(ctx)
	}(txn, context.Background())

	mutation := &api.Mutation{
		SetJson: []byte(`
			{
				"name": "Ann",
				"age": 28,
				"follows": [
					{
						"name": "Ben",
						"age": 31
					}
				]
			}
		`),
		CommitNow: true,
	}

	_, err := txn.Mutate(context.Background(), mutation)
	if err != nil {
		db.log.Error("Mutation failed", "error", err)
		panic(err)
	}

}
