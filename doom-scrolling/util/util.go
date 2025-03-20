package util

import (
	"io"
	"log/slog"
)

var log = slog.Default()

func SafeClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Error("Failed to close file", "error", err)
	}
}
