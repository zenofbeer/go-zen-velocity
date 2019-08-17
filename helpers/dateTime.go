package helpers

import (
	"time"
)

const layoutISO = "1900-01-01"

// GetFormattedDate pass in the date string and format it
// for display
func GetFormattedDate(rawDate string) string {
	t, _ := time.Parse(layoutISO, rawDate)
	return t.String()
}
