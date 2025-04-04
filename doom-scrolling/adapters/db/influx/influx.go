package influx

import (
	"context"
	"fmt"
	"log/slog"
	"rshd/lab1/v2/core"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxDb struct {
	log      *slog.Logger
	client   influxdb2.Client
	org      string
	bucket   string
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

func New(log *slog.Logger, url, token, org, bucket string) (*InfluxDb, error) {
	client := influxdb2.NewClient(url, token)
	return &InfluxDb{
		log:      log,
		client:   client,
		org:      org,
		bucket:   bucket,
		writeAPI: client.WriteAPIBlocking(org, bucket),
		queryAPI: client.QueryAPI(org),
	}, nil
}

func (i *InfluxDb) WriteEvent(ctx context.Context, event core.TimeSeriesEvent) error {
	p := influxdb2.NewPoint(
		event.Measurement,
		event.Tags,
		event.Fields,
		event.Timestamp,
	)

	return i.writeAPI.WritePoint(ctx, p)
}

func (i *InfluxDb) GetEvents(ctx context.Context, measurement string, filters map[string]string, from, to time.Time) ([]core.TimeSeriesEvent, error) {
	query := fmt.Sprintf(`from(bucket: "%s")
        |> range(start: %s, stop: %s)
        |> filter(fn: (r) => r._measurement == "%s")`,
		i.bucket, from.Format(time.RFC3339), to.Format(time.RFC3339), measurement)

	for k, v := range filters {
		query += fmt.Sprintf(` |> filter(fn: (r) => r.%s == "%s")`, k, v)
	}

	result, err := i.queryAPI.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var events []core.TimeSeriesEvent
	for result.Next() {
		record := result.Record()
		tags := make(map[string]string)
		for key, value := range record.Values() {
			tags[key] = fmt.Sprintf("%v", value)
		}

		events = append(events, core.TimeSeriesEvent{
			Measurement: measurement,
			Tags:        tags,
			Fields:      map[string]interface{}{record.Field(): record.Value()},
			Timestamp:   record.Time(),
		})
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return events, nil
}

func (i *InfluxDb) GetEventCount(ctx context.Context, measurement string, duration time.Duration) (int64, error) {
	query := fmt.Sprintf(`from(bucket: "%s")
        |> range(start: -%s)
        |> filter(fn: (r) => r._measurement == "%s")
        |> count()`,
		i.bucket, duration.String(), measurement)

	result, err := i.queryAPI.Query(ctx, query)
	if err != nil {
		return 0, err
	}

	if result.Next() {
		if count, ok := result.Record().Value().(int64); ok {
			return count, nil
		} else if countFloat, ok := result.Record().Value().(float64); ok {
			return int64(countFloat), nil
		}
	}
	if result.Err() != nil {
		return 0, result.Err()
	}
	return 0, nil
}

func (i *InfluxDb) GetRatePerMinute(ctx context.Context, measurement string) (float64, error) {
	query := fmt.Sprintf(`from(bucket: "%s")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "%s")
        |> aggregateWindow(every: 1m, fn: count)
        |> mean()`,
		i.bucket, measurement)

	result, err := i.queryAPI.Query(ctx, query)
	if err != nil {
		return 0, err
	}

	if result.Next() {
		if rate, ok := result.Record().Value().(float64); ok {
			return rate, nil
		}
	}
	if result.Err() != nil {
		return 0, result.Err()
	}
	return 0, nil
}

func (i *InfluxDb) WriteSystemMetric(ctx context.Context, metric core.SystemMetric) error {
	p := write.NewPoint(
		"system_metrics",
		map[string]string{
			"service":     metric.Service,
			"instance":    metric.InstanceID,
			"metric_type": metric.Type,
		},
		map[string]any{"value": metric.Value},
		time.Now(),
	)
	return i.writeAPI.WritePoint(ctx, p)
}

func (i *InfluxDb) GetSystemHealth(ctx context.Context) (core.SystemHealthStats, error) {
	stats := core.SystemHealthStats{}

	query := `from(bucket: "` + i.bucket + `")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "system_metrics")
        |> filter(fn: (r) => r.metric_type == "cpu")
        |> last()`

	if result, err := i.queryAPI.Query(ctx, query); err == nil && result.Next() {
		if cpuUsage, ok := result.Record().Value().(float64); ok {
			stats.CPUUsage = cpuUsage
		}
	}

	return stats, nil
}
