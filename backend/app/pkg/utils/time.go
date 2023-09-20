package utils

import "time"

var (
	jst, _ = time.LoadLocation("Asia/Tokyo")
)

func Now() time.Time {
	now := time.Now().In(jst)
	return now
}

func GetTimeDelay(lifetime time.Duration) time.Time {
	expiredAt := time.Now().In(jst).Add(lifetime)
	return expiredAt
}

func GetHourDuration(hours int) time.Duration {
	return time.Duration(hours) * time.Hour
}

func GetMinuteDuration(minutes int) time.Duration {
	return time.Duration(minutes) * time.Minute
}
