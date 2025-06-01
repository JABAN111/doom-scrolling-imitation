package logger

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"os"
	"rshd/lab1/v2/core"
	"sync"
	"time"
)

//кастомный логгер, который будет записывать все в файл data.json
// и раз в заданное время загружать в s3

type Custom struct {
	SSS     core.S3
	log     *slog.Logger
	logFile *os.File
	mu      sync.Mutex
}

func GetCustom(sss core.S3, duration time.Duration) *Custom {
	f, err := os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	handler := slog.NewJSONHandler(f, nil)
	logger := slog.New(handler)

	c := &Custom{
		SSS:     sss,
		log:     logger,
		logFile: f,
	}

	go c.startUploader(duration)

	return c
}

func (c *Custom) Info(msg string, args ...any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.log.Info(msg, args...)
}

func (c *Custom) startUploader(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			f, err := os.Open("data.json")
			if err != nil {
				c.mu.Unlock()
				continue
			}
			defer f.Close()

			var entries []json.RawMessage
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				entries = append(entries, json.RawMessage(scanner.Bytes()))
			}

			tmpFile := "data_array.json"
			out, _ := os.Create(tmpFile)
			enc := json.NewEncoder(out)
			enc.SetIndent("", "  ")
			enc.Encode(entries)
			out.Close()

			_ = c.SSS.UploadLogs(tmpFile)
			_ = os.Remove(tmpFile)
			c.mu.Unlock()
		}
	}
}
