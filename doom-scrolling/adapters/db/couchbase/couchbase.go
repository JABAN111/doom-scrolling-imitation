package couchbase

import (
	"context"
	"github.com/couchbase/gocb/v2"
	"log/slog"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/internal/db/document"
	"time"
)

type Collections struct {
	userCollection *gocb.Collection
	postCollection *gocb.Collection
}

type CouchDB struct {
	log         *slog.Logger
	conn        *gocb.Cluster
	bucket      *gocb.Bucket
	collections Collections
}

// collections
const (
	userCollection = "users"
	postCollection = "posts"
)

// scopes
const scope = "doom-data"

func New(log *slog.Logger, cfg config.Config) (*CouchDB, error) {
	cluster, err := document.InitializeCluster(cfg, log)
	if err != nil {
		return nil, err
	}

	res := &CouchDB{log: log, conn: cluster}

	res.bucket = cluster.Bucket("doom-scrolling")
	err = res.bucket.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		log.Error("Failed to wait Couchbase bucket ready", "error", err)
		return nil, err
	}

	mgr := res.bucket.CollectionsV2()
	if err = mgr.CreateScope(scope, nil); err != nil {
		log.Warn("Failed to create scope", "scope", scope, "error", err)
	}

	res.collections.userCollection = res.bucket.Scope(scope).Collection(userCollection)
	res.collections.postCollection = res.bucket.Scope(scope).Collection(postCollection)

	spec := gocb.CollectionSpec{
		ScopeName: scope,
		Name:      userCollection,
	}
	if err = mgr.CreateCollection(spec.ScopeName, spec.Name, nil, nil); err != nil {
		log.Warn("Failed to create collection of users", "error", err)
	}
	spec.Name = postCollection
	if err = mgr.CreateCollection(spec.ScopeName, spec.Name, nil, nil); err != nil {
		log.Warn("Failed to create collection of posts", "error", err)
	}

	return res, nil
}

func (d *CouchDB) CreateUser(_ context.Context, user core.User) (core.User, error) {
	_, err := d.collections.userCollection.Upsert(user.Username, user, nil)
	if err != nil {
		return core.User{}, err
	}
	return user, nil
}

func (d *CouchDB) CreatePost(_ context.Context, post core.Post) (core.Post, error) {
	_, err := d.collections.postCollection.Upsert(post.ID, post, nil)
	if err != nil {
		return core.Post{}, err
	}
	return post, nil
}

func (d *CouchDB) GetPost(ctx context.Context, postId string) (core.Post, error) {
	getResult, err := d.collections.postCollection.Get(postId, nil)
	if err != nil {
		return core.Post{}, err
	}

	var post core.Post
	if err := getResult.Content(&post); err != nil {
		return core.Post{}, err
	}
	return post, nil

}
