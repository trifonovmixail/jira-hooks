package objects

import (
	"fmt"
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

func (t *Time) MarshalJSON() ([]byte, error) {
	baseTime := time.Time(*t)
	if baseTime.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, baseTime.Format(timeLayout))), nil
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

func (d *Date) MarshalJSON() ([]byte, error) {
	date := time.Time(*d)
	if date.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, date.Format(dateLayout))), nil
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(i, 0))
	return nil
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(*t).Unix(), 10)), nil
}
