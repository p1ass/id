package internal_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewResponseType(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    internal.ResponseType
		wantErr bool
	}{
		{
			name: "should return ResponseTypeCode when code",
			args: args{
				str: "code",
			},
			want:    internal.ResponseTypeCode,
			wantErr: false,
		},
		{
			name: "should return error when unknown value",
			args: args{
				str: "unknown",
			},
			want:    internal.ResponseUnknown,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := internal.NewResponseType(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResponseType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewResponseType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponseTypes(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    internal.ResponseTypes
		wantErr bool
	}{
		{
			name: "should return response types when all args are valid",
			args: args{
				strs: []string{"code", "id_token"},
			},
			want:    []internal.ResponseType{internal.ResponseTypeCode, internal.ResponseTypeIdToken},
			wantErr: false,
		},
		{
			name: "should return error when one of the args is invalid",
			args: args{
				strs: []string{"code", "invalid"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := internal.NewResponseTypes(tt.args.strs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResponseTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("NewResponseTypes() diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestResponseTypes_ContainsOnlyCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    internal.ResponseTypes
		want bool
	}{
		{
			name: "should return true when contains only code",
			s:    internal.ResponseTypes{internal.ResponseTypeCode},
			want: true,
		},
		{
			name: "should return false when contains other response type",
			s:    internal.ResponseTypes{internal.ResponseTypeCode, internal.ResponseTypeIdToken},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.s.ContainsOnlyCode(); got != tt.want {
				t.Errorf("ContainsOnlyCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
