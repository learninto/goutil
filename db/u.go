package db

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
)

// U 更新
type U struct {
	Filters []F    `json:"filters"`
	Data    string `json:"data"`
	SQL     string `json:"sql"` //
}

// BuildSQL 构建SQL
func (m U) BuildSQL(ctx context.Context) (sql string, err error) {
	if m.SQL == "" {
		return "", errors.New("param SQL is null")
	}

	switch len(m.Data) {
	case 0:
		return "", errors.New("param Data is null")
	}

	// json 字符串转map
	var dataMap map[string]json.RawMessage
	if err = json.Unmarshal([]byte(m.Data), &dataMap); err != nil {
		return "", err
	}

	var b strings.Builder
	for k, v := range dataMap {
		b.WriteString(" ")
		b.WriteString(k)
		b.WriteString("=")
		b.Write(v)
		b.WriteString(",")
	}
	replaceValues := strings.TrimRight(b.String(), ",")

	m.SQL = strings.ReplaceAll(m.SQL, "{{.replace_values}}", replaceValues)

	m.SQL, err = F{}.BuildFilter(ctx, m.Filters, m.SQL)
	if err != nil {
		return "", err
	}
	return strings.Trim(m.SQL, " "), nil
}
