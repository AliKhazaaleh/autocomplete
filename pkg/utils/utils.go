package utils

import (
	"strconv"
)

func ParseLimit(limitParam string, defaultLimit int, maxLimit int) int {
	limit := defaultLimit
	if limitParam != "" {
		if limitValue, err := strconv.Atoi(limitParam); err == nil && limitValue > 0 && limitValue <= maxLimit {
			limit = limitValue
		}
	}
	return limit
}
