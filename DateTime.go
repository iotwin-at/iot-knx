package knx

import (
	"time"
)

type DateTime struct {
	time.Time
	Fault              bool
	WorkingDay         bool
	WorkingDayInvalid  bool
	YearInvalid        bool
	MonthAndDayInvalid bool
	DayOfWeekInvalid   bool
	TimeInvalid        bool
	ClockWithExtSync   bool
	ReliableSyncSource bool
}

func (d DateTime) String() string {
	return d.Format(time.RFC3339)
}

func (d DateTime) Pack() []byte {
	result := make([]byte, 8)
	// Encode year
	yearRaw := d.Year()
	var year uint8
	if yearRaw <= 1900 {
		year = 0
	} else if yearRaw >= 2155 {
		year = 255
	} else {
		year = uint8(yearRaw - 1900)
	}
	result[0] |= year
	// Encode Month
	result[1] |= (uint8(d.Month()) & 0b0000_1111)
	// Encode DayOfMonth
	result[2] |= (uint8(d.Day()) & 0b0001_1111)
	// Encode DayOfWeek
	wd := d.Weekday()
	if wd == 0 { // Sunday
		result[3] |= 0b1110_0000 // 7..Sunday
	} else {
		result[3] |= ((uint8(d.Weekday()) & 0b0000_0111) << 5)
	}
	// Encode Hour
	result[3] |= (uint8(d.Hour()) & 0b0001_1111)
	// Encode Minutes
	result[4] |= (uint8(d.Minute()) & 0b0011_1111)
	// Encode Seconds
	result[5] |= (uint8(d.Second()) & 0b0011_1111)
	// Encode F
	if d.Fault {
		result[6] |= 0b1000_0000
	}
	// Encode WD
	if d.WorkingDay {
		result[6] |= 0b0100_0000
	}
	// Encode NWD
	if d.WorkingDayInvalid {
		result[6] |= 0b0010_0000
	}
	// Encode NY
	if d.YearInvalid {
		result[6] |= 0b0001_0000
	}
	// Encode ND
	if d.MonthAndDayInvalid {
		result[6] |= 0b0000_1000
	}
	// Encode NDoW
	if d.DayOfWeekInvalid {
		result[6] |= 0b0000_0100
	}
	// Encode NT
	if d.TimeInvalid {
		result[6] |= 0b0000_0010
	}
	// Encode SUTI
	if d.IsDST() {
		result[6] |= 0b0000_0001
	}
	// Encode CLQ
	if d.ClockWithExtSync {
		result[7] |= 0b1000_0000
	}
	// Encode SRC
	if d.ReliableSyncSource {
		result[7] |= 0b0100_0000
	}

	return result
}

func NewDateTime(data []byte) (DateTime, error) {
	if len(data) != 8 {
		return DateTime{}, NewErrInvalidDataType("given data cannot be unpacked to DateTime value")
	}
	// Extract year
	year := int(data[0]) + 1900 // Value binary encoded, offset 1900
	// Extract Month
	month := int(data[1] & 0b0000_1111)
	// Extract DayOfMonth
	day := int(data[2] & 0b0001_1111)
	// Extract Hour
	h := int((data[3] & 0b0001_1111))
	// Extract Minutes
	m := int(data[4] & 0b0011_1111)
	// Extract Seconds
	s := int(data[5] & 0b0011_1111)
	// Extract F
	f := ((data[6] & 0b1000_0000) >> 7) == 1
	// Extract WD
	wd := ((data[6] & 0b0100_0000) >> 6) == 1
	// Extract NWD
	nwd := ((data[6] & 0b0010_0000) >> 5) == 1
	// Extract NY
	ny := ((data[6] & 0b0001_0000) >> 4) == 1
	// Extract ND
	nd := ((data[6] & 0b0000_1000) >> 3) == 1
	// Extract NDoW
	ndow := ((data[6] & 0b0000_0100) >> 2) == 1
	// Extract NT
	nt := ((data[6] & 0b0000_0010) >> 1) == 1
	// Extract CLQ
	clq := ((data[7] & 0b1000_0000) >> 7) == 1
	// Extract SRC
	src := ((data[7] & 0b0100_0000) >> 6) == 1
	// Return
	return DateTime{
		Time:               time.Date(year, time.Month(month), day, h, m, s, 0, time.UTC),
		Fault:              f,
		WorkingDay:         wd,
		WorkingDayInvalid:  nwd,
		YearInvalid:        ny,
		MonthAndDayInvalid: nd,
		DayOfWeekInvalid:   ndow,
		TimeInvalid:        nt,
		ClockWithExtSync:   clq,
		ReliableSyncSource: src,
	}, nil
}
