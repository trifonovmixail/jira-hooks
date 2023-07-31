package objects

import (
	"strconv"
	"time"
)

const (
	dateLayout = "\"2006-01-02\""
	timeLayout = "\"2006-01-02T15:04:05.999-0700\""
)

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse(timeLayout, string(b))
	if err != nil {
		return err
	}
	*t = Time(ti)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	baseTime := time.Time(t)
	if baseTime.IsZero() {
		return []byte("null"), nil
	}
	return []byte(baseTime.Format(timeLayout)), nil
}

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse(dateLayout, string(b))
	if err != nil {
		return err
	}
	*d = Date(ti)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	date := time.Time(d)
	if date.IsZero() {
		return []byte("null"), nil
	}
	return []byte(date.Format(dateLayout)), nil
}

type Timestamp time.Time

func (ts *Timestamp) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	*ts = Timestamp(time.Unix(i, 0))
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(ts).Unix(), 10)), nil
}
