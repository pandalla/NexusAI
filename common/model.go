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

// MarshalJSON 实现json.Marshaler接口
func (j JSON) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON 实现json.Unmarshaler接口
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull 检查JSON值是否为null
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// String 实现Stringer接口，方便打印调试
func (j JSON) String() string {
	if len(j) == 0 {
		return "null"
	}
	return string(j)
}

// FromStruct 将任意结构体序列化为JSON类型
func FromStruct(v interface{}) (JSON, error) {
	if v == nil {
		return nil, nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return JSON(data), nil
}

// ToStruct 将JSON数据反序列化为指定的结构体
// dest 必须是指针类型
func (j JSON) ToStruct(dest interface{}) error {
	if len(j) == 0 {
		return nil
	}
	return json.Unmarshal(j, dest)
}
