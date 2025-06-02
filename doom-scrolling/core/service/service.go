package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/internal/logger"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type Service struct {
	DocumentDB   core.DocumentDB
	GraphDB      core.GraphDB
	TimeSeriesDB core.TimeSeriesDB
	ColumnarDB   core.ColumnarDB
	S3           core.S3
	log          *slog.Logger
	numWorkers   int
	customLogger *logger.Custom
}

func NewService(
	log *slog.Logger,
	db core.DocumentDB,
	graph core.GraphDB,
	tsDB core.TimeSeriesDB,
	columnarDB core.ColumnarDB,
	s3 core.S3,
	numWorkers int,
) *Service {
	return &Service{
		DocumentDB:   db,
		GraphDB:      graph,
		TimeSeriesDB: tsDB,
		ColumnarDB:   columnarDB,
		log:          log,
		S3:           s3,
		numWorkers:   numWorkers,
		customLogger: logger.GetCustom(s3, 5*time.Second),
	}
}

func (s *Service) CreateUser(ctx context.Context, user core.User) (core.User, error) {
	user, err := s.DocumentDB.CreateUser(ctx, user)
	if err != nil {
		s.log.Info("Failed to create user", "error", err)
		return core.User{}, err
	}

	go s.logUserActivity(ctx, user.Username, "user_created")

	err = s.GraphDB.CreateUser(ctx, user.Username)
	if err != nil {
		s.log.Info("Failed to create user in graph", "error", err)
		return core.User{}, err
	}

	go s.customLogger.Info("create user", "user", user)

	return user, nil
}

func (s *Service) CreatePost(ctx context.Context, post core.Post, filePath string) (core.Post, error) {
	post.CreatedAt = time.Now()
	post, err := s.DocumentDB.CreatePost(ctx, post)
	if err != nil {
		s.log.Error("Failed to create post", "error", err)
		return core.Post{}, err
	}

	err = s.S3.UploadPostImage(ctx, post.ID, filePath)
	if err != nil {
		s.log.Error("failed to save an image for the post", "error", err)
		return core.Post{}, err
	}

	go func() {
		serEvnt := core.TimeSeriesEvent{
			Measurement: "post_created",
			Tags: map[string]string{
				"user_id": post.UserID,
				"post_id": post.ID},
			Fields:    map[string]any{"post": 1},
			Timestamp: time.Now(),
		}
		err = s.TimeSeriesDB.WriteEvent(ctx, serEvnt)
		if err != nil {
			s.log.Error("Failed to write post event", "error", err, "event", serEvnt)
		}
	}()

	go func() {
		_ = s.ColumnarDB.InsertAnalyticsEvent(context.Background(), core.AnalyticsEvent{
			Type:       core.EventCreatePost,
			UserID:     post.UserID,
			PostID:     post.ID,
			Timestamp:  time.Now(),
			Additional: fmt.Sprintf(`{"tag": "%s"}`, core.EventCreatePost),
		})
	}()

	err = s.GraphDB.CreatePost(ctx, post.ID, post.UserID)
	if err != nil {
		s.log.Info("Failed to create post in graph", "error", err)
		return core.Post{}, err
	}

	go s.customLogger.Info("create post", "post", post)

	return post, nil
}

func (s *Service) LikePost(ctx context.Context, userID, postID string) error {
	if err := s.GraphDB.LikePost(ctx, userID, postID); err != nil {
		s.log.Info("Failed to like post", "error", err)
		return err
	}

	go s.logUserActivity(ctx, userID, core.EventLike)

	err := s.TimeSeriesDB.WriteEvent(ctx, core.TimeSeriesEvent{
		Measurement: "post_likes",
		Tags: map[string]string{
			"user_id": userID,
			"post_id": postID,
		},
		Fields:    map[string]any{string(core.EventLike): 1},
		Timestamp: time.Now(),
	})
	if err != nil {
		s.log.Error("Failed to write like event", "error", err)
	}

	go s.customLogger.Info("like post", "userID", userID, "postID", postID)

	return nil
}

func (s *Service) GetPopularTags(ctx context.Context, days, limit int) ([]core.TagStat, error) {
	return s.ColumnarDB.GetTopActions(ctx, days, limit)
}

func (s *Service) GetSystemHealth(ctx context.Context) (core.SystemHealthStats, error) {
	return s.TimeSeriesDB.GetSystemHealth(ctx)
}

