package helpers

import "time"

func ConvertTimeToLocal() time.Time {
	now := time.Now()
	rounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return rounded
}
func TimeNow() time.Time {
	return time.Now()
}