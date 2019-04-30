package garbage

import (
	"errors"
	"time"
)

type WeekdayOfMonth int

const (
	First WeekdayOfMonth = 1 + iota
	Second
	Third
	Fourth
	Last
)

var weekdaysOfMonth = [...]string{
	"First",
	"Second",
	"Third",
	"Fourth",
	"Last",
}

func (d WeekdayOfMonth) String() string {
	return weekdaysOfMonth[d-1]
}

type Date struct {
	time.Time
}

func (d Date) WeekdayOfMonth() WeekdayOfMonth {
	wday := (d.Day()-1)/7 + 1
	return WeekdayOfMonth(wday)
}

const RFC3339FullDate = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(RFC3339FullDate)+2)
	b = append(b, '"')
	b = d.AppendFormat(b, RFC3339FullDate)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t, err := time.Parse(`"`+RFC3339FullDate+`"`, string(data))
	d.Time = t
	return err
}

func Parse(value string) (Date, error) {
	t, err := time.Parse(RFC3339FullDate, value)
	if err != nil {
		return Date{}, err
	}

	return Date{t}, nil
}

func NewDate(year int, month time.Month, day int, loc *time.Location) *Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return &Date{t}
}
