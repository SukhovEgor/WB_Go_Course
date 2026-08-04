package grep

import (
	"strings"

	"DistributedGrep/internal/app"
)

func Match(line string, req app.GrepRequest) bool {

	text := line
	pattern := req.Pattern

	if req.IgnoreCase {
		text = strings.ToLower(text)
		pattern = strings.ToLower(pattern)
	}

	found := strings.Contains(text, pattern)

	if req.Invert {
		return !found
	}

	return found
}