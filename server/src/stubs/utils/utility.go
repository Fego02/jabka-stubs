package utils

import (
	"io"
	"log/slog"
	"strings"
)

func HandleClose(stream io.Closer) {
	err := stream.Close()
	if err != nil {
		slog.Info("error while closing", "context", err.Error())
	}
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}
