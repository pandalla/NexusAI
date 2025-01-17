package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func GetTime() time.Time {
	return time.Now()
}

func GetTimeStamp() int64 { // 获取当前时间戳
	return time.Now().Unix()
}

func GetTimeString() string { // 获取当前时间字符串
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), time.Now().UnixNano()%1e9)
}

func GetTimeStringMySQL() string { // 获取当前时间字符串，格式为MySQL的datetime类型
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatTimeToMySQL(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

type MySQLTime time.Time

// Value 实现 driver.Valuer 接口
func (t MySQLTime) Value() (driver.Value, error) {
	return time.Time(t).Format("2006-01-02 15:04:05"), nil
}

// Scan 实现 sql.Scanner 接口
func (t *MySQLTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = MySQLTime(v.Truncate(time.Second))
	case []byte:
		if len(v) == 0 {
			return nil
		}
		// 先尝试解析带微秒的格式
		parsedTime, err := time.Parse("2006-01-02 15:04:05.999999", string(v))
		if err != nil {
			// 如果失败，尝试标准格式
			parsedTime, err = time.Parse("2006-01-02 15:04:05", string(v))
			if err != nil {
				return err
			}
		}
		// 去除微秒部分
		*t = MySQLTime(parsedTime.Truncate(time.Second))
	case string:
		if v == "" {
			return nil
		}
		// 先尝试解析带微秒的格式
		parsedTime, err := time.Parse("2006-01-02 15:04:05.999999", v)
		if err != nil {
			// 如果失败，尝试标准格式
			parsedTime, err = time.Parse("2006-01-02 15:04:05", v)
			if err != nil {
				return err
			}
		}
		// 去除微秒部分
		*t = MySQLTime(parsedTime.Truncate(time.Second))
	default:
		return fmt.Errorf("无法将 %T 转换为 MySQLTime", value)
	}
	return nil
}

// MarshalJSON 实现 json.Marshaler 接口
func (t MySQLTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *MySQLTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tt, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}
	*t = MySQLTime(tt)
	return nil
}

// String 实现 Stringer 接口
func (t MySQLTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}
