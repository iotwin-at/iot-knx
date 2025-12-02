package knx

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) String() string {
	return fmt.Sprintf("%d.%d.%d", d.Day(), d.Month(), d.Year())
}

func (d Date) Pack() []byte {
	result := []byte{0, 0, 0}
	// Encode day
	result[0] |= (uint8(d.Day()) & 0b00011111)
	// Encode month
	result[1] |= (uint8(d.Month()) & 0b00001111)
	// Encode year (0..99)
	result[2] |= (uint8(d.Year()%100) & 0b01111111)
	return result
}

func NewDate(data []byte) (Date, error) {
	if len(data) != 3 {
		return Date{}, NewErrInvalidDataType("given data cannot be unpacked to Date value")
	}
	// Extract day
	day := uint8(data[0] & 0b00011111)
	// Extract month
	month := uint8(data[1] & 0b00001111)
	// Extract year
	var year = 0
	yearRaw := uint8(data[2] & 0b01111111)
	if yearRaw < 90 { // interpret as 20th century
		year = 2000 + int(yearRaw)
	} else { // interpret as 19th century
		year = 1900 + int(yearRaw)
	}
	// Return result
	return Date{time.Date(year, time.Month(month), int(day), 0, 0, 0, 0, time.UTC)}, nil
}
