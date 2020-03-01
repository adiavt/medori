package main

import (
"database/sql/driver"
"fmt"
"github.com/dwadp/medori/formats"
"strings"
"time"
)

type Datetime struct {
	time.Time
}

// Implementation for json.Unmarshal function
func (d *Datetime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	trimmed := strings.Trim(string(data), "\"")
	layout := formats.DatetimeLayout(trimmed)

	if trimmed == "" {
		*d = Datetime{time.Now()}
		return nil
	}

	t, err := time.Parse(layout, trimmed)

	if err != nil {
		return err
	}

	*d = Datetime{t}

	return nil
}

// Implementation of json.Marshall function
func (d *Datetime) MarshalJSON() ([]byte, error) {
	t := fmt.Sprintf("\"%s\"", d.Format("2006-01-02 15:04:05"))

	return []byte(t), nil
}

// String returns the time formatted using the format string
func (d *Datetime) String() string {
	return d.Format("2006-01-02 15:04:05")
}

// Implementation of driver.Valuer
func (d Datetime) Value() (driver.Value, error) {
	return d.Format("2006-01-02 15:04:05"), nil
}

// Implementation of sql.Scanner
func (d *Datetime) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		return d.UnmarshalJSON(value)
	case string:
		return d.UnmarshalJSON([]byte(value))
	case time.Time:
		return d.UnmarshalJSON([]byte(value.Format("2006-01-02 15:04:05")))
	}

	return nil
}
