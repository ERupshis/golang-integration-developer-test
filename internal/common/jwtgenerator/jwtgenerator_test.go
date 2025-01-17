package jwtgenerator

import (
	"testing"
)

func TestJwtGenerator_Overall(t *testing.T) {
	type fields struct {
		jwtKey   string
		tokenExp int
	}
	type args struct {
		tokenStringSuffix string
		userID            int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				jwtKey:   "secret",
				tokenExp: 1,
			},
			args: args{
				userID: 3,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "expired token",
			fields: fields{
				jwtKey:   "secret",
				tokenExp: 0,
			},
			args: args{
				userID: 3,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "invalid token",
			fields: fields{
				jwtKey:   "secret",
				tokenExp: 0,
			},
			args: args{
				tokenStringSuffix: "af10",
				userID:            3,
			},
			want:    -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JwtGenerator{
				jwtKey:   tt.fields.jwtKey,
				tokenExp: tt.fields.tokenExp,
			}
			tokenString, err := j.BuildJWTString(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildJWTString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tokenString += tt.args.tokenStringSuffix

			if userID, _ := j.GetUserID(tokenString); userID != tt.want {
				t.Errorf("GetUserID() = %v, want %v", userID, tt.want)
			}
		})
	}
}
