package localtime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"

// LocalTime 本地时间
type LocalTime struct {
	time.Time
}

// UnmarshalJSON gin bind 反射结构体
func (t *LocalTime) UnmarshalJSON(bytes []byte) (err error) {
	t.Time, err = time.ParseInLocation(DateTimeFormat, strings.Trim(string(bytes), "\""), time.Local)
	return
}

// MarshalJSON gorm marshal 序列化结构体
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(DateTimeFormat))), nil
}

// Value LocalTime 转 time
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan gorm Scan 扫描时的数据赋值
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
