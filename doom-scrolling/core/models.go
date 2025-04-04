package core

import "time"

type EventType string

const (
	EventPost       EventType = "post"
	EventLike       EventType = "like"
	EventFollow     EventType = "follow"
	EventCreatePost EventType = "create_post"
	EventCreateUser EventType = "create_user"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

type AnalyticsEvent struct {
	Type       EventType
	UserID     string
	PostID     string
	TargetID   string
	Timestamp  time.Time
	Additional string
}

type TagStat struct {
	Tag   string
	Count uint64
}

type ContentStat struct {
	ContentID string
	Views     int
	Likes     int
	Shares    int
}

type UserActivity struct {
	UserID     string
	TotalPosts int
	ActiveDays int
}

type RevenueStat struct {
	Date         time.Time
	Source       string
	Amount       float64
	Country      string
	CampaignName string
}

type TimeSeriesEvent struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]any
	Timestamp   time.Time
}

type SystemMetric struct {
	Type       string
	Value      float64
	Service    string
	InstanceID string
}

type SystemHealthStats struct {
	CPUUsage     float64
	MemoryUsage  float64
	ActiveUsers  int
	RequestRate  float64
	ErrorRate    float64
	ResponseTime float64
}
