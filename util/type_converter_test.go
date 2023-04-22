package util

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetNullableString(t *testing.T) {
	// test with valid string
	s := sql.NullString{String: "foo", Valid: true}
	result := GetNullableString(s)
	require.NotNil(t, result)
	require.Equal(t, "foo", *result)

	// test with invalid string
	s = sql.NullString{String: "", Valid: false}
	result = GetNullableString(s)
	require.Nil(t, result)
}

func TestGetNullableInt64(t *testing.T) {
	var validValue sql.NullInt64 = sql.NullInt64{Int64: 123, Valid: true}
	var invalidValue sql.NullInt64 = sql.NullInt64{Int64: 0, Valid: false}

	// Test with valid input
	result := GetNullableInt64(validValue)
	expected := int64(123)
	if *result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, *result)
	}

	// Test with invalid input
	result = GetNullableInt64(invalidValue)
	if result != nil {
		t.Errorf("Expected result to be nil, but got %v", *result)
	}
}

func TestGetNullableTime(t *testing.T) {
	var validValue sql.NullTime = sql.NullTime{Time: time.Now(), Valid: true}
	var invalidValue sql.NullTime = sql.NullTime{Time: time.Time{}, Valid: false}

	// Test with valid input
	result := GetNullableTime(validValue)
	expected := validValue.Time
	if !result.Equal(expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, *result)
	}

	// Test with invalid input
	result = GetNullableTime(invalidValue)
	if result != nil {
		t.Errorf("Expected result to be nil, but got %v", *result)
	}
}

func TestStringToNullString(t *testing.T) {
	var validString string = "hello"
	var emptyString string = ""

	// Test with non-empty string
	result := StringToNullString(validString)
	if !result.Valid {
		t.Errorf("Expected result to be valid, but got invalid")
	}
	if result.String != validString {
		t.Errorf("Expected result to be %v, but got %v", validString, result.String)
	}

	// Test with empty string
	result = StringToNullString(emptyString)
	if result.Valid {
		t.Errorf("Expected result to be invalid, but got valid")
	}
}

func TestInt64ToNullInt64(t *testing.T) {
	var validValue int64 = 123

	// Test with valid input
	result := Int64ToNullInt64(validValue)
	if !result.Valid {
		t.Errorf("Expected result to be valid, but got invalid")
	}
	if result.Int64 != validValue {
		t.Errorf("Expected result to be %v, but got %v", validValue, result.Int64)
	}
}

func TestTimeToNullTime(t *testing.T) {
	var validTime time.Time = time.Now()

	// Test with valid input
	result := TimeToNullTime(validTime)
	if !result.Valid {
		t.Errorf("Expected result to be valid, but got invalid")
	}
	if !result.Time.Equal(validTime) {
		t.Errorf("Expected result to be %v, but got %v", validTime, result.Time)
	}
}
