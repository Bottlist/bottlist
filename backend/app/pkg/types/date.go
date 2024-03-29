package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(s string) error {
	t, err := time.Parse("2006-01-02", s)
	*d = Date(t)
	return err
}

func (d *Date) Scan(value interface{}) error {
	*d = Date(value.(time.Time))
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d Date) GetDate() string {
	return time.Time(d).Format("2006-01-02")
}
