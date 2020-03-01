package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/dwadp/medori/formats"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

// Implementation for json.Unmarshal function
func (d *Date) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	trimmed := strings.Trim(string(data), "\"")
	layout := formats.DateLayout(trimmed)

	if trimmed == "" {
		*d = Date{time.Now()}
		return nil
	}

	t, err := time.Parse(layout, trimmed)

	if err != nil {
		return err
	}

	*d = Date{t}

	return nil
}

// Implementation of json.Marshall function
func (d *Date) MarshalJSON() ([]byte, error) {
	t := fmt.Sprintf("\"%s\"", d.Format("2006-01-02"))

	return []byte(t), nil
}

// String returns the time formatted using the format string
func (d *Date) String() string {
	return d.Format("2006-01-02")
}

// Implementation of driver.Valuer
func (d Date) Value() (driver.Value, error) {
	return d.Format("2006-01-02"), nil
}

// Implementation of sql.Scanner
func (d *Date) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		return d.UnmarshalJSON(value)
	case string:
		return d.UnmarshalJSON([]byte(value))
	case time.Time:
		return d.UnmarshalJSON([]byte(value.Format("2006-01-02")))
	}

	return nil
}
