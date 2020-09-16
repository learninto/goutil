package db

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// Q 查询
type Q struct {
	Sql            string `json:"sql"`
	Filters        []F    `json:"filters"`
	ReplaceOrderBy string `json:"replace_order_by"`
	ReplaceGroupBy string `json:"replace_group_by"`
	CurrentPage    int64  `json:"current_page"`
	PageSize       int64  `json:"page_size"`
}

// BuildSql 构建SQL
func (m Q) BuildSql(ctx context.Context) (sql string, err error) {
	if m.Sql == "" {
		return "", errors.New("params SQL is null")
	}

	// 追加的排序条件
	if m.ReplaceOrderBy != "" {
		var b strings.Builder
		b.WriteString(" order by ")
		b.WriteString(m.ReplaceOrderBy)
		m.Sql = strings.ReplaceAll(m.Sql, "{{.replace_order_by}}", b.String())
	}

	// 追加的分组条件
	if m.ReplaceGroupBy != "" {
		var b strings.Builder
		b.WriteString(" group by ")
		b.WriteString(m.ReplaceGroupBy)
		m.Sql = strings.ReplaceAll(m.Sql, "{{.replace_group_by}}", b.String())
	}

	// 追加的分页条件
	if m.PageSize > 0 && m.CurrentPage >= 0 {
		var b strings.Builder
		b.WriteString(" limit ")
		b.WriteString(strconv.FormatInt((m.CurrentPage-1)*m.PageSize, 10))
		b.WriteString(",")
		b.WriteString(strconv.FormatInt(m.PageSize, 10))
		m.Sql = strings.ReplaceAll(m.Sql, "{{.replace_limit}}", b.String())
	}

	m.Sql, err = F{}.BuildFilter(ctx, m.Filters, m.Sql)

	return m.Sql, err
}
