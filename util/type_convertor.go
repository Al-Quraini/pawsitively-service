package util

import (
	"database/sql"
	"time"
)

// ****************************** //
// nullable string   //
// ****************************** //

func GetNullableString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

// ****************************** //
// nullable int64   //
// ****************************** //

func GetNullableInt64(i sql.NullInt64) *int64 {
	if i.Valid {
		return &i.Int64
	}
	return nil
}

// ****************************** //
// update time   //
// ****************************** //
func GetNullableTime(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

// ****************************** //
// string to nullable   //
// ****************************** //
func StringToNullString(str string) sql.NullString {
	var nullableStr sql.NullString
	if str == "" {
		nullableStr.Valid = false
	} else {
		nullableStr.String = str
		nullableStr.Valid = true
	}
	return nullableStr
}

// ****************************** //
// int64 to nullable   //
// ****************************** //
func Int64ToNullInt64(n int64) sql.NullInt64 {
	var nullableInt sql.NullInt64
	nullableInt.Int64 = n
	nullableInt.Valid = true
	return nullableInt
}

// ****************************** //
// time to nullable   //
// ****************************** //
func TimeToNullTime(t time.Time) sql.NullTime {
	var nullableTime sql.NullTime
	nullableTime.Time = t
	nullableTime.Valid = true
	return nullableTime
}
