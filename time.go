package main

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

// Implementation for json.Unmarshal function
func (d *Time) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	trimmed := strings.Trim(string(data), "\"")
	layout := "15:04:05"

	if trimmed == "" {
		*d = Time{time.Now()}
		return nil
	}

	t, err := time.Parse(layout, trimmed)

	if err != nil {
		return err
	}

	*d = Time{t}

	return nil
}

// Implementation of json.Marshall function
func (d *Time) MarshalJSON() ([]byte, error) {
	t := fmt.Sprintf("\"%s\"", d.Format("15:04:05"))

	return []byte(t), nil
}

// String returns the time formatted using the format string
func (d *Time) String() string {
	return d.Format("15:04:05")
}

// Implementation of driver.Valuer
func (d Time) Value() (driver.Value, error) {
	return d.Format("15:04:05"), nil
}

// Implementation of sql.Scanner
func (d *Time) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		return d.UnmarshalJSON(value)
	case string:
		return d.UnmarshalJSON([]byte(value))
	case time.Time:
		return d.UnmarshalJSON([]byte(value.Format("15:04:05")))
	}

	return nil
}