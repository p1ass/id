package internal_test

import (
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewGrantType(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    internal.GrantType
		wantErr bool
	}{
		{
			name: "should return GrantTypeAuthorizationCode when authorization_code",
			args: args{
				str: "authorization_code",
			},
			want:    internal.GrantTypeAuthorizationCode,
			wantErr: false,
		},
		{
			name: "should return error when unknown value",
			args: args{
				str: "unknown",
			},
			want:    internal.GrantTypeUnknown,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := internal.NewGrantType(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGrantType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewGrantType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
