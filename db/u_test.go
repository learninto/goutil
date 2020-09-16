package db

import (
	"context"
	"testing"

	"github.com/learninto/goutil/ctxkit"
)

func TestU_BuildSQL(t *testing.T) {
	ctx := context.TODO()
	ctx = ctxkit.WithCompanyID(ctx, 10000)

	type fields struct {
		Filters []F
		Data    string
		SQL     string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantSql string
		wantErr bool
	}{
		{
			name: "测试-01",
			fields: fields{
				Filters: []F{
					{FieldName: "name", Value: "李四"},
				},
				Data: `{"name":"张三", "sex":"14"}`,
				SQL:  "UPDATE table_name SET {{.replace_values}} WHERE company_id = {{.company_id}} {{.append_filter}}",
			},
			args:    args{ctx: ctx},
			wantErr: false,
			wantSql: `UPDATE table_name SET  name="张三", sex="14" WHERE company_id = 10000  and name  like '%李四%'`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := U{
				Filters: tt.fields.Filters,
				Data:    tt.fields.Data,
				SQL:     tt.fields.SQL,
			}
			gotSql, err := m.BuildSQL(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSql != tt.wantSql {
				t.Errorf("BuildSQL() gotSql = %v, want %v", gotSql, tt.wantSql)
			}
		})
	}
}
