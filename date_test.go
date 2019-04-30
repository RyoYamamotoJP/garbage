package garbage

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestWeekdayOfMonthString(t *testing.T) {
	date := Date{time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC)}
	want := "First"

	s := date.WeekdayOfMonth().String()
	if s != want {
		t.Errorf("WeekdayOfMonth().String() = %v, want %v", s, want)
	}
}

func TestWeekdayOfMonth(t *testing.T) {
	date := Date{time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC)}
	want := First

	wday := date.WeekdayOfMonth()
	if wday != want {
		t.Errorf("WeekdayOfMonth() = %v, want %v", wday, want)
	}
}

func TestMarshalJSON(t *testing.T) {
	date := Date{time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC)}
	want := []byte(`"2019-05-01"`)

	b, err := json.Marshal(date)
	if err != nil {
		t.Errorf("MarshalJSON(%v) error = %v, want nil", date, err)
	} else if !bytes.Equal(b, want) {
		t.Errorf("MarshalJSON(%v) = %s, want = %s", date, b, want)
	}
}

func TestMarshalJSONError(t *testing.T) {
	date := Date{time.Date(10000, time.January, 1, 0, 0, 0, 0, time.UTC)}
	want := "json: error calling MarshalJSON for type garbage.Date: Date.MarshalJSON: year outside of range [0,9999]"

	_, err := json.Marshal(date)
	if err == nil || err.Error() != want {
		t.Errorf("MarshalJSON(%v) error = %v, want %v", date, err, want)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	data := []byte(`"2019-05-01"`)
	want := time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC)

	var date Date
	err := json.Unmarshal(data, &date)
	if err != nil {
		t.Errorf("UnmarshalJSON(%s) error = %v, want nil", data, err)
	} else if !date.Equal(want) {
		t.Errorf("UnmarshalJSON(%s) = %v, want = %v", data, date, want)
	}
}

func TestUnmarshalJSONError(t *testing.T) {
	data := []byte("2019/05/01")
	want := "invalid character '/' after top-level value"

	var date Date
	err := json.Unmarshal(data, &date)
	if err == nil || err.Error() != want {
		t.Errorf("UnmarshalJSON(%s, %T) error = %v, want %v", data, &date, err, want)
	}
}

func TestUnmarshalJSONNull(t *testing.T) {
	data := []byte("null")
	want := Date{}.Time

	var date Date
	err := json.Unmarshal(data, &date)
	if err != nil {
		t.Errorf("UnmarshalJSON(%s) error = %v, want nil", data, err)
	} else if !date.Equal(want) {
		t.Errorf("UnmarshalJSON(%s) = %v, want = %v", data, date, want)
	}
}

func TestParseDate(t *testing.T) {
	value := "2019-05-01"
	want := time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC)

	date, err := Parse(value)
	if err != nil {
		t.Errorf("Parse(%s) error = %v, want nil", value, err)
	} else if !date.Equal(want) {
		t.Errorf("Parse(%s) = %v, want = %v", value, date, want)
	}
}

func TestParseDateError(t *testing.T) {
	value := "2019/05/01"
	want := `parsing time "2019/05/01" as "2006-01-02": cannot parse "/05/01" as "-"`

	_, err := Parse(value)
	if err == nil || err.Error() != want {
		t.Errorf("Parse(%s) error = %v, want %v", value, err, want)
	}
}

func TestNewDate(t *testing.T) {
	year := 2019
	month := time.May
	day := 1
	want := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	date := NewDate(year, month, day, time.UTC)
	if !date.Equal(want) {
		t.Errorf("NewDate(%d, %d, %d) = %v, want = %v", year, month, day, date, want)
	}
}
