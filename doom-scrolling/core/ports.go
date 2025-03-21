package core

import "context"

type DocumentDB interface {
	CreateUser(ctx context.Context, user User) (User, error)
	CreatePost(ctx context.Context, post Post) (Post, error)
}

type GraphDB interface {
	FollowUser(ctx context.Context, followerID, followingID string) error
	LikePost(ctx context.Context, userID, postID string) error
	GetFeed(ctx context.Context, userID string) (string, error)

	CreateUser(ctx context.Context, username string) error
	CreatePost(ctx context.Context, postId string) error
}
