package core

import (
	"context"
	"log/slog"
)

type Service struct {
	DocumentDB DocumentDB
	GraphDB    GraphDB
	log        *slog.Logger
}

func NewService(
	log *slog.Logger, db DocumentDB, graph GraphDB,
) *Service {

	return &Service{
		DocumentDB: db,
		GraphDB:    graph,
		log:        log,
	}
}

func (s *Service) CreateUser(ctx context.Context, user User) (User, error) {
	user, err := s.DocumentDB.CreateUser(ctx, user)
	if err != nil {
		s.log.Info("Failed to create user", "error", err)
		return User{}, err
	}
	s.log.Debug("User created on document side", "user", user)

	err = s.GraphDB.CreateUser(ctx, user.Username)
	if err != nil {
		s.log.Info("Failed to create user", "error", err) // NOTE: Сюда бы двухфазный коммит, но требование было к минимальному api)
		return User{}, err
	}
	return user, nil
}

func (s *Service) CreatePost(ctx context.Context, post Post) (Post, error) {
	post, err := s.DocumentDB.CreatePost(ctx, post)
	if err != nil {
		s.log.Info("Failed to create post", "error", err)
		return Post{}, err
	}
	s.log.Debug("Post created on document side", "post", post)

	err = s.GraphDB.CreatePost(ctx, post.ID, post.UserID)
	if err != nil {
		s.log.Info("Failed to create post", "error", err)
		return Post{}, err
	}
	return post, nil
}

func (s *Service) FollowUser(ctx context.Context, followerName, followingName string) error {
	if err := s.GraphDB.FollowUser(ctx, followerName, followingName); err != nil {
		s.log.Info("Failed to follow user", "error", err)
		return err
	}
	return nil
}

func (s *Service) LikePost(ctx context.Context, userID, postID string) error {
	if err := s.GraphDB.LikePost(ctx, userID, postID); err != nil {
		s.log.Info("Failed to like post", "error", err)
		return err
	}
	return nil
}

func (s *Service) GetFeed(ctx context.Context, userID string) (string, error) {
	res, err := s.GraphDB.GetFeed(ctx, userID)
	if err != nil {
		s.log.Info("Failed to get feed", "error", err)
		return "", err
	}
	return res, nil
}
