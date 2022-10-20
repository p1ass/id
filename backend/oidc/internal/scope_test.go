package internal_test

import (
	"reflect"
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewScope(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    internal.Scope
		wantErr bool
	}{
		{
			name: "should return OpenId when openid",
			args: args{
				str: "openid",
			},
			want:    internal.OpenId,
			wantErr: false,
		},
		{
			name: "should return error when unknown value",
			args: args{
				str: "unknown",
			},
			want:    internal.Unknown,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := internal.NewScope(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewScope() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewScopes(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    internal.Scopes
		wantErr bool
	}{
		{
			name: "should return scopes when all args are valid",
			args: args{
				strs: []string{"openid", "email"},
			},
			want:    []internal.Scope{internal.OpenId, internal.Email},
			wantErr: false,
		},
		{
			name: "should return error when one of the args is invalid",
			args: args{
				strs: []string{"openid", "invalid"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := internal.NewScopes(tt.args.strs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScopes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScopes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScopes_ContainsOpenId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		scopes internal.Scopes
		want   bool
	}{
		{
			name:   "should return true when contains openid",
			scopes: internal.Scopes{internal.OpenId, internal.Email},
			want:   true,
		},
		{
			name:   "should return false when not contains openid",
			scopes: internal.Scopes{internal.Email},
			want:   false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.scopes.ContainsOpenId(); got != tt.want {
				t.Errorf("ContainsOpenId() = %v, want %v", got, tt.want)
			}
		})
	}
}