func (s *Service) logUserActivity(ctx context.Context, userID string, eventType core.EventType) {
	go func() {
		err := s.TimeSeriesDB.WriteEvent(ctx, core.TimeSeriesEvent{
			Measurement: "user_activity",
			Tags:        map[string]string{"user_id": userID, "type": string(eventType)},
			Timestamp:   time.Now(),
			Fields:      map[string]any{"count": 0},
		})
		if err != nil {
			s.log.Error("Failed to log user activity", "error", err)
		}
	}()

	go func() {
		_ = s.ColumnarDB.InsertAnalyticsEvent(context.Background(), core.AnalyticsEvent{
			Type:       eventType,
			UserID:     userID,
			Timestamp:  time.Now(),
			Additional: fmt.Sprintf(`{"tag": "%s"}`, string(eventType)),
		})
	}()
}
func (s *Service) FollowUser(ctx context.Context, username, usernameToFollow string) error {
	if err := s.GraphDB.FollowUser(ctx, username, usernameToFollow); err != nil {
		s.log.Info("Failed to follow user", "error", err)
		return err
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_ = s.TimeSeriesDB.WriteEvent(context.Background(), core.TimeSeriesEvent{
			Measurement: "user_follow",
			Tags: map[string]string{
				"follower":  username,
				"following": usernameToFollow,
			},
			Fields:    map[string]any{"count": 0},
			Timestamp: time.Now(),
		})

	}()
	go func() {
		defer wg.Done()

		_ = s.ColumnarDB.InsertAnalyticsEvent(context.Background(), core.AnalyticsEvent{
			Type:       "follow",
			UserID:     username,
			TargetID:   usernameToFollow,
			Timestamp:  time.Now(),
			Additional: `{"tag": "follow"}`,
		})
	}()
	wg.Wait()
	go s.customLogger.Info("follow user", "username", username, "usernameToFollow", usernameToFollow)
	return nil
}

func (s *Service) GetFeed(ctx context.Context, userID string) ([]core.Post, error) {
	feeds, err := s.GraphDB.GetFeed(ctx, userID)
	if err != nil {
		s.log.Info("Failed to get feed", "error", err)
		return nil, err
	}

	sema := make(chan struct{}, s.numWorkers)
	postChan := make(chan core.Post)
	wg := sync.WaitGroup{}
	wg.Add(len(feeds))

	for _, postId := range feeds {
		go func() {
			defer wg.Done()
			defer func() {
				<-sema
			}()

			sema <- struct{}{}

			tmp, err := os.CreateTemp("", "img-*.jpg")
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					return
				}
			}(tmp.Name())
			if err != nil {
				return
			}
			post, err := s.DocumentDB.GetPost(ctx, postId)

			if err := s.S3.DownloadPostImage(ctx, post.ID, tmp.Name()); err != nil {
				s.log.Error("fail to download image for the post", "err", err)
				return
			} else {
				data, err := os.ReadFile(tmp.Name())
				if err == nil {
					post.ImageBase64 = base64.StdEncoding.EncodeToString(data)
				}
			}
			if err != nil {
				s.log.Error("Failed to parse post", "postid", postId)
			}
			postChan <- post
		}()
	}

	go func() {
		wg.Wait()
		close(postChan)
	}()

	var posts []core.Post

	for val := range postChan {
		posts = append(posts, val)
	}

	go func() {
		_ = s.TimeSeriesDB.WriteEvent(context.Background(), core.TimeSeriesEvent{
			Measurement: "feed_viewed",
			Tags:        map[string]string{"user_id": userID},
			Fields:      map[string]any{"count": 1},
			Timestamp:   time.Now(),
		})
	}()

	return posts, nil
}

func (s *Service) GetUserStats(ctx context.Context, userID string) (core.UserActivity, error) {
	return s.ColumnarDB.GetUserActivityStats(ctx, userID)
}

func (s *Service) CollectAndStoreSystemMetrics(ctx context.Context) {
	defer s.log.Info("Update of system metrics finished")

	cpuUsage, err := getCPUUsage()
	if err != nil {
		s.log.Error("Failed to get CPU usage", "error", err)
	}
	go func() {
		cpuMetric := core.SystemMetric{
			Service:    "my_service",
			InstanceID: "instance_1",
			Type:       "cpu_usage",
			Value:      cpuUsage,
		}
		if err := s.TimeSeriesDB.WriteSystemMetric(ctx, cpuMetric); err != nil {
			s.log.Error("Failed to write CPU system metric", "error", err)
		}
	}()

	memUsage, err := getMemoryUsage()
	s.log.Info("Updating memory", "mem", memUsage, "cpu", cpuUsage)
	if err != nil {
		s.log.Error("Failed to get memory usage", "error", err)
		return
	}
	go func() {
		memMetric := core.SystemMetric{
			Service:    "my_service",
			InstanceID: "instance_1",
			Type:       "memory_usage",
			Value:      memUsage,
		}
		if err := s.TimeSeriesDB.WriteSystemMetric(ctx, memMetric); err != nil {
			s.log.Error("Failed to write memory system metric", "error", err)
		}
	}()

}

func getCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) == 0 {
		return 0, fmt.Errorf("no CPU data available")
	}
	return percentages[0], nil
}

func getMemoryUsage() (float64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.UsedPercent, nil
}

func (s *Service) StartSystemMetricsCollection(ctx context.Context, tick time.Duration) {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.CollectAndStoreSystemMetrics(ctx)
		}
	}
}
