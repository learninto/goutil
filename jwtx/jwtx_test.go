package jwtx

import (
	"context"
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
)

func TestGetSignKey(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
		{"default", []byte("gwt_sign_key")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSignKey(); string(got) != string(tt.want) {
				t.Errorf("GetSignKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_CreateToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		claims CustomClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case 01",
			fields{SigningKey: []byte("gwt_sign_key")},
			args{CustomClaims{Data: []byte("123456")}},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.args.claims.CreateToken(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if got != tt.want {
			//	t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestJWT_ParseToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomClaims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "case 01",
			fields: fields{SigningKey: []byte("gwt_sign_key")},
			args:   args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjoxMjM0NTYsImV4cCI6MTY0NTI2MDIxM30.cON6OEoVqwxncQI9WcRRKJPIp2gCEuCMJczAVR2prdY"},
			want: &CustomClaims{
				Data: []byte("123456"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CustomClaims{}.ParseToken(context.Background(), tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				_, _ = pp.Println(got)
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_RefreshToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case 01",
			fields{SigningKey: []byte("gwt_sign_key")},
			args{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuYW1lIjoi5byg5LiJIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJydWxlcyI6bnVsbH0.agleKMaE-ncgJetG8jGU4eLMlNsCBZN4CyN2pOSht4o"},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjpudWxsfQ.AKy7KIJnXUwB20EmOoxWn6BGeAskGtnotlLPo10uGbk",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := CustomClaims{}
			got, err := j.RefreshToken(context.TODO(), tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("RefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
