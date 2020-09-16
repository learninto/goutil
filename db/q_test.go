package db

import (
	"context"
	"testing"
)

func TestQ_BuildSql(t *testing.T) {
	type fields struct {
		Sql            string
		Filters        []F
		ReplaceOrderBy string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{{
		name: "Case01",
		fields: fields{
			Sql:            "select * from coo_user_sign Where company_id={{.company_id}} {{.append_filter}} {{.replace_order_by}}",
			ReplaceOrderBy: "create_time desc",
			Filters: []F{
				{
					FieldName:  "create_time",
					FieldType:  0,
					FieldLogic: 201,
					IsMultiple: 0,
					Value:      "1592150400",
				},
				{
					FieldName:  "create_time",
					FieldType:  100,
					FieldLogic: 301,
					IsMultiple: 0,
					Value:      "1592236799",
				},
			},
		},
		args:    args{ctx: context.Background()},
		want:    "select * from coo_user_sign Where company_id=0  and create_time  >= 1592150400  and create_time  <= 1592236799   order by create_time desc",
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Q{
				Sql:            tt.fields.Sql,
				Filters:        tt.fields.Filters,
				ReplaceOrderBy: tt.fields.ReplaceOrderBy,
			}
			got, err := q.BuildSql(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSql() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildSql() got = %v, want %v", got, tt.want)
			}
		})
	}
}
