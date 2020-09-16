package db

import (
	"bytes"
	"context"
	"errors"
	"html"
	"html/template"
	"strings"
	"time"

	"github.com/learninto/goutil/ctxkit"
)

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

// BuildFilter 获取追加的筛选条件
func (f F) BuildFilter(ctx context.Context, filters []F, sql string) (string, error) {
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
			appendFilter += f.buildFilter(filter)
		}
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
	}
	tmpl, err := template.New("tmpl").Parse(sql)
	if err != nil {
		return "", err
	}

	tmplBytes := &bytes.Buffer{}
	if err = tmpl.Execute(tmplBytes, tempVars); err != nil {
		return "", err
	}

	bytesSQL := tmplBytes.String()
	if bytesSQL == "" {
		return "", errors.New("build F is null")
	}

	return html.UnescapeString(bytesSQL), nil
}

// 构造单个筛选
func (F) buildFilter(f F) (res string) {
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
	default:
		break
	}
	return
}
