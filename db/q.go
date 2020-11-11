package db

import (
	"bytes"
	"context"
	"errors"
	"html"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/learninto/goutil/ctxkit"
)

// FieldTypeText 文本
const FieldTypeText = 0

// FieldTypeDate 日期
const FieldTypeDate = 100

// FieldTypeDateTime 日期时间
const FieldTypeDateTime = 101

// FieldTypePositiveInteger 正整数（包含0）
const FieldTypePositiveInteger = 200

// FieldTypeInteger 整数
const FieldTypeInteger = 201

// FieldTypePositiveRealNumber 正实数（包含0）
const FieldTypePositiveRealNumber = 200

// FieldTypeRealNumber 实数
const FieldTypeRealNumber = 301

// FieldTypeListOne 列表（只能单选）
const FieldTypeListOne = 400

// FieldTypeListMultiple 列表（可多选）
const FieldTypeListMultiple = 401

// FieldLogicLike 过滤逻辑 0：like ；
const FieldLogicLike = 0

// FieldLogicLeftLike 过滤逻辑 1：左 like；
const FieldLogicLeftLike = 1

// FieldLogicRightLike 过滤逻辑 2：右 like；
const FieldLogicRightLike = 2

// FieldLogicEq 过滤逻辑 100：= ；
const FieldLogicEq = 100

// FieldLogicGt 过滤逻辑 200：>;
const FieldLogicGt = 200

// FieldLogicEgt 过滤逻辑 201：>=
const FieldLogicEgt = 201

// FieldLogicLt 过滤逻辑 300 <；
const FieldLogicLt = 300

// FieldLogicElt 过滤逻辑 301: <=；
const FieldLogicElt = 301

// FieldLogicNeq 过滤逻辑 400： <>；
const FieldLogicNeq = 400

// FieldLogicIn 过滤逻辑 500： in
const FieldLogicIn = 500

// FieldLogicNotIn 过滤逻辑 500： not in
const FieldLogicNotIn = 501

// IsMultipleTrue 是否多重应用条件。是指在sql中，该过滤值在多个子查询中需要用到。 0：否；100：是
const IsMultipleTrue = 100

// IsMultipleFalse 是否多重应用条件。是指在sql中，该过滤值在多个子查询中需要用到。 0：否；100：是
const IsMultipleFalse = 0

// Q 查询
type Q struct {
	Sql            string `json:"sql"`
	Filters        []F    `json:"filters"`
	ReplaceOrderBy string `json:"replace_order_by"`
	ReplaceGroupBy string `json:"replace_group_by"`
	CurrentPage    int64  `json:"current_page"`
	PageSize       int64  `json:"page_size"`
}

// F 筛选条件
type F struct {
	// Comment: 字段名称。 用于后端 构造sql。  如果是多重应用条件，则成为占位名称标识
	FieldName string `json:"field_name"`
	// Comment: 字段类型。 用于前端 不同输入控件，以及限制输入的内容。 0：文本；100： 日期  101：日期时间；200：正整数（包含0）；201：整数；300：正实数（包含0）；301：实数； 400：列表（只能单选）；401：列表（可多选）
	FieldType int64 `json:"field_type"`
	// Comment: 过滤逻辑。 用于前端：显示查询条件逻辑； 用于后端 构造sql。 0：like ；1：左 like；2：右 like；100：= ；200：>; 201：>=  300 <；301: <=； 400： <>；500： in
	FieldLogic int64 `json:"field_logic"`
	// Comment: 是否多重应用条件。是指在sql中，该过滤值在多个子查询中需要用到。 0：否；100：是
	// Default: 0
	IsMultiple int64 `json:"is_multiple"`
	// Comment：值
	Value string `json:"value"`
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

	m.Sql, err = m.BuildFilter(ctx, m.Filters, m.Sql)

	return m.Sql, err
}

