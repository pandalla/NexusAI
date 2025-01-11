package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSON 自定义JSON类型，基于json.RawMessage实现
// 用于在数据库操作和JSON序列化/反序列化时灵活处理JSON数据
// 支持存储任意有效的JSON数据结构（对象、数组、字符串等）
type JSON json.RawMessage

// Value 实现driver.Valuer接口，用于将JSON数据保存到数据库
// 返回值：
//   - 如果JSON为空，返回nil
//   - 否则返回JSON字符串
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

// Scan 实现sql.Scanner接口，用于从数据库扫描JSON数据
// 参数：
//   - value: 数据库中的原始数据
//
// 返回值：
//   - 如果输入为nil，将JSON设置为nil
//   - 如果输入类型不是[]byte，返回错误
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source: expected []byte")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON 实现json.Marshaler接口，用于JSON序列化
// 返回值：
//   - 如果JSON为nil，返回"null"
//   - 否则返回原始JSON数据
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	// 验证JSON数据的有效性
	if !json.Valid(j) {
		return nil, errors.New("invalid json data")
	}
	return j, nil
}

// UnmarshalJSON 实现json.Unmarshaler接口，用于JSON反序列化
// 参数：
//   - data: 要解析的JSON数据
//
// 返回值：
//   - 如果接收者为nil，返回错误
//   - 如果输入的JSON数据无效，返回错误
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null pointer exception")
	}
	// 验证输入的JSON数据是否有效
	if !json.Valid(data) {
		return errors.New("invalid json data")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
