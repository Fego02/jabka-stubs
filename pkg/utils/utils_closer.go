package utils

import (
	"io"
	"log/slog"
)

func HandleClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		slog.Log(nil, LevelCloseError, "error while closing", "context", err.Error())
	}
}