// BuildFilter 获取追加的筛选条件
func (m Q) BuildFilter(ctx context.Context, filters []F, sql string) (string, error) {
	appendFilter := ""
	for _, filter := range filters {

		// 如果是文本类型， filter.Value 需要转换单引号
		filterValue := filter.Value

		if filter.FieldType == FieldTypeText {
			filterValue = strings.ReplaceAll(filterValue, "'", "''")
		}

		// 多重应用条件
		if filter.IsMultiple == IsMultipleTrue {
			sql = strings.ReplaceAll(sql, filter.FieldName, filterValue)
		} else { // 非多重应用条件
			appendFilter += filterLogic(filter)
		}
	}

	// 追加的分页条件
	replaceLimit := ""
	if m.PageSize > 0 && m.CurrentPage >= 0 {
		var b strings.Builder
		b.WriteString(" limit ")
		b.WriteString(strconv.FormatInt((m.CurrentPage-1)*m.PageSize, 10))
		b.WriteString(",")
		b.WriteString(strconv.FormatInt(m.PageSize, 10))
		replaceLimit = b.String()
	}

	t := time.Now()
	curDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Unix()
	tempVars := map[string]interface{}{
		"user_id":        ctxkit.GetUserID(ctx),        // 用户id列表
		"part_ids":       ctxkit.GetPartIds(ctx),       // 角色id列表
		"department_ids": ctxkit.GetDepartmentIds(ctx), // 部门id列表
		"company_id":     ctxkit.GetCompanyID(ctx),     // 公司id
		"cur_date":       curDate,                      // 当天开始时间戳值
		"yesterday":      curDate - 86400,              // 昨天开始时间戳值
		"tomorrow":       curDate + 86400,              // 明天开始间戳值
		"append_filter":  appendFilter,                 // 追加的筛选条件
		"replace_limit":  replaceLimit,                 // 追加的筛选条件
	}
	tmpl, err := template.New("tmpl").Parse(sql)
	if err != nil {
		return "", err
	}

	tmplBytes := &bytes.Buffer{}
	if err = tmpl.Execute(tmplBytes, tempVars); err != nil {
		return "", err
	}

	sql = tmplBytes.String()
	if sql == "" {
		return "", errors.New("build F is null")
	}

	sql = html.UnescapeString(sql)
	sql = strings.ReplaceAll(sql, "<no value>", "")

	return sql, nil
}

// 构造单个筛选
func filterLogic(f F) (res string) {
	// 拼接sql FieldLogic —— 0：like;1：左 like；2：右 like；100：=;200：>; 201：>=  300 <；301: <=； 400： <>；500： in
	res = res + " and " + f.FieldName + " "

	switch f.FieldLogic {
	case FieldLogicLike:
		res = res + " like '%" + f.Value + "%' "
		break
	case FieldLogicLeftLike:
		res = res + " like '%" + f.Value + "' "
		break
	case FieldLogicRightLike:
		res = res + " like '" + f.Value + "%' "
		break
	case FieldLogicEq:
		res = res + " = " + f.Value + " "
		break
	case FieldLogicGt:
		res = res + " > " + f.Value + " "
		break
	case FieldLogicEgt:
		res = res + " >= " + f.Value + " "
		break
	case FieldLogicLt:
		res = res + " < " + f.Value + " "
		break
	case FieldLogicElt:
		res = res + " <= " + f.Value + " "
		break
	case FieldLogicNeq:
		res = res + " <> " + f.Value + " "
		break
	case FieldLogicIn:
		str := " in ("
		typeArr := strings.Split(f.Value, ",")
		for _, v := range typeArr {
			str = str + "'" + v + "'" + ","
		}
		str = strings.TrimRight(str, ",") + ")"
		res = res + str
		break
	case FieldLogicNotIn:
		str := " not in ("
		typeArr := strings.Split(f.Value, ",")
		for _, v := range typeArr {
			str = str + "'" + v + "'" + ","
		}
		str = strings.TrimRight(str, ",") + ")"
		res = res + str
		break
	default:
		break
	}
	return
}
