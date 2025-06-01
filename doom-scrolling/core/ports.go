package core

import (
	"context"
	"time"
)

type DocumentDB interface {
	CreateUser(ctx context.Context, user User) (User, error)
	CreatePost(ctx context.Context, post Post) (Post, error)

	GetPost(ctx context.Context, postId string) (Post, error)
}

type GraphDB interface {
	FollowUser(ctx context.Context, username, usernameToFollow string) error
	LikePost(ctx context.Context, userID, postID string) error
	GetFeed(ctx context.Context, userID string) ([]string, error)

	CreateUser(ctx context.Context, username string) error
	CreatePost(ctx context.Context, postId string, userId string) error
}

type ColumnarDB interface {
	InsertAnalyticsEvent(ctx context.Context, event AnalyticsEvent) error
	GetTopActions(ctx context.Context, days int, limit int) ([]TagStat, error)

	GetUserActivityStats(ctx context.Context, userID string) (UserActivity, error)
}

type TimeSeriesDB interface {
	WriteEvent(ctx context.Context, event TimeSeriesEvent) error
	GetEvents(ctx context.Context, measurement string, filters map[string]string, from, to time.Time) ([]TimeSeriesEvent, error)

	GetEventCount(ctx context.Context, measurement string, duration time.Duration) (int64, error)
	GetRatePerMinute(ctx context.Context, measurement string) (float64, error)

	WriteSystemMetric(ctx context.Context, metric SystemMetric) error
	GetSystemHealth(ctx context.Context) (SystemHealthStats, error)
}

type S3 interface {
	UploadPostImage(ctx context.Context, id string, filePath string) error
	// DownloadPostImage contains filepath because some s3(especially minio) are saving file to specified path,
	// without returning os.File or something like this.
	DownloadPostImage(ctx context.Context, id string, filePath string) error
	DeletePostImage(ctx context.Context, id string) error

	UploadLogs(filepath string) error
}
