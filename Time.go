package knx

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) String() string {
	dow := "No Day"
	switch t.Weekday() {
	case time.Monday:
		dow = "Monday"
	case time.Tuesday:
		dow = "Tuesday"
	case time.Wednesday:
		dow = "Wednesday"
	case time.Thursday:
		dow = "Thursday"
	case time.Friday:
		dow = "Friday"
	case time.Saturday:
		dow = "Saturday"
	case time.Sunday:
		dow = "Sunday"
	}
	return fmt.Sprintf("%s %d:%d:%d", dow, t.Hour(), t.Minute(), t.Second())
}

func (t Time) Pack() []byte {
	result := []byte{0, 0, 0}
	// Set day of week
	var dow uint8 = 0
	if t.Weekday() == time.Sunday {
		dow = 7 // Knx encodes Sunday as 7
	} else {
		dow = uint8(t.Weekday())
	}
	result[0] |= (dow << 5)
	// Set Hour
	result[0] |= (uint8(t.Hour()) & 0b00011111)
	// Set Minutes
	result[1] |= (uint8(t.Minute() & 0b00111111))
	// Set Seconds
	result[2] |= (uint8(t.Second() & 0b00111111))
	// Return result
	return result
}

func NewTime(data []byte) (Time, error) {
	if len(data) != 3 {
		return Time{}, NewErrInvalidDataType("given data cannot be unpacked to Time value")
	}
	// Extract DayOfWeek
	dow := uint8((data[0] & 0b11100000) >> 5)
	if dow == 7 {
		dow = 0 // Convert sunday from 7 to 0
	}
	// Extract Hour
	h := uint8(data[0] & 0b00011111)
	// Extract Minutes
	m := uint8(data[1] & 0b00111111)
	// Extract Seconds
	s := uint8(data[2] & 0b00111111)

	// Calculate date for weekday
	now := time.Now()
	currentDay := now.Weekday()
	offset := int(dow - uint8(currentDay))
	date := now.AddDate(0, 0, offset)
	// Create final timestamp
	t := time.Date(date.Year(), date.Month(), date.Day(), int(h), int(m), int(s), 0, time.UTC)
	return Time{t}, nil
}
