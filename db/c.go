package db

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
)

// U 更新
type C struct {
	TableName            string   `json:"-"`
	Data                 string   `json:"data"`
	OnDuplicateKeyUpdate []string `json:"on_duplicate_key_update"`
}

func (m C) BuildSQL(ctx context.Context) (sql string, err error) {
	if m.TableName == "" {
		return "", errors.New("param TableName is null")
	}

	switch len(m.Data) {
	case 0:
		return "", errors.New("param Data is null")
	}

	// json 字符串转map
	var values []map[string]json.RawMessage
	if err = json.Unmarshal([]byte(m.Data), &values); err != nil {
		return "", err
	}

	// 获取字段
	fields := m.getFields(values)

	var sqlStr strings.Builder
	sqlStr.WriteString("INSERT INTO ")
	sqlStr.WriteString(m.TableName)
	sqlStr.WriteString(" (")
	sqlStr.WriteString(strings.Join(fields, ","))
	sqlStr.WriteString(") VALUES ")

	// 获取第一个
	value := m.getValue(fields, values[0])
	sqlStr.WriteString(value)

	for _, valueMap := range values[1:] {
		sqlStr.WriteString(",")
		// 获取值
		value := m.getValue(fields, valueMap)
		sqlStr.WriteString(value)
	}

	// 更新条件
	if len(m.OnDuplicateKeyUpdate) > 0 {
		updateFields := m.getOnDuplicateKeyUpdate()
		sqlStr.WriteString(updateFields)
	}

	return strings.Trim(sqlStr.String(), " "), nil
}

func (m C) getOnDuplicateKeyUpdate() string {
	var sqlStr strings.Builder
	sqlStr.WriteString(" on duplicate key update ")

	sqlStr.WriteString(m.OnDuplicateKeyUpdate[0])
	sqlStr.WriteString(" = values(")
	sqlStr.WriteString(m.OnDuplicateKeyUpdate[0])
	sqlStr.WriteString(")")

	for _, i2 := range m.OnDuplicateKeyUpdate[1:] {
		sqlStr.WriteString(",")
		sqlStr.WriteString(i2)
		sqlStr.WriteString(" = values(")
		sqlStr.WriteString(i2)
		sqlStr.WriteString(")")
	}
	return sqlStr.String()
}

// getFields 获取字段
func (m C) getFields(values []map[string]json.RawMessage) []string {
	var fields []string
	for k, _ := range values[0] {
		fields = append(fields, k)
	}
	return fields
}

// getValue 获取值
func (m C) getValue(fields []string, valueMap map[string]json.RawMessage) string {
	var valStr strings.Builder
	valStr.WriteString("(")
	valStr.Write(valueMap[fields[0]])

	for _, fieldName := range fields[1:] {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := valueMap[fieldName]; ok {
			valStr.WriteString(",")
			valStr.Write(Value)
		}
	}
	valStr.WriteString(")")
	return valStr.String()
}
