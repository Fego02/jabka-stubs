package handlers

import "log/slog"

var (
	LevelNonMatchedRequests = slog.Level(-2)
	LevelMatchedRequests    = slog.Level(-1)
)
