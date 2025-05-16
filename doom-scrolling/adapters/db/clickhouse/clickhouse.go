package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"rshd/lab1/v2/internal/util"

	"rshd/lab1/v2/core"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type clickhouseDb struct {
	log  *slog.Logger
	conn clickhouse.Conn
}

func New(log *slog.Logger) *clickhouseDb {
	conn, err := connect(log)
	if err != nil {
		log.Error("ClickHouse connection error", "err", err)
		return nil
	}

	_ = createTables(conn)
	_ = conn.Ping(context.Background())
	return &clickhouseDb{
		log:  log,
		conn: conn,
	}
}

func connect(log *slog.Logger) (clickhouse.Conn, error) {
	ctx := context.Background()
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9020"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "an-example-go-client", Version: "0.1"},
			},
		},
		Debugf: func(format string, v ...any) {
			fmt.Printf(format, v...)
		},
	})
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	log.Info("ClickHouse connection established")
	return conn, nil
}

func (db *clickhouseDb) InsertAnalyticsEvent(ctx context.Context, event core.AnalyticsEvent) error {
	query := `INSERT INTO analytics_events (type, user_id, post_id, target_id, timestamp, additional) VALUES (?, ?, ?, ?, ?, ?)`

	return db.conn.Exec(ctx, query,
		event.Type,
		event.UserID,
		event.PostID,
		event.TargetID,
		event.Timestamp,
		fmt.Sprintf("%v", event.Additional),
	)
}

func (db *clickhouseDb) GetTopActions(ctx context.Context, days, limit int) ([]core.TagStat, error) {
	query := fmt.Sprintf(`
		SELECT tag, count(*) AS cnt
		FROM analytics_events
		WHERE timestamp >= now() - INTERVAL %d DAY
		GROUP BY tag
		ORDER BY cnt DESC
		LIMIT %d
	`, days, limit)

	rows, err := db.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer util.SafeClose(rows)

	var stats []core.TagStat
	for rows.Next() {
		var tag string
		var count uint64
		if err := rows.Scan(&tag, &count); err != nil {
			return nil, err
		}
		stats = append(stats, core.TagStat{
			Tag:   tag,
			Count: count,
		})
	}
	return stats, nil
}

func (db *clickhouseDb) GetUserActivityStats(ctx context.Context, userID string) (core.UserActivity, error) {
	query := `
		SELECT sum(total_posts) as total_posts, active_days
		FROM user_activity
		WHERE user_id = ? 
		GROUP BY active_days
	`
	var totalPosts uint64
	var activeDays uint64
	if err := db.conn.QueryRow(ctx, query, userID).Scan(&totalPosts, &activeDays); err != nil {
		return core.UserActivity{}, err
	}
	activity := core.UserActivity{
		UserID:     userID,
		TotalPosts: int(totalPosts),
		ActiveDays: int(activeDays),
	}
	return activity, nil
}

func (db *clickhouseDb) ExecuteAnalyticsQuery(ctx context.Context, query string, args ...any) ([]any, error) {
	rows, err := db.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer util.SafeClose(rows)

	columns := rows.Columns()
	var results []any

	for rows.Next() {

		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			rowMap[colName] = values[i]
		}
		results = append(results, rowMap)
	}
	return results, nil
}

// вместо нормальной миграции. Потому что не постгря и долбиться мне искренно лень
func createTables(conn clickhouse.Conn) error {
	ctx := context.Background()

	err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS analytics_events (
			type String,
			user_id String,
			post_id String,
			target_id String,
			timestamp DateTime64(3),
			additional String,
			tag String MATERIALIZED JSONExtractString(additional, 'tag')
		) ENGINE = MergeTree()
		ORDER BY (type, timestamp)
	`)
	if err != nil {
		return fmt.Errorf("failed to create analytics_events: %w", err)
	}

	err = conn.Exec(ctx, `
        CREATE MATERIALIZED VIEW IF NOT EXISTS tags_mv
        ENGINE = SummingMergeTree()
        ORDER BY (tag)
        POPULATE AS
        SELECT 
            tag,
            count() as count
        FROM analytics_events
        GROUP BY tag
    `)
	if err != nil {
		return fmt.Errorf("failed to create tags view: %w", err)
	}

	err = conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS content_stats (
            content_id String,
            views UInt64,
            likes UInt64,
            shares UInt64,
            content_type String,
            timestamp DateTime64(3)
        ) ENGINE = MergeTree()
        ORDER BY (content_type, timestamp)
    `)
	if err != nil {
		return fmt.Errorf("failed to create tags view: %w", err)
	}

	err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS user_activity (
			user_id String,
			period Int32,
			total_posts UInt64,
			active_days UInt64
		) ENGINE = MergeTree()
		ORDER BY (user_id, period)
	`)
	if err != nil {
		return fmt.Errorf("failed to create user_activity: %w", err)
	}

	err = conn.Exec(ctx, `
		CREATE MATERIALIZED VIEW IF NOT EXISTS user_activity_mv
		TO user_activity
		AS
		SELECT 
			user_id,
			30 AS period,
			countIf(type = 'create_post') AS total_posts,
			countDistinct(toDate(timestamp)) AS active_days
		FROM analytics_events
		GROUP BY user_id
	`)
	if err != nil {
		return fmt.Errorf("failed to create user_activity materialized view: %w", err)
	}
	return err
}
