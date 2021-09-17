package main

import (
	"fmt"
)

func toHuman(count int) string {
	if count < 1000 {
		return fmt.Sprintf("%d", count)
	}
	if 1000 < count && count < 1000000 {
		return fmt.Sprintf("%.1fK", float64(count)/1000.0)
	}
	if 1000000 < count && count < 1000000000 {
		return fmt.Sprintf("%.1fM", float64(count)/1000000.0)
	}
	if 1000000000 < count && count < 1000000000000 {
		return fmt.Sprintf("%.1fG", float64(count)/1000000000.0)
	}
	return fmt.Sprintf("%.1fT", float64(count/1000000000000.0))
}
