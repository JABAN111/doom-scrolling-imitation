package logger

import (
	"fmt"
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
	f, err := os.OpenFile("/app/data.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
			err := c.logFile.Sync()
			if err != nil {
				fmt.Println("не смогли синкнуть")
			}
			err = c.SSS.UploadLogs("/app/data.json")
			if err != nil {
				fmt.Println("произошла критическая да")
				fmt.Println(err)
			}
			c.mu.Unlock()
		}
	}
}
