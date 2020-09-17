package jwt

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"

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
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjoxMjM0NTZ9.oQC5aJRtHlHkxBvOKNj6ne5FFUnznwO8hjdcoWClNjo",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.CreateToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
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
			args:   args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjpudWxsfQ.AKy7KIJnXUwB20EmOoxWn6BGeAskGtnotlLPo10uGbk"},
			want: &CustomClaims{
				Data: json.RawMessage{0x6e, 0x75, 0x6c, 0x6c},
				StandardClaims: jwt.StandardClaims{
					Audience: "", ExpiresAt: 0, Id: "",
					IssuedAt: 0, Issuer: "", NotBefore: 0, Subject: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				pp.Println(got)
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
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.RefreshToken(tt.args.tokenString)
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

func TestNewJWT(t *testing.T) {
	tests := []struct {
		name string
		want *JWT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
