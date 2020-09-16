package db

import (
	"context"
	"testing"

	"github.com/learninto/goutil/ctxkit"
)

func TestC_BuildSQL(t *testing.T) {
	ctx := context.TODO()
	ctx = ctxkit.WithCompanyID(ctx, 10000)
	type fields struct {
		TableName            string
		Data                 string
		OnDuplicateKeyUpdate []string
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
				TableName:            "table_name",
				Data:                 `[{"name":"张三","sex": 14}, {"name":"李四","sex": 15}]`,
				OnDuplicateKeyUpdate: []string{"name", "sex"},
			},
			args:    args{ctx: ctx},
			wantErr: false,
			wantSql: `INSERT INTO table_name (name,sex) VALUES ("张三",14),("李四",15) on duplicate key update name = values(name),sex = values(sex)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := C{
				TableName:            tt.fields.TableName,
				Data:                 tt.fields.Data,
				OnDuplicateKeyUpdate: tt.fields.OnDuplicateKeyUpdate,
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
