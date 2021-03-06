package global

import (
	"strconv"
	"time"
)

// Yet to de developeds

// GetCurrentDateTimeInStr is to get the current date and time in string format
func GetCurrentDateTimeInStr() string {
	return time.Now().String()
}

// GetCurrentDateTimeInTime is to get the current date and time in string format
func GetCurrentDateTimeInTime() time.Time {
	return time.Now()
}

// GetUnixTimeInLong is to get the Unix time in string format
func GetUnixTimeInStr() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10)
}

// GetUnixTimeInLong is to get the Unix time in int64 format
func GetUnixTimeInLong() int64 {
	return time.Now().UTC().Unix()
}

// GetDefaultStr is to get the default value which is of string type
func GetDefaultStr(in string) string {
	return in
}

// GetDefaultInt is to get the default integer value
func GetDefaultInt(in int32) int32 {
	return in
}

// GetDefaultLong is to get the default int64 value
func GetDefaultLong(in int64) int64 {
	return in
}

// GetDefaultBool is to get the default bool value
func GetDefaultBool(in bool) bool {
	return in
}
