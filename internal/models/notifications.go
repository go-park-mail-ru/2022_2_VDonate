package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type DataMap map[string]interface{}

type NotificationCancel struct {
	Cancel bool `json:"cancel"`
}

type Notification struct {
	Name      string    `json:"name"`
	Data      DataMap   `json:"data"`
	Timestamp time.Time `json:"time"`
}

func (p DataMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *DataMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
